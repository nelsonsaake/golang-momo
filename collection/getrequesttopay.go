package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetRequestToPayResponse struct {
	Amount                 float64
	Currency               string
	FinancialTransactionID uint64
	ExternalID             uint64
	Payer                  Payer
	Status                 string
}

func GetRequestToPay(uuid, apiuser, apikey, targetEnv, apimSubKey string) (reqResp interface{}, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/collection/v1_0/requesttopay/" + uuid
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request to get request-to-pay information, err: %v", err)
		return
	}

	req.SetBasicAuth(apiuser, apikey)
	req.Header.Add("X-Target-Environment", targetEnv)
	req.Header.Add("Ocp-Apim-Subscription-Key", apimSubKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send request to get request to pay information: %v", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: err %v", err)
		return
	}

	switch res.StatusCode {

	case http.StatusOK:
		var grtpRes GetRequestToPayResponse
		err = json.Unmarshal(payload, &grtpRes)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		reqResp = grtpRes

	case http.StatusBadRequest:
		err = fmt.Errorf("failed to get request to pay information: %v", res.Status)

	case http.StatusNotFound, http.StatusInternalServerError:
		var resErr ErrorResponse
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		reqResp = resErr
		err = errors.New(res.Status)

	default:
		err = fmt.Errorf("Something went wrong!, %v", res.Status)
	}

	return
}
