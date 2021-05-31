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
	GetTokenBalances(chainID, address string, params BalanceParams) (Portfolios, error)
	GetHistoricalPortfolio(chainID, address string) (Portfolios, error)
	GetTransactions(chainID, address string) (Transactions, error)
	GetERCTokenTransfers(chainID, address string, params TransferParams) (Transactions, error)
	GetBlock(chainID, blockHeight string) (Blocks, error)
	GetLogEventsByContract(chainID, address string, params LogEventsParams) (LogEvents, error)
	GetLogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEvents, error)
	GetExternalNFTMetadata(chainID, address, tokenID string) (NFTMetadata, error)
}

var _ ClassAInterface = (*Client)(nil)

// GetTokenBalances returns a list of all ERC20 and NFT token balances along with their current spot prices.
func (c *Client) GetTokenBalances(chainID, address string, params BalanceParams) (Portfolios, error) {
	u := fmt.Sprintf("%v/address/%v/balances_v2/", chainID, address)
	response := BalanceResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// GetHistoricalPortfolio returns wallet value for the last 30 days at 24 hour timestamps
// for given chain_id and wallet address.
func (c *Client) GetHistoricalPortfolio(chainID, address string) (Portfolios, error) {
	u := fmt.Sprintf("%v/address/%v/portfolio_v2/", chainID, address)
	response := Portfolios{}
	err := c.API.Request("GET", u, nil, &response)
	return response, err
}

// GetTransactions retrieves all transactions for address including their decoded log events.
//This endpoint does a deep-crawl of the blockchain to retrieve all kinds of transactions that references the address.
func (c *Client) GetTransactions(chainID, address string) (Transactions, error) {
	u := fmt.Sprintf("%v/address/%v/transactions_v2/", chainID, address)
	response := TransactionResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// GetERCTokenTransfers returns ERC20 token transfers. Passing in an ENS resolves automatically.
func (c *Client) GetERCTokenTransfers(chainID, address string, params TransferParams) (Transactions, error) {
	u := fmt.Sprintf("%v/address/%v/transfers_v2/", chainID, address)
	response := TransactionResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// GetBlock retrieves a single block at block_height.
// If block_height is set to the value latest, return the latest block available.
func (c *Client) GetBlock(chainID, blockHeight string) (Blocks, error) {
	u := fmt.Sprintf("%v/block_v2/%v/", chainID, blockHeight)
	response := BlockResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// GetLogEventsByContract returns a paginated list of decoded log events emitted by a particular smart contract.
func (c *Client) GetLogEventsByContract(chainID, address string, params LogEventsParams) (LogEvents, error) {
	u := fmt.Sprintf("%v/events/address/%v/", chainID, address)
	response := LogEventsResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// GetLogEventsByTopic returns a paginated list of decoded log events with one or more topic hashes separated by a comma.
func (c *Client) GetLogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEvents, error) {
	u := fmt.Sprintf("%v/events/topics/%v/", chainID, topic)
	response := LogEventsResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// GetExternalNFTMetadata returns the external metadata for given a NFT contract address and a token ID.
func (c *Client) GetExternalNFTMetadata(chainID, address, tokenID string) (NFTMetadata, error) {
	u := fmt.Sprintf("%v/tokens/%v/nft_metadata/%v/", chainID, address, tokenID)
	response := NFTExternalMetadataResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}
