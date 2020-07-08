package collection

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MoMo trials

// request to pay
type Payer struct {
	PartyIdType string
	PartyId     string
}

type CreateRequestToPayRequest struct {
	Amount       string
	Currency     string
	ExternalId   string
	Payer        Payer
	PayerMessage string
	PayeeNote    string
}

type ErrorResponse struct {
	Code    string
	Message string
}

func CreateRequestToPay(apiUser, apiKey, apimSubKey, uuid, targetEvn string, reqBody CreateRequestToPayRequest) (resErr ErrorResponse, err error) {

	var url = "https://sandbox.momodeveloper.mtn.com/collection/v1_0/requesttopay"
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		err = fmt.Errorf("Error marshalling json%v", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		err = fmt.Errorf("failed to create request to request-to-pay \nerr: %v", err)
		return
	}

	req.SetBasicAuth(apiUser, apiKey)
	req.Header.Add("Ocp-Apim-Subscription-Key", apimSubKey)
	req.Header.Add("X-Callback-Url", "")
	req.Header.Add("X-Reference-Id", uuid)
	req.Header.Add("X-Target-Environment", targetEvn)
	req.Header.Add("Content-Type", "application/json")

	var client = &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to sending request to request-to-pay \nerr: %v", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: err %v", err)
		return
	}

	switch res.StatusCode {

	case http.StatusBadRequest:
		err = fmt.Errorf("failed to create request-to-pay%v", res.Status)

	case http.StatusConflict, http.StatusInternalServerError:
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}
		err = errors.New(res.Status)

	default:
		err = fmt.Errorf("some thing went wrong!, %v", res.Status)
	}

	return
}
