//Package govalent provides the binding for Covalent Rest APIs.
package govalent

import (
	"github.com/zaebee/govalent/class_a"
	"github.com/zaebee/govalent/client"
	"net/http"
	"time"
)

//constant used for API client
const (
	APIURL             = "https://api.covalenthq.com/v1/"
	defaultHTTPTimeout = 80 * time.Second
)

//Default HTTP client
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

// APIKey is Covalent API Key.
var APIKey string

// ClassA uses endpoint without client.
func ClassA() *class_a.Client {
	api := client.New(APIURL, APIKey, httpClient)
	return &class_a.Client{API: *api}
}

//Client is the Covalent client. It contains all resources available.
type Client struct {
	ClassA class_a.Client
}

//Init initializes the Binance client with given API key, secret key.
func (c *Client) Init(apiKey string) {
	api := client.New(APIURL, apiKey, httpClient)
	c.ClassA = class_a.Client{API: *api}
}
