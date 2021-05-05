//Package class_a general provides the binding for Covalent Rest APIs Class A endpoints
package class_a

import (
	"fmt"
	"github.com/zaebee/govalent/client"
)

type Client struct {
	API client.API
}

// ClassAInterface describes methods for Class A endpoints.
type ClassAInterface interface {
	GetTokenBalances(chainID, address string, params BalanceParams) (Portfolio, error)
	GetHistoricalPortfolio(chainID, address string) (Portfolio, error)
}

var _ ClassAInterface = (*Client)(nil)

// GetTokenBalances returns a list of all ERC20 and NFT token balances along with their current spot prices.
func (c *Client) GetTokenBalances(chainID, address string, params BalanceParams) (Portfolio, error) {
	u := fmt.Sprintf("/v1/%v/address/%v/balances_v2/", chainID, address)
	balance := Balance{}
	err := c.API.Request("GET", u, params, &balance)
	return balance.Data, err
}

// GetHistoricalPortfolio returns wallet value for the last 30 days at 24 hour timestamps
// for given chain_id and wallet address.
func (c *Client) GetHistoricalPortfolio(_, _ string) (Portfolio, error) {
	portfolio := Portfolio{}
	return portfolio, nil
}
