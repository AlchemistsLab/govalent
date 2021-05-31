package class_a

import "time"

// BalanceParams sets parameters for token balances endpoint.
type BalanceParams struct {
	Nft        bool `json:"nft"`
	NoNftFetch bool `json:"no-nft-fetch"`
}

// TransferParams sets parameters for ERC20 token transfer endpoint.
type TransferParams struct {
	ContractAddress string `json:"contract-address"`
	PageNumber      int    `json:"page-number"`
	// PageSize        int    `json:"page-size"`
}

// LogEventsParams sets parameters for get log events endpoint.
type LogEventsParams struct {
	StartingBlock string `json:"starting-block"`
	EndingBlock   string `json:"ending-block"`
	PageNumber    int    `json:"page-number"`
}

// Pagination returns pagination metadata for each endpoint.
type Pagination struct {
	HasMore    bool `json:"has_more"`
	PageNumber int  `json:"page_number"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
}

// BalanceResponse returns portfolio data for token balances response.
type BalanceResponse struct {
	Data         Portfolios `json:"data"`
	Error        bool       `json:"error"`
	ErrorMessage string     `json:"error_message"`
	ErrorCode    int        `json:"error_code"`
}

// TransactionResponse returns transaction data.
type TransactionResponse struct {
	Data         Transactions `json:"data"`
	Error        bool         `json:"error"`
	ErrorMessage string       `json:"error_message"`
	ErrorCode    int          `json:"error_code"`
}

// BlockResponse returns block data for given contract address.
type BlockResponse struct {
	Data         Blocks `json:"data"`
	Error        bool   `json:"error"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    int    `json:"error_code"`
}

// LogEventsResponse returns decoded log events.
type LogEventsResponse struct {
	UpdatedAt    time.Time `json:"updated_at"`
	Data         LogEvents `json:"data"`
	Error        bool      `json:"error"`
	ErrorMessage string    `json:"error_message"`
	ErrorCode    int       `json:"error_code"`
}

// NFTExternalMetadataResponse returns NFT external metadata for nft external metadata endpoint.
type NFTExternalMetadataResponse struct {
	Data         NFTTokens `json:"data"`
	Error        bool      `json:"error"`
	ErrorMessage string    `json:"error_message"`
	ErrorCode    int       `json:"error_code"`
}

// NFTTokenResponse returns NFT tokens data for given contract address.
type NFTTokenResponse struct {
	UpdatedAt    time.Time `json:"updated_at"`
	Data         NFTTokens `json:"data"`
	Error        bool      `json:"error"`
	ErrorMessage string    `json:"error_message"`
	ErrorCode    int       `json:"error_code"`
}

// Portfolios returns list of items for portfolio endpoint.
type Portfolios struct {
	Address       string      `json:"address"`
	UpdatedAt     time.Time   `json:"updated_at"`
	NextUpdateAt  time.Time   `json:"next_update_at"`
	QuoteCurrency string      `json:"quote_currency"`
	ChainID       int         `json:"chain_id"`
	Items         []Portfolio `json:"items"`
	Pagination    Pagination  `json:"pagination"`
}

// Transactions returns list of items for transaction endpoint.
type Transactions struct {
	Address       string        `json:"address"`
	UpdatedAt     time.Time     `json:"updated_at"`
	NextUpdateAt  time.Time     `json:"next_update_at"`
	QuoteCurrency string        `json:"quote_currency"`
	ChainId       int           `json:"chain_id"`
	Items         []Transaction `json:"items"`
	Pagination    Pagination    `json:"pagination"`
}

