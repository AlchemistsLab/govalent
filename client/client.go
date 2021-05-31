//Package client contains methods to make request to Covalent API server.
package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
	Code int    `json:"error_code"`
	Msg  string `json:"error_message"`
}

// Error returns error message from Covalent API.
func (e CovalentError) Error() string {
	return e.Msg
}

//New initializes API with given URL, api key. It also provides a way to overwrite *http.Client
func New(url, key string, httpClient *http.Client) *API {
	return &API{
		URL:        url,
		Key:        key,
		HTTPClient: httpClient,
	}
}

func (a *API) Request(method, endpoint string, params interface{}, out interface{}) error {
	u, err := url.ParseRequestURI(a.URL)
	if err != nil {
		return err
	}
	u.Path = u.Path + endpoint
	if method == "GET" {
		b, err := json.Marshal(params)
		if err != nil {
			return err
		}
		m := map[string]interface{}{}
		if err := json.Unmarshal(b, &m); err != nil {
			return err
		}
		q := u.Query()
		q.Set("key", a.Key)
		for k, v := range m {
			q.Set(k, fmt.Sprintf("%v", v))
		}
		u.RawQuery = q.Encode()
		log.Printf("send request: %v %v", method, u.String())
		req, err := http.NewRequest(method, u.String(), nil)
		if err != nil {
			return err
		}
		req.Header.Add("content-type", "application/json")
		res, err := a.HTTPClient.Do(req)
		if err != nil {
			return err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Printf("error while close response: %v", err)
			}
		}(res.Body)
		if res.StatusCode != 200 {
			e := CovalentError{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				return CovalentError{Msg: err.Error()}
			}
			return e
		}
		if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
			return CovalentError{Msg: "Invalid JSON"}
		}
	}
	return nil
}
