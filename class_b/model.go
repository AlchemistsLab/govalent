package class_b

import "time"

type SushiSwapActsParams struct {
	Swaps         bool   `json:"swaps"`
	QuoteCurrency string `json:"quote-currency"`
}

// Pagination returns pagination metadata for each endpoint.
type Pagination struct {
	HasMore    bool `json:"has_more"`
	PageNumber int  `json:"page_number"`
	PageSize   int  `json:"page_size"`
	TotalCount int  `json:"total_count"`
}

type Response struct{}

// SushiSwapActsResponse returns response for acts endpoint.
type SushiSwapActsResponse struct {
	Data         SushiSwapActs `json:"data"`
	Error        bool          `json:"error"`
	ErrorMessage string        `json:"error_message"`
	ErrorCode    int           `json:"error_code"`
}

// SushiSwapActs returns actions for sushiswap act endpoint.
type SushiSwapActs struct {
	Address       string    `json:"address"`
	UpdatedAt     time.Time `json:"updated_at"`
	NextUpdateAt  time.Time `json:"next_update_at"`
	QuoteCurrency string    `json:"quote_currency"`
	ChainID       int       `json:"chain_id"`
	Items         []struct {
		ActAt       time.Time `json:"act_at"`
		Act         string    `json:"act"`
		Description string    `json:"description"`
		TxHash      string    `json:"tx_hash"`
		Token0      Token     `json:"token_0"`
		Token1      Token     `json:"token_1"`
		PoolToken   Token     `json:"pool_token"`
	} `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type Token struct {
	ContractDecimals     int     `json:"contract_decimals"`
	ContractTickerSymbol string  `json:"contract_ticker_symbol"`
	ContractAddress      string  `json:"contract_address"`
	LogoUrl              string  `json:"logo_url"`
	Balance              string  `json:"balance"`
	Quote                float64 `json:"quote"`
	QuoteRate            float64 `json:"quote_rate"`
	TotalSupply          string  `json:"total_supply"`
}
