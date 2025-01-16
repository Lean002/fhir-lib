package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

// GetPatient gets a patient from the FHIR server using the ID

func (c *FHIRClient) GetPatient(id string) (*fhir.Patient, error) {
	url := fmt.Sprintf("%s/Patient/%s", c.BaseURL, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creando la solicitud GET: %w", err)
	}

	// Usar autenticación básica
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Set("Content-Type", "application/fhir+json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error realizando la solicitud GET: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error obteniendo el paciente, código de estado: %d, detalle: %s",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error leyendo la respuesta del servidor: %w", err)
	}

	var patient fhir.Patient
	if err := json.Unmarshal(body, &patient); err != nil {
		return nil, fmt.Errorf("error deserializando el recurso Patient: %w", err)
	}

	return &patient, nil
}
