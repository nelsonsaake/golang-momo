package sandbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/bjt/momotrials/collection"
)

// MoMo trials
// write an api req to create api user
func CreateApiUser(UUID, apim_sub_key string) (response string, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/v1_0/apiuser"
	reqBody := []byte(`{"providerCallbackHost":""}`)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		err = fmt.Errorf("failed to create request to create mtn momo api user \nerr: %v", err)
		return
	}

	req.Header.Add("X-Reference-Id", UUID)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", apim_sub_key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to sending request to create mtn momo api user \nerr :", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body \nerr: %v", err)
		return
	}

	fmt.Println("create api user payload")
	fmt.Println(string(payload))
	fmt.Println()

	switch res.StatusCode {

	case http.StatusCreated:
		response = string(payload)

	case http.StatusBadRequest, http.StatusInternalServerError:
		err = errors.New(res.Status)

	case http.StatusConflict:
		err = errors.New(res.Status)
		var resErr collection.ErrorResponse
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		response = fmt.Sprint(resErr)

	default:
		err = errors.New("Something went wrong!")
	}

	return
}
