package goanda

import (
	"encoding/json"
	"fmt"
	"time"
)

func GetAccounts() (*accountsResponse, error) {
	const extraEndpoint = "/v3/accounts"
	api := NewOandaAPI(OANDA_API_KEY)
	fullEndpoint := api.baseDemoEnpoint + extraEndpoint
	res, bodyBytes, err := api.GetRequestJSON(fullEndpoint)
	if err != nil {
		fmt.Println("Error: ")
		return &accountsResponse{}, err
	}

	if res.StatusCode != 200 {
		err := fmt.Errorf("Non-200 code when getting accounts: %v", res.StatusCode)
		return &accountsResponse{}, err
	}

	var ar accountsResponse
	err = json.Unmarshal(bodyBytes, &ar)
	if err != nil {
		fmt.Println("Error: ")
		return &accountsResponse{}, err
	}

	return &ar, nil
}

func GetAccount(acctID string) (*accountResponse, error) {
	api := NewOandaAPI(OANDA_API_KEY)
	fullEndpoint := fmt.Sprintf("%s/v3/accounts/%s", api.baseDemoEnpoint, acctID)
	res, bodyBytes, err := api.GetRequestJSON(fullEndpoint)

	// fmt.Println(string(bodyBytes))

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

	return &ar, nil
}

type account struct {
	ID   string   `json:"id"`
	Tags []string `json:"tags"`
}

type accountDetails struct {
	Nav                         string    `json:"NAV"`
	Alias                       string    `json:"alias"`
	Balance                     string    `json:"balance"`
	CreatedByUserID             int       `json:"createdByUserID"`
	CreatedTime                 time.Time `json:"createdTime"`
	Currency                    string    `json:"currency"`
	HedgingEnabled              bool      `json:"hedgingEnabled"`
	ID                          string    `json:"id"`
	LastTransactionID           string    `json:"lastTransactionID"`
	MarginAvailable             string    `json:"marginAvailable"`
	MarginCloseoutMarginUsed    string    `json:"marginCloseoutMarginUsed"`
	MarginCloseoutNAV           string    `json:"marginCloseoutNAV"`
	MarginCloseoutPercent       string    `json:"marginCloseoutPercent"`
	MarginCloseoutPositionValue string    `json:"marginCloseoutPositionValue"`
	MarginCloseoutUnrealizedPL  string    `json:"marginCloseoutUnrealizedPL"`
	MarginRate                  string    `json:"marginRate"`
	MarginUsed                  string    `json:"marginUsed"`
	OpenPositionCount           int       `json:"openPositionCount"`
	OpenTradeCount              int       `json:"openTradeCount"`
	Orders                      []any     `json:"orders"`
	PendingOrderCount           int       `json:"pendingOrderCount"`
	Pl                          string    `json:"pl"`
	PositionValue               string    `json:"positionValue"`
	Positions                   []struct {
		Instrument string `json:"instrument"`
		Long       struct {
			Pl           string `json:"pl"`
			ResettablePL string `json:"resettablePL"`
			Units        string `json:"units"`
			UnrealizedPL string `json:"unrealizedPL"`
		} `json:"long"`
		Pl           string `json:"pl"`
		ResettablePL string `json:"resettablePL"`
		Short        struct {
			Pl           string `json:"pl"`
			ResettablePL string `json:"resettablePL"`
			Units        string `json:"units"`
			UnrealizedPL string `json:"unrealizedPL"`
		} `json:"short"`
		UnrealizedPL string `json:"unrealizedPL"`
	} `json:"positions"`
	ResettablePL    string `json:"resettablePL"`
	Trades          []any  `json:"trades"`
	UnrealizedPL    string `json:"unrealizedPL"`
	WithdrawalLimit string `json:"withdrawalLimit"`
}

type accountsResponse struct {
	Accounts []account `json:"accounts"`
}

type accountResponse struct {
	Account accountDetails `json:"account"`
}
