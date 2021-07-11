package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var siteid string = "appointplus430/360"
var apiKey = "f5c9e47015ecd81423b5a65aef0240e9da17b921"
var serviceIds = []string{"7749", "1658", "7750"}

func main() {
	fmt.Println(siteid, apiKey, serviceIds)

	fmt.Println(wsCall("Services/GetServices", map[string]string{
		"response_type": "json",
		"service_id":    strings.Join(serviceIds, ",")}))

}

func wsCall(endpoint string, params map[string]string) string {
	base := "https://ws.appointment-plus.com/"
	baseEndpoint := base + endpoint

	apiGun := &http.Client{Timeout: time.Second * 10}
	paramValues := url.Values{}

	for k, v := range params {
		paramValues.Set(k, v)
	}

	fullURL := baseEndpoint + "?" + paramValues.Encode()

	req, err := http.NewRequest("POST", fullURL, nil)

	if err != nil {
		return fmt.Errorf("Got error %s", err.Error()).Error()
	}

	req.SetBasicAuth(siteid, apiKey)

	res, err := apiGun.Do(req)
	if err != nil {
		return fmt.Errorf("Got error %s", err.Error()).Error()
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(result)
}