// Blocks returns list of items for block endpoint.
type Blocks struct {
	UpdatedAt time.Time `json:"updated_at"`
	Items     []struct {
		SignedAt time.Time `json:"signed_at"`
		Height   int       `json:"height"`
	} `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// LogEvents returns list of log items.
type LogEvents struct {
	UpdatedAt  time.Time  `json:"updated_at"`
	Items      []LogEvent `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// NFTTokens returns list of token items.
type NFTTokens struct {
	Items      []Portfolio `json:"items"`
	Pagination Pagination  `json:"pagination"`
}

type Portfolio struct {
	ContractDecimals     int        `json:"contract_decimals"`
	ContractName         string     `json:"contract_name"`
	ContractTickerSymbol string     `json:"contract_ticker_symbol"`
	ContractAddress      string     `json:"contract_address"`
	SupportsErc          []string   `json:"supports_erc"`
	LogoUrl              string     `json:"logo_url"`
	Type                 string     `json:"type"`
	Balance              string     `json:"balance"`
	QuoteRate            float64    `json:"quote_rate"`
	Quote                float64    `json:"quote"`
	TokenID              string     `json:"token_id"`
	NftData              []NftData  `json:"nft_data"`
	Holdings             []Holdings `json:"holdings"`
}

type Holdings struct {
	Timestamp time.Time `json:"timestamp"`
	QuoteRate float64   `json:"quote_rate"`
	Open      Holding   `json:"open"`
	High      Holding   `json:"high"`
	Low       Holding   `json:"low"`
	Close     Holding   `json:"close"`
}

type Holding struct {
	Balance string  `json:"balance"`
	Quote   float64 `json:"quote"`
}

type Transaction struct {
	BlockSignedAt    time.Time   `json:"block_signed_at"`
	TxHash           string      `json:"tx_hash"`
	TxOffset         int         `json:"tx_offset"`
	Successful       bool        `json:"successful"`
	FromAddress      string      `json:"from_address"`
	FromAddressLabel interface{} `json:"from_address_label"`
	ToAddress        string      `json:"to_address"`
	ToAddressLabel   interface{} `json:"to_address_label"`
	Value            string      `json:"value"`
	ValueQuote       float64     `json:"value_quote"`
	GasOffered       int         `json:"gas_offered"`
	GasSpent         int         `json:"gas_spent"`
	GasPrice         int64       `json:"gas_price"`
	GasQuote         float64     `json:"gas_quote"`
	GasQuoteRate     float64     `json:"gas_quote_rate"`
	LogEvents        []LogEvent  `json:"log_events"`
}

type LogEvent struct {
	BlockSignedAt              time.Time   `json:"block_signed_at"`
	BlockHeight                int         `json:"block_height"`
	TxOffset                   int         `json:"tx_offset"`
	LogOffset                  int         `json:"log_offset"`
	TxHash                     string      `json:"tx_hash"`
	RawLogTopicsBytes          interface{} `json:"_raw_log_topics_bytes"`
	RawLogTopics               []string    `json:"raw_log_topics"`
	SenderContractDecimals     interface{} `json:"sender_contract_decimals"`
	SenderName                 interface{} `json:"sender_name"`
	SenderContractTickerSymbol interface{} `json:"sender_contract_ticker_symbol"`
	SenderAddress              string      `json:"sender_address"`
	SenderAddressLabel         interface{} `json:"sender_address_label"`
	SenderLogoUrl              interface{} `json:"sender_logo_url"`
	RawLogData                 string      `json:"raw_log_data"`
	Decoded                    struct {
		Name      string `json:"name"`
		Signature string `json:"signature"`
		Params    []struct {
			Name    string      `json:"name"`
			Type    string      `json:"type"`
			Indexed bool        `json:"indexed"`
			Decoded bool        `json:"decoded"`
			Value   interface{} `json:"value"`
		} `json:"params"`
	} `json:"decoded"`
}

type NftData struct {
	TokenID           string          `json:"token_id"`
	TokenBalance      string          `json:"token_balance"`
	TokenUrl          string          `json:"token_url"`
	SupportsErc       []string        `json:"supports_erc"`
	TokenPriceWei     interface{}     `json:"token_price_wei"`
	TokenQuoteRateEth interface{}     `json:"token_quote_rate_eth"`
	Owner             string          `json:"owner"`
	ExternalData      NFTExternalData `json:"external_data"`
}

type NFTExternalData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ExternalUrl string `json:"external_url"`
	Attributes  []struct {
		TraitType string      `json:"trait_type"`
		Value     interface{} `json:"value"`
	} `json:"attributes"`
	Owner string `json:"owner"`
}
