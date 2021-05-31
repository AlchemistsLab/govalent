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

// Pagination returns pagination data for each endpoint.
type Pagination struct {
	HasMore    bool        `json:"has_more"`
	PageNumber int         `json:"page_number"`
	PageSize   int         `json:"page_size"`
	TotalCount interface{} `json:"total_count"`
}

// Balance returns portfolio data for token balances response.
type Balance struct {
	Data         PortfolioData `json:"data"`
	Error        bool          `json:"error"`
	ErrorMessage interface{}   `json:"error_message"`
	ErrorCode    interface{}   `json:"error_code"`
}

// PortfolioData returns list of items for portfolio endpoint.
type PortfolioData struct {
	Address       string          `json:"address"`
	UpdatedAt     time.Time       `json:"updated_at"`
	NextUpdateAt  time.Time       `json:"next_update_at"`
	QuoteCurrency string          `json:"quote_currency"`
	ChainID       int             `json:"chain_id"`
	Items         []PortfolioItem `json:"items"`
	Pagination    Pagination      `json:"pagination"`
}

type PortfolioItem struct {
	ContractDecimals     int       `json:"contract_decimals"`
	ContractName         string    `json:"contract_name"`
	ContractTickerSymbol string    `json:"contract_ticker_symbol"`
	ContractAddress      string    `json:"contract_address"`
	SupportsErc          []string  `json:"supports_erc"`
	LogoUrl              string    `json:"logo_url"`
	Type                 string    `json:"type"`
	Balance              string    `json:"balance"`
	QuoteRate            float64   `json:"quote_rate"`
	Quote                float64   `json:"quote"`
	NftData              []NftData `json:"nft_data"`
	Holdings             []Holding `json:"holdings"`
}

type Holding struct {
	Timestamp time.Time `json:"timestamp"`
	QuoteRate float64   `json:"quote_rate"`
	Open      struct {
		Balance string  `json:"balance"`
		Quote   float64 `json:"quote"`
	} `json:"open"`
	High struct {
		Balance string  `json:"balance"`
		Quote   float64 `json:"quote"`
	} `json:"high"`
	Low struct {
		Balance string  `json:"balance"`
		Quote   float64 `json:"quote"`
	} `json:"low"`
	Close struct {
		Balance string  `json:"balance"`
		Quote   float64 `json:"quote"`
	} `json:"close"`
}

type NftData struct {
	TokenID           string       `json:"token_id"`
	TokenBalance      string       `json:"token_balance"`
	TokenUrl          string       `json:"token_url"`
	SupportsErc       []string     `json:"supports_erc"`
	TokenPriceWei     interface{}  `json:"token_price_wei"`
	TokenQuoteRateEth interface{}  `json:"token_quote_rate_eth"`
	Owner             string       `json:"owner"`
	ExternalData      ExternalData `json:"external_data"`
}

type ExternalData struct {
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

type Transaction struct {
	Data         TransactionData `json:"data"`
	Error        bool            `json:"error"`
	ErrorMessage interface{}     `json:"error_message"`
	ErrorCode    interface{}     `json:"error_code"`
}

type TransactionData struct {
	Address       string            `json:"address"`
	UpdatedAt     time.Time         `json:"updated_at"`
	NextUpdateAt  time.Time         `json:"next_update_at"`
	QuoteCurrency string            `json:"quote_currency"`
	ChainId       int               `json:"chain_id"`
	Items         []TransactionItem `json:"items"`
	Pagination    Pagination        `json:"pagination"`
}

type TransactionItem struct {
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
	LogEvents        []struct {
		BlockSignedAt      time.Time   `json:"block_signed_at"`
		BlockHeight        int         `json:"block_height"`
		TxOffset           int         `json:"tx_offset"`
		LogOffset          int         `json:"log_offset"`
		TxHash             string      `json:"tx_hash"`
		RawLogTopicsBytes  interface{} `json:"_raw_log_topics_bytes"`
		RawLogTopics       []string    `json:"raw_log_topics"`
		SenderAddress      string      `json:"sender_address"`
		SenderAddressLabel interface{} `json:"sender_address_label"`
		RawLogData         string      `json:"raw_log_data"`
		Decoded            struct {
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
	} `json:"log_events"`
}

type Block struct {
	Data         BlockData   `json:"data"`
	Error        bool        `json:"error"`
	ErrorMessage interface{} `json:"error_message"`
	ErrorCode    interface{} `json:"error_code"`
}

type BlockData struct {
	UpdatedAt time.Time `json:"updated_at"`
	Items     []struct {
		SignedAt time.Time `json:"signed_at"`
		Height   int       `json:"height"`
	} `json:"items"`
	Pagination Pagination `json:"pagination"`
}
