package class_a

import (
	"github.com/google/go-cmp/cmp"
	"github.com/zaebee/govalent/client"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClassA_GetTokenBalances_WhenValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		response string
		want     Portfolio
	}{
		{
			desc:     "WhenEmptyResponse",
			response: `{"data": null}`,
			want:     Portfolio{},
		},
		{
			desc:     "WhenNoNFT",
			response: `{"data":{"address":"0x01","updated_at":"2021-05-05T18:00:00Z","quote_currency":"USD","chain_id":56,"items":[{"contract_name":"Binance Coin","contract_ticker_symbol":"BNB","contract_address":"0xb1","type":"cryptocurrency","balance":"10","quote_rate":645,"quote":0.02}]}}`,
			want: Portfolio{
				Address:       "0x01",
				UpdatedAt:     time.Date(2021, 5, 5, 18, 0, 0, 0, time.UTC),
				QuoteCurrency: "USD",
				ChainID:       56,
				Items: []Item{
					{
						ContractDecimals:     0,
						ContractName:         "Binance Coin",
						ContractTickerSymbol: "BNB",
						ContractAddress:      "0xb1",
						SupportsErc:          nil,
						LogoUrl:              "",
						Type:                 "cryptocurrency",
						Balance:              "10",
						QuoteRate:            645,
						Quote:                0.02,
						NftData:              nil,
					},
				},
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_, err := io.WriteString(w, tc.response)
				if err != nil {
					t.Fatalf("while write response got err: %v", err)
				}
			}))
			defer s.Close()
			api := client.New(s.URL, "ckey_test", s.Client())
			classA := Client{API: *api}
			got, err := classA.GetTokenBalances("chain", "address", BalanceParams{Nft: true})

			if diff := cmp.Diff(got, tc.want); diff != "" || err != nil {
				t.Errorf("%v.GetTokenBalances(chain, address) has diff (-got/+want)\n: %v", classA, diff)
				t.Errorf("%v.GetTokenBalances(chain, address) got err: %v want nil", classA, err)
			}
		})
	}
}
