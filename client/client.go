//Package client contains methods to make request to Covalent API server.
package client

import (
	"net/http"
)

//API is a Binance API client.
type API struct {
	URL        string
	Key        string
	HTTPClient *http.Client
	UserAgent  string
}

// CovalentError handles api errors from covalenthq.com.
type CovalentError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error returns error message from Covalent API.
func (e CovalentError) Error() string {
	return e.Msg
}

//New initializes API with given URL, api key. It also provides a way to overwrite *http.Client
func New(url, key string, httpClient *http.Client, userAgent string) *API {
	return &API{
		URL:        url,
		Key:        key,
		HTTPClient: httpClient,
		UserAgent:  userAgent,
	}
}

func (a *API) Request(method, endpoint string, params interface{}, out interface{}) error {
	return nil
}
