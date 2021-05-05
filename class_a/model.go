package class_a

import "time"

type BalanceParams struct {
	Nft        bool `json:"nft"`
	NoNftFetch bool `json:"no-nft-fetch"`
}

type Balance struct {
	Data         Portfolio   `json:"data"`
	Error        bool        `json:"error"`
	ErrorMessage interface{} `json:"error_message"`
	ErrorCode    interface{} `json:"error_code"`
}

type Portfolio struct {
	Address       string      `json:"address"`
	UpdatedAt     time.Time   `json:"updated_at"`
	NextUpdateAt  time.Time   `json:"next_update_at"`
	QuoteCurrency string      `json:"quote_currency"`
	ChainID       int         `json:"chain_id"`
	Items         []Item      `json:"items"`
	Pagination    interface{} `json:"pagination"`
}

type Item struct {
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
	TokenID           string      `json:"token_id"`
	TokenBalance      string      `json:"token_balance"`
	TokenUrl          string      `json:"token_url"`
	SupportsErc       []string    `json:"supports_erc"`
	TokenPriceWei     interface{} `json:"token_price_wei"`
	TokenQuoteRateEth interface{} `json:"token_quote_rate_eth"`
	ExternalData      *struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
		ExternalUrl string `json:"external_url"`
		Attributes  []struct {
			TraitType string      `json:"trait_type"`
			Value     interface{} `json:"value"`
		} `json:"attributes"`
		Owner string `json:"owner"`
	} `json:"external_data"`
	Owner string `json:"owner"`
}
