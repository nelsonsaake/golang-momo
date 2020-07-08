package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MoMo trials

type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIn   uint64
}

type GetTokenErrorResponse struct {
	Error string
}

// write an api req to get barrer token
func GetToken(apiuser, apikey, apimsubkey string) (reqRes interface{}, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/collection/token/"
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request to get token \nerr: %v", err)
		return
	}

	req.SetBasicAuth(apiuser, apikey)
	req.Header.Add("Ocp-Apim-Subscription-Key", apimsubkey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to sending request to get token \nerr: %v", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body \nerr: %v", err)
		return
	}

	switch res.StatusCode {

	case http.StatusOK:
		var token Token
		err = json.Unmarshal(payload, &token)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}
		reqRes = token

	case http.StatusUnauthorized:
		var resErr GetTokenErrorResponse
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		reqRes = resErr
		err = errors.New(res.Status)

	case http.StatusInternalServerError:
		err = fmt.Errorf("failed to get token: %v", res.Status)

	default:
		err = fmt.Errorf("Something went wrong!, %v", res.Status)
	}

	return
}
