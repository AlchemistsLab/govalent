//Package class_a general provides the binding for Covalent Rest APIs Class A endpoints
package class_a

import (
	"fmt"
	"github.com/AlchemistsLab/govalent/client"
)

type Client struct {
	API client.API
}

// ClassAInterface describes methods for Class A endpoints.
type ClassAInterface interface {
	GetTokenBalances(chainID, address string, params BalanceParams) (PortfolioData, error)
	GetHistoricalPortfolio(chainID, address string) (PortfolioData, error)
	GetTransactions(chainID, address string) (TransactionData, error)
	GetERCTokenTransfers(chainID, address string, params TransferParams) (TransactionData, error)
	GetBlock(chainID, blockHeight string) (BlockData, error)
	GetLogEventsByContract(chainID, address string, params LogEventsParams) (LogEventsData, error)
	GetLogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEventsData, error)
}

var _ ClassAInterface = (*Client)(nil)

// GetTokenBalances returns a list of all ERC20 and NFT token balances along with their current spot prices.
func (c *Client) GetTokenBalances(chainID, address string, params BalanceParams) (PortfolioData, error) {
	u := fmt.Sprintf("%v/address/%v/balances_v2/", chainID, address)
	balance := Balance{}
	err := c.API.Request("GET", u, params, &balance)
	return balance.Data, err
}

// GetHistoricalPortfolio returns wallet value for the last 30 days at 24 hour timestamps
// for given chain_id and wallet address.
func (c *Client) GetHistoricalPortfolio(chainID, address string) (PortfolioData, error) {
	u := fmt.Sprintf("%v/address/%v/portfolio_v2/", chainID, address)
	portfolio := PortfolioData{}
	err := c.API.Request("GET", u, nil, &portfolio)
	return portfolio, err
}

// GetTransactions retrieves all transactions for address including their decoded log events.
//This endpoint does a deep-crawl of the blockchain to retrieve all kinds of transactions that references the address.
func (c *Client) GetTransactions(chainID, address string) (TransactionData, error) {
	u := fmt.Sprintf("%v/address/%v/transactions_v2/", chainID, address)
	transaction := Transaction{}
	err := c.API.Request("GET", u, nil, &transaction)
	return transaction.Data, err
}

// GetERCTokenTransfers returns ERC20 token transfers. Passing in an ENS resolves automatically.
func (c *Client) GetERCTokenTransfers(chainID, address string, params TransferParams) (TransactionData, error) {
	u := fmt.Sprintf("%v/address/%v/transfers_v2/", chainID, address)
	transaction := Transaction{}
	err := c.API.Request("GET", u, params, &transaction)
	return transaction.Data, err
}

// GetBlock retrieves a single block at block_height.
// If block_height is set to the value latest, return the latest block available.
func (c *Client) GetBlock(chainID, blockHeight string) (BlockData, error) {
	u := fmt.Sprintf("%v/block_v2/%v/", chainID, blockHeight)
	block := Block{}
	err := c.API.Request("GET", u, nil, &block)
	return block.Data, err
}

// GetLogEventsByContract returns a paginated list of decoded log events emitted by a particular smart contract.
func (c *Client) GetLogEventsByContract(chainID, address string, params LogEventsParams) (LogEventsData, error) {
	u := fmt.Sprintf("%v/events/address/%v/", chainID, address)
	logEvents := LogEvents{}
	err := c.API.Request("GET", u, params, &logEvents)
	return logEvents.Data, err
}

// GetLogEventsByTopic returns a paginated list of decoded log events with one or more topic hashes separated by a comma..
func (c *Client) GetLogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEventsData, error) {
	u := fmt.Sprintf("%v/events/topics/%v/", chainID, topic)
	logEvents := LogEvents{}
	err := c.API.Request("GET", u, params, &logEvents)
	return logEvents.Data, err
}
