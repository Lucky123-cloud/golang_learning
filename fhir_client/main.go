package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	patientURL := "https://hapi.fhir.org/baseR4/Patient/example"
	outputFile := "patient.json"

	data, err := fetchPatient(patientURL)
	check(err)

	err = saveToFile(outputFile, data)
	check(err)

	err = validateFHIR(outputFile)
	check(err)

	fmt.Println("Patient validated successfully")
}

// fetchPatient retrieves a Patient resource from a FHIR server
func fetchPatient(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch patient: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Basic sanity check: ensure it's a Patient resource
	var resource map[string]interface{}
	if err := json.Unmarshal(body, &resource); err != nil {
		return nil, fmt.Errorf("invalid JSON received: %w", err)
	}

	if resource["resourceType"] != "Patient" {
		return nil, fmt.Errorf("resource is not a Patient")
	}

	return body, nil
}

// saveToFile writes JSON data to disk
func saveToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// validateFHIR runs the official FHIR validator CLI
func validateFHIR(filename string) error {
	cmd := exec.Command("java", "-jar", "validator_cli.jar", filename)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("validation failed:\n%s", string(output))
	}

	fmt.Println(string(output))
	return nil
}

// check fails fast (Pragmatic Programmer principle)
func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
