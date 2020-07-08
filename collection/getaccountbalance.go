package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountBalance struct {
	AvailableBalance string
	Currency         string
}

func GetAccountBalance(apiuser, apikey, targetEnv, apimSubKey string) (reqRes interface{}, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/collection/v1_0/account/balance"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request to get account balance: %v", err)
		return
	}

	req.SetBasicAuth(apiuser, apikey)
	req.Header.Add("X-Target-Environment", targetEnv)
	req.Header.Add("Ocp-Apim-Subscription-Key", apimSubKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send request to get account balance: %v", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body: err %v", err)
		return
	}

	switch res.StatusCode {

	case http.StatusOK:
		var accBal AccountBalance
		err = json.Unmarshal(payload, &accBal)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}
		reqRes = accBal

	case http.StatusBadRequest:
		err = fmt.Errorf("failed to get account balance: %v", res.Status)

	case http.StatusInternalServerError:
		var resErr ErrorResponse
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		reqRes = resErr
		err = errors.New(res.Status)

	default:
		err = fmt.Errorf("Something went wrong!, %v", res.Status)
	}

	return
}
