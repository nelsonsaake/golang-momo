package sandbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"projects/investorsmarket/bjt/momotrials/collection"
)

// MoMo trials
// write an api req to get the api uer info after create
func GetApiUserInfo(UUID, apim_sub_key string) (response string, err error) {

	url := "https://sandbox.momodeveloper.mtn.com/v1_0/apiuser/" + UUID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request to get mtn momo api user info \nerr: %v", err)
		return
	}

	req.Header.Add("Ocp-Apim-Subscription-Key", apim_sub_key)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to sending request to get mtn momo api user \nerr: %v", err)
		return
	}

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body \nerr: %v", err)
		return
	}

	fmt.Println("get api user information payload")
	fmt.Println(string(payload))
	fmt.Println()

	switch res.StatusCode {

	case http.StatusOK:

	case http.StatusBadRequest:
		err = errors.New(res.Status)

	case http.StatusNotFound, http.StatusInternalServerError:
		err = errors.New(res.Status)
		var resErr collection.ErrorResponse
		err = json.Unmarshal(payload, &resErr)
		if err != nil {
			err = fmt.Errorf("failed to unmarshal response: %v", err)
			return
		}

		response = fmt.Sprint(resErr)

		fmt.Println("err details:")
		fmt.Println(string(payload))

	default:
		err = errors.New("Something went wrong!")
	}

	return
}
