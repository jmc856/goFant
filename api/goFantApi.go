package api

import (
	"time"
	"net/http"
	"os"
	"bytes"
	"net/url"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

var BaseYahooApi = "https://fantasysports.yahooapis.com/fantasy/v2"

func DoRequest(apiUrl string, accessToken string) ([]byte, error) {
	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	data := url.Values{}
	data.Set("client_id", os.Getenv("client_id_gofant"))
	data.Add("client_secret", os.Getenv("client_secret_gofant"))

	req, err := http.NewRequest("GET", apiUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer " + accessToken)

	response, err := client.Do(req)
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if response.StatusCode == 401 {
		//TODO: Try refresh and recall
		return DoRequest(apiUrl, accessToken)
	}

	if response.StatusCode >= 400 {
		var YahooApiError YahooApiError
		err := xml.Unmarshal(contents, &YahooApiError)
		if err != nil {
			return nil, ApiError{
				Status: "1234",
				Message: fmt.Sprintf("Failed on api call to %s", apiUrl),
			}
		}
		return nil, ApiError{
			Status: "1001",
			Message: fmt.Sprint("%s %s", YahooApiError.Description, response.StatusCode),
		}
	}
	return contents, nil
}
