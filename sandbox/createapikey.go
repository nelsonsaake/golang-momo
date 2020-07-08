package sandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MoMo trials

type CreateApiKeyResponse struct {
	ApiKey string
}

// write an api req to create api key
func CreateApiKey(UUID, apim_sub_key string) (apiKey string, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/v1_0/apiuser/" + UUID + "/apikey"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request to create mtn momo api key \nerr: %v", err)
		return
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", apim_sub_key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to sending request to create mtn momo api key \nerr: %v", err)
		return
	}

	if res.StatusCode != http.StatusCreated {
		errMsg := fmt.Sprintf("error with the status code, expected ", http.StatusCreated, ", received ", res.StatusCode)
		errMsg = errMsg + fmt.Sprintf("failed to create api user")
		err = errors.New(errMsg)
		if res.StatusCode == 404 {
			payload, readErr := ioutil.ReadAll(res.Body)
			if readErr == nil {
				fmt.Println(string(payload))
			}
		}
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body \nerr: %v", err)
		return
	}

	var createApiKeyRes CreateApiKeyResponse
	err = json.Unmarshal(payload, &createApiKeyRes)
	apiKey = createApiKeyRes.ApiKey
	if err != nil {
		err = fmt.Errorf("failed to unmarhsal payload \nerr: %v", err)
		return
	}

	return
}
