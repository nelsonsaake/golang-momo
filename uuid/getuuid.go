package uuid

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MoMo trials

// write an api req to get a UUID
func GetUUID() (UUID string, err error) {

	url := "http://www.uuidgenerator.net/api/version4"
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("failed to sending request to get UUID \nerr: %v", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.New("error with the status code, expected status to be OK")
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body \nerr: %v", err)
		return
	}

	UUID = string(payload)

	return
}
