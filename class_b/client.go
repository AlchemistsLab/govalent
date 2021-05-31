//Package class_b general provides the binding for Covalent Rest APIs Class B endpoints.
package class_b

import (
	"fmt"
	"github.com/AlchemistsLab/govalent/client"
)

type Client struct {
	API client.API
}

// ClassBInterface describes methods for Class B endpoints.
type ClassBInterface interface {
	SushiSwapActs(chainID, address string, params SushiSwapActsParams) (SushiSwapActs, error)
	SushiSwapBalances(chainID, address string) (Response, error)
	SushiSwapAssets(chainID string) (Response, error)
}

var _ ClassBInterface = (*Client)(nil)

// SushiSwapActs returns Sushiswap address exchange liquidity transactions.
func (c *Client) SushiSwapActs(chainID, address string, params SushiSwapActsParams) (SushiSwapActs, error) {
	u := fmt.Sprintf("%v/address/%v/stacks/sushiswap/acts/", chainID, address)
	response := SushiSwapActsResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// SushiSwapBalances returns Sushiswap address exchange balances. Passing in an ENS resolves automatically.
func (c *Client) SushiSwapBalances(chainID, address string) (Response, error) {
	u := fmt.Sprintf("%v/address/%v/stacks/sushiswap/balances/", chainID, address)
	response := Response{}
	err := c.API.Request("GET", u, nil, &response)
	return response, err
}

// SushiSwapAssets returns a paginated list of Sushiswap pools sorted by exchange volume.
func (c *Client) SushiSwapAssets(chainID string) (Response, error) {
	u := fmt.Sprintf("%v/networks/sushiswap/assets/", chainID)
	response := Response{}
	err := c.API.Request("GET", u, nil, &response)
	return response, err
}
