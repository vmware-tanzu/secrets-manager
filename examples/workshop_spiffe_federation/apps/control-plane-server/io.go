package main

import (
	"encoding/json"
	"os"
)

func loadEndpoints(filename string) Endpoints {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading endpoints file: " + err.Error())
	}

	var endpoints Endpoints
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		panic("Error parsing endpoints JSON: " + err.Error())
	}

	return endpoints
}

func loadSecrets(filename string) Secrets {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("Error reading secrets file: " + err.Error())
	}

	var secrets Secrets
	err = json.Unmarshal(data, &secrets)
	if err != nil {
		panic("Error parsing secrets JSON: " + err.Error())
	}

	return secrets
}
