package api

import (
	"agent/internal/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendMetrics(dto [31]dto.Metric, serverURL string) error {
	fmt.Println(1)
	client := &http.Client{}
	metricsURL := fmt.Sprintf("%s/api/updates", serverURL)

	// dto to JSON
	jsonStr, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// create io.Reader из JSON
	response, err := client.Post(metricsURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(body))

	if response.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}
