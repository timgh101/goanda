package goanda

import (
	"encoding/json"
	"fmt"
)

func GetAccounts() (*accountResponse, error) {
	const extraEndpoint = "/v3/accounts"
	api := NewOandaAPI(OANDA_API_KEY)
	fullEndpoint := api.baseDemoEnpoint + extraEndpoint
	res, bodyBytes, err := api.GetRequestJSON(fullEndpoint)
	if err != nil {
		fmt.Println("Error: ")
		return &accountResponse{}, err
	}

	if res.StatusCode != 200 {
		err := fmt.Errorf("Non-200 code when getting accounts: %v", res.StatusCode)
		return &accountResponse{}, err
	}

	var ar accountResponse
	err = json.Unmarshal(bodyBytes, &ar)
	if err != nil {
		fmt.Println("Error: ")
		return &accountResponse{}, err
	}

	fmt.Println("Response for accounts: ")
	fmt.Println(ar)
	return &ar, nil
}

type account struct {
	ID   string   `json:"id"`
	Tags []string `json:"tags"`
}

type accountResponse struct {
	Accounts []account `json:"accounts"`
}
