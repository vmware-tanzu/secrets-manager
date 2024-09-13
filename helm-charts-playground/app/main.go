package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

func removeComments(input string) string {
	var result strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(strings.TrimSpace(line), "#") {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	return result.String()
}

func createManifests(inputFile, outputDir string) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	documents := strings.Split(string(content), "---\n")

	for i, doc := range documents {
		if strings.TrimSpace(doc) == "" {
			continue
		}

		doc = removeComments(doc)

		var m map[string]interface{}
		if err := yaml.Unmarshal([]byte(doc), &m); err != nil {
			fmt.Printf("Error parsing YAML document %d: %v\n", i+1, err)
			continue
		}

		kind, ok := m["kind"].(string)
		if !ok {
			fmt.Printf("Error: 'kind' not found or not a string in document %d\n", i+1)
			continue
		}

		metadata, ok := m["metadata"].(map[string]interface{})
		if !ok {
			fmt.Printf("Error: 'metadata' not found or not a map in document %d\n", i+1)
			continue
		}

		name, ok := metadata["name"].(string)
		if !ok {
			fmt.Printf("Error: 'name' not found or not a string in document %d\n", i+1)
			continue
		}

		filename := fmt.Sprintf("%s-%s.yaml", strings.ToLower(kind), strings.ToLower(name))
		outputPath := filepath.Join(outputDir, filename)

		if err := os.WriteFile(outputPath, []byte(doc), 0644); err != nil {
			fmt.Printf("Error writing document %d to file: %v\n", i+1, err)
			continue
		}

		fmt.Printf("Created file: %s\n", outputPath)
	}
}

func copyVSecMCrds(inputDir, outputDir string) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	files, err := os.ReadDir(inputDir)
	if err != nil {
		fmt.Printf("Error reading input directory: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		inputPath := filepath.Join(inputDir, file.Name())
		outputPath := filepath.Join(outputDir, file.Name())

		content, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", inputPath, err)
			continue
		}

		content = []byte(removeComments(string(content)))

		if err := os.WriteFile(outputPath, content, 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", outputPath, err)
			continue
		}

		fmt.Printf("Copied file: %s\n", outputPath)
	}
}

func main() {
	inputFile := "/Users/volkan/Desktop/WORKSPACE/secrets-manager/k8s/0.27.2/spire.yaml"
	outputDir := "/Users/volkan/Desktop/WORKSPACE/secrets-manager/helm-charts-playground/vsecm-manifests"

	createManifests(inputFile, outputDir)

	//inputDir := "/Users/volkan/Desktop/WORKSPACE/secrets-manager/k8s/0.27.2/crds"
	//outputDir = "/Users/volkan/Desktop/WORKSPACE/secrets-manager/helm-charts-playground/vsecm-manifests/crds"
	//
	//copyVSecMCrds(inputDir, outputDir)

	inputFile = "/Users/volkan/Desktop/WORKSPACE/secrets-manager/helm-charts-playground/spire-manifest-no-openshift.yaml"
	outputDir = "/Users/volkan/Desktop/WORKSPACE/secrets-manager/helm-charts-playground/helm-charts-manifests"

	createManifests(inputFile, outputDir)
}
