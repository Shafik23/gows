package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var siteid string = "appointplus430/360"
var apiKey = ""
var serviceIds = []string{"7749", "1658", "7750"}

func main() {
	// fmt.Println(siteid, apiKey, serviceIds)

	for _, serviceId := range serviceIds {
		jsonResponseMap := wsCall("Appointments/GetOpenDates", map[string]string{
			"response_type": "json",
			"service_id":    serviceId,
			"num_days":      "20"})

		dates, ok := jsonResponseMap["data"].([]interface{})

		if !ok {
			panic("Received empty data from WebServices - check your credentials and try again")
		}

		firstResult := dates[0].(map[string]interface{})

		firstAvailableDate := firstResult["date"].(string)

		fmt.Println("Service ID, First Available Date:", serviceId, firstAvailableDate)
		fmt.Println("--------")
	}
}

func wsCall(endpoint string, params map[string]string) map[string]interface{} {
	base := "https://ws.appointment-plus.com/"
	baseEndpoint := base + endpoint

	apiGun := &http.Client{Timeout: time.Second * 30}
	paramValues := url.Values{}

	for k, v := range params {
		paramValues.Set(k, v)
	}

	fullURL := baseEndpoint + "?" + paramValues.Encode()

	req, err := http.NewRequest("POST", fullURL, nil)

	if err != nil {
		panic(fmt.Errorf("Could not build URL: %s", err.Error()).Error())
	}

	req.SetBasicAuth(siteid, apiKey)

	res, err := apiGun.Do(req)
	if err != nil {
		panic(fmt.Errorf("API error from WebServices: %s", err.Error()).Error())
	}

	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("Error reading response body: %s", err.Error()).Error())
	}

	jsonMap := make(map[string]interface{})
	json.Unmarshal([]byte(result), &jsonMap)

	return jsonMap
}
