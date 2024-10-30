package api

import (
	"agent/internal/dto"
	"agent/internal/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendMetrics(dto [31]dto.Metric, serverURL string, key string) error {
	fmt.Println(1)
	client := &http.Client{}
	metricsURL := fmt.Sprintf("%s/api/updates", serverURL)

	// dto to JSON
	jsonStr, err := json.Marshal(dto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	hash := utils.GenerateHash(key)

	fmt.Println("JSON Payload:", string(jsonStr))

	// create new request with header Hash
	req, err := http.NewRequest("POST", metricsURL, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Hash", hash)

	// do request
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.StatusCode)

	if response.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}
