package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Lean002/fhirlib/client"
)

func main() {
	client := client.NewFHIRClient("https://aidbox-qa.bupa.cl", "admin", "secret")

	patientID := "a8fea2b4-c38f-46d4-b295-a8b8ee664d87"
	patient, err := client.GetPatient(patientID)
	if err != nil {
		log.Fatalf("Error obteniendo el paciente: %v", err)
	}

	patientJSON, err := json.MarshalIndent(patient, "", "  ")
	if err != nil {
		log.Fatalf("Error convirtiendo a JSON: %v", err)
	}

	fmt.Printf("Paciente obtenido en formato JSON: \n%s\n", string(patientJSON))
}
