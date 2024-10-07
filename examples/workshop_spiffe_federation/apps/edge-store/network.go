package main

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

func run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr("unix:///spire-agent-socket/spire-agent.sock"),
		),
	)

	if err != nil {
		panic("Error acquiring source")
	}

	defer func(source *workloadapi.X509Source) {
		err := source.Close()
		if err != nil {
			fmt.Println("error closing source")
		}
	}(source)

	svid, err := source.GetX509SVID()
	if err != nil {
		panic("error getting svid")
	}

	// example:
	// spiffe://mephisto.vsecm.com/workload/mephisto-edge-store/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
	svidId := svid.ID.String()
	parts := strings.Split(svidId, "/")
	// The workload name is the 4th element (index 3) in the split result
	workloadName := ""
	if len(parts) > 4 {
		workloadName = parts[4]
	}

	// https://github.com/vmware-tanzu/secrets-manager -- /examples/workshop_spiffe_federation/

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		// In a real-world scenario, you'd implement proper authorization logic here
		return nil
	})

	baseURL := os.Getenv("CONTROL_PLANE_URL")
	if baseURL == "" {
		baseURL = "https://10.211.55.112"
	}

	p, err := url.JoinPath(baseURL, "/")
	if err != nil {
		panic("problem in url " + baseURL)
	}

	// Generate a new keypair for this request
	privateKey, publicKey, err := generateKeyPair()
	if err != nil {
		panic("Error generating keypair")
	}

	// Encode the public key to PEM format
	publicKeyPEM, err := encodePublicKeyToPEM(publicKey)
	if err != nil {
		panic("Error encoding public key to PEM")
	}

	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   tlsconfig.MTLSClientConfig(source, source, authorizer),
		},
	}

	// Create a new request with the public key in the body
	req, err := http.NewRequest("POST", p, strings.NewReader(publicKeyPEM))
	if err != nil {
		panic("Error creating request")
	}
	req.Header.Set("Content-Type", "application/x-pem-file")

	r, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		panic("error getting from client")
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic("error reading body")
	}

	var response EncryptedResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic("error unmarshaling response")
	}

	// Decode the server's public key from the response header
	serverPublicKeyPEM, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Public-Key"))
	if err != nil {
		panic("error decoding server public key")
	}

	serverPublicKey, err := decodePublicKeyFromPEM(string(serverPublicKeyPEM))
	if err != nil {
		fmt.Println("err", err.Error())
		panic("error parsing server public key")
	}

	// Decode the encrypted AES key, encrypted data, and signature
	encryptedAESKey, err := base64.StdEncoding.DecodeString(response.EncryptedAESKey)
	if err != nil {
		panic("error decoding encrypted AES key")
	}

	encryptedData, err := base64.StdEncoding.DecodeString(response.EncryptedData)
	if err != nil {
		panic("error decoding encrypted data")
	}

	signature, err := base64.StdEncoding.DecodeString(response.Signature)
	if err != nil {
		panic("error decoding signature")
	}

	// Verify the signature
	err = verifySignature(encryptedData, signature, serverPublicKey)
	if err != nil {
		panic("signature verification failed")
	}

	// Decrypt the AES key using the client's private key
	aesKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedAESKey, nil)
	if err != nil {
		panic("error decrypting AES key")
	}

	// Decrypt the data using the AES key
	decryptedData, err := decryptAES(encryptedData, aesKey)
	if err != nil {
		panic("error decrypting data")
	}

	var secretValue []string
	err = json.Unmarshal(decryptedData, &secretValue)
	if err != nil {
		panic("error unmarshaling decrypted data")
	}

	fmt.Printf("My secret (encrypted with VSecM root key) : '%v'.\n", secretValue)

	if len(secretValue) == 0 {
		fmt.Println("No secret found")
		return
	}

	fmt.Println("")
	fmt.Println("")
	registerSecretToVSecM(workloadName, secretValue[0])
}

func registerSecretToVSecM(workloadName, secretValue string) {
	fmt.Println("will register the secret to VSecM for… ", workloadName)

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	// Assume we have a valid *rest.Config named config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	namespace := "vsecm-system"
	containerName := "" // Leave empty to use the first container
	podPrefix := "vsecm-sentinel"

	podName, err := getSentinelPodName(clientset, namespace, podPrefix)
	if err != nil {
		panic(err)
	}

	command := []string{"safe", "-w", workloadName, "-s", secretValue, "-e"}

	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: containerName,
			Command:   command,
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err)
	}

	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  nil,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Registered the secret to VSecM. Now fetching it with VSecM SDK.")
	fmt.Println("(this will be done in a separate workload, normally.")

	res, err := sentry.Fetch()
	if err != nil {
		fmt.Println("error fetching secret", err.Error())
		return
	}

	fmt.Println("my secret data (decrypted):", res.Data)

	fmt.Println()
	fmt.Println("---")
	fmt.Println()
}

func getSentinelPodName(clientset *kubernetes.Clientset, namespace, prefix string) (string, error) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list pods: %v", err)
	}

	for _, pod := range pods.Items {
		if strings.HasPrefix(pod.Name, prefix) {
			return pod.Name, nil
		}
	}

	return "", fmt.Errorf("no pod with prefix %s found in namespace %s", prefix, namespace)
}
