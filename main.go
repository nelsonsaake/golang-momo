package main

import (
	"fmt"
	"os"
	"projects/investorsmarket/bjt/momotrials/sandbox"
	"projects/investorsmarket/bjt/momotrials/uuid"

	"github.com/olekukonko/tablewriter"
)

func doTheThingWithError(msg string, err error) {
	fmt.Println(msg)
	fmt.Println(err)
}

func requestToPay() {

	// get api user
	/// create a uuid for api user x-reference
	subKey := ""
	xRef, err := uuid.GetUUID()
	if err != nil {
		doTheThingWithError("error creating api user, failed to uuid", err)
		return
	}

	_, err = sandbox.CreateApiUser(xRef, subKey)
	if err != nil {
		doTheThingWithError("error creating api user, something went wrong", err)
		return
	}

	/// get the user api information
	apiUser, err := sandbox.GetApiUserInfo(xRef, subKey)
	if err != nil {
		doTheThingWithError("error getting api user information", err)
		return
	}

	// get api key
	apiKey, err := sandbox.CreateApiKey(xRef, subKey)
	if err != nil {
		doTheThingWithError("error", err)
		return
	}

	data := [][]string{
		[]string{"Name", "Value"},
		[]string{"X-Reference-Id", xRef},
		[]string{"api user", apiUser},
		[]string{"api key", apiKey},
		[]string{"Ocp-Apim-Subscription-Key", subKey},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
	table.AppendBulk(data)
	table.Render()

}

func main() {

	requestToPay()
}
