//Package class_a general provides the binding for Covalent Rest APIs Class A endpoints.
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
	Chains() (Chains, error)
	ChainsStatus() (Chains, error)
	TokenLists(chainID, id string) (Contracts, error)

	TokenBalances(chainID, address string, params BalanceParams) (Portfolios, error)
	HistoricalPortfolio(chainID, address string) (Portfolios, error)

	Transactions(chainID, address string) (Transactions, error)
	Transaction(chainID, txHash string) (Transactions, error)
	ERCTokenTransfers(chainID, address string, params TransferParams) (Transactions, error)

	Block(chainID, blockHeight string) (Blocks, error)
	LogEventsByContract(chainID, address string, params LogEventsParams) (LogEvents, error)
	LogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEvents, error)

	ExternalNFTMetadata(chainID, address, tokenID string) (NFTTokens, error)
	NFTTokenIDs(chainID, address string) (NFTTokens, error)
	NFTTransactions(chainID, address, tokenID string) (NFTTokens, error)

	TokenHoldersChanges(chainID, address string, params TokenHoldersParams) (TokenHoldersChanges, error)
	TokenHolders(chainID, address string) (TokenHolders, error)
}

var _ ClassAInterface = (*Client)(nil)

// Chains returns all chains.
func (c *Client) Chains() (Chains, error) {
	response := ChainsResponse{}
	err := c.API.Request("GET", "chains/", nil, &response)
	return response.Data, err
}

// ChainsStatus returns all chain statuses..
func (c *Client) ChainsStatus() (Chains, error) {
	response := ChainsResponse{}
	err := c.API.Request("GET", "chains/status/", nil, &response)
	return response.Data, err
}

// TokenLists returns a list of all contracts on a blockchain along with their metadata.
func (c *Client) TokenLists(chainID, id string) (Contracts, error) {
	u := fmt.Sprintf("%v/tokens/tokenlists/%v/", chainID, id)
	response := ContractsResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// TokenBalances returns a list of all ERC20 and NFT token balances along with their current spot prices.
func (c *Client) TokenBalances(chainID, address string, params BalanceParams) (Portfolios, error) {
	u := fmt.Sprintf("%v/address/%v/balances_v2/", chainID, address)
	response := BalanceResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// HistoricalPortfolio returns wallet value for the last 30 days at 24 hour timestamps
// for given chain_id and wallet address.
func (c *Client) HistoricalPortfolio(chainID, address string) (Portfolios, error) {
	u := fmt.Sprintf("%v/address/%v/portfolio_v2/", chainID, address)
	response := Portfolios{}
	err := c.API.Request("GET", u, nil, &response)
	return response, err
}

// Transactions retrieves all transactions for address including their decoded log events.
// This endpoint does a deep-crawl of the blockchain to retrieve all kinds of transactions
// that references the address.
func (c *Client) Transactions(chainID, address string) (Transactions, error) {
	u := fmt.Sprintf("%v/address/%v/transactions_v2/", chainID, address)
	response := TransactionResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// Transaction retrieves a single transaction for tx_hash including their decoded log events.
func (c *Client) Transaction(chainID, txHash string) (Transactions, error) {
	u := fmt.Sprintf("%v/transaction_v2/%v/", chainID, txHash)
	response := TransactionResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// ERCTokenTransfers returns ERC20 token transfers. Passing in an ENS resolves automatically.
func (c *Client) ERCTokenTransfers(chainID, address string, params TransferParams) (Transactions, error) {
	u := fmt.Sprintf("%v/address/%v/transfers_v2/", chainID, address)
	response := TransactionResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// Block retrieves a single block at block_height.
// If block_height is set to the value latest, return the latest block available.
func (c *Client) Block(chainID, blockHeight string) (Blocks, error) {
	u := fmt.Sprintf("%v/block_v2/%v/", chainID, blockHeight)
	response := BlockResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// LogEventsByContract returns a paginated list of decoded log events emitted by a particular smart contract.
func (c *Client) LogEventsByContract(chainID, address string, params LogEventsParams) (LogEvents, error) {
	u := fmt.Sprintf("%v/events/address/%v/", chainID, address)
	response := LogEventsResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// LogEventsByTopic returns a paginated list of decoded log events with one or more topic hashes
// separated by a comma.
func (c *Client) LogEventsByTopic(chainID, topic string, params LogEventsParams) (LogEvents, error) {
	u := fmt.Sprintf("%v/events/topics/%v/", chainID, topic)
	response := LogEventsResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// ExternalNFTMetadata returns the external metadata for given a NFT contract address and a token ID.
func (c *Client) ExternalNFTMetadata(chainID, address, tokenID string) (NFTTokens, error) {
	u := fmt.Sprintf("%v/tokens/%v/nft_metadata/%v/", chainID, address, tokenID)
	response := NFTTokenResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// NFTTokenIDs returns a list of all token IDs for a NFT contract on a blockchain network.
func (c *Client) NFTTokenIDs(chainID, address string) (NFTTokens, error) {
	u := fmt.Sprintf("%v/tokens/%v/nft_token_ids/", chainID, address)
	response := NFTTokenResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// NFTTransactions returns a list of transactions given a NFT contract and a token ID on a blockchain network.
func (c *Client) NFTTransactions(chainID, address, tokenID string) (NFTTokens, error) {
	u := fmt.Sprintf("%v/tokens/%v/nft_transactions/%v/", chainID, address, tokenID)
	response := NFTTokenResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}

// TokenHoldersChanges gets token balance changes for token holders between starting-block and ending-block.
// Returns a paginated list of token holders and their current/historical balances. If ending-block is omitted,
// the latest block is used. Note: Token holder balances exclude passive rewards through static reflection.
func (c *Client) TokenHoldersChanges(chainID, address string, params TokenHoldersParams) (TokenHoldersChanges, error) {
	u := fmt.Sprintf("%v/tokens/%v/token_holders_changes/", chainID, address)
	response := TokenHoldersChangesResponse{}
	err := c.API.Request("GET", u, params, &response)
	return response.Data, err
}

// TokenHolders returns a paginated list of token holders. If block-height is omitted, the latest block is used.
// Note: Token holder balances exclude passive rewards through static reflection.
func (c *Client) TokenHolders(chainID, address string) (TokenHolders, error) {
	u := fmt.Sprintf("%v/tokens/%v/token_holders/", chainID, address)
	response := TokenHoldersResponse{}
	err := c.API.Request("GET", u, nil, &response)
	return response.Data, err
}
