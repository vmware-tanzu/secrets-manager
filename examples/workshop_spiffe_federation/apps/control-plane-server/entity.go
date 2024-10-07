package main

type Endpoint struct {
	Name              string `json:"name"`
	BundleEndpointURL string `json:"bundleEndpointUrl"`
	TrustDomain       string `json:"trustDomain"`
	EndpointSPIFFEID  string `json:"endpointSPIFFEID"`
}

type Endpoints struct {
	Endpoints    map[string]Endpoint `json:""`
	FederateWith []string            `json:"federateWith"`
}

type Secret struct {
	Name         string   `json:"name"`
	Value        []string `json:"value"`
	Created      string   `json:"created"`
	Updated      string   `json:"updated"`
	NotBefore    string   `json:"notBefore"`
	ExpiresAfter string   `json:"expiresAfter"`
}

type EncryptedResponse struct {
	EncryptedAESKey string `json:"encryptedAESKey"`
	EncryptedData   string `json:"encryptedData"`
	Signature       string `json:"signature"`
}

type Secrets struct {
	Secrets   []Secret `json:"secrets"`
	Algorithm string   `json:"algorithm"`
}
