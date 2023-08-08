package goanda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://api-fxtrade.oanda.com"
const baseDemoURL = "https://api-fxpractice.oanda.com"

type oandaAPI struct {
	baseEnpoint     string
	baseDemoEnpoint string
	apiKey          string
	authHeader      string
}

func NewOandaAPI(apiKey string) *oandaAPI {
	o := oandaAPI{
		baseEnpoint:     baseURL,
		baseDemoEnpoint: baseDemoURL,
		apiKey:          apiKey,
		authHeader:      "Bearer " + apiKey,
	}
	return &o
}

// Makes a GET request with the correct auth header included
// and the header setup for a JSON response. Returns Byte Array of
// response body
func (o *oandaAPI) GetRequestJSON(url string) (*http.Response, []byte, error) {
	c := http.Client{Timeout: time.Duration(3) * time.Second}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return &http.Response{}, []byte{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", o.authHeader)
	resp, err := c.Do(req)
	if err != nil {
		return &http.Response{}, []byte{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return &http.Response{}, []byte{}, err
	}

	fmt.Println("Response Status = ", resp.Status)
	return resp, body, nil
}
