package client

import (
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type out struct {
	Account int
	Address string
}

func TestAPI_Request_WhenValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		response string
		want     *out
	}{
		{
			desc:     "WhenStatusOK",
			response: `{"account": 1, "address": "test_address"}`,
			want: &out{
				Account: 1,
				Address: "test_address",
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
			client := s.Client()
			api := New(s.URL, "ckey_test", client)

			got := &out{}
			err := api.Request("GET", "/v1/", nil, got)

			if diff := cmp.Diff(got, tc.want); diff != "" || err != nil {
				t.Errorf("%v.Request(params) got %v want %v", api, got, tc.want)
				t.Errorf("%v.Request(params) got err: %v want nil", api, err)
			}
		})
	}
}

func TestAPI_Request_WhenInvalid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		desc     string
		response string
		status   int
		want     CovalentError
	}{
		{
			desc:     "WhenStatusForbidden",
			response: `{"code": 401, "msg": "forbidden"}`,
			status:   http.StatusForbidden,
			want: CovalentError{
				Code: 401,
				Msg:  "forbidden",
			},
		},
		{
			desc:     "WhenStatusOKWithInvalidJSON",
			response: `{"code": 200, "key":}`,
			status:   http.StatusOK,
			want: CovalentError{
				Msg: "Invalid JSON",
			},
		},
		{
			desc:     "WhenStatusNotFound",
			response: `{"code": 404, "msg": "url not found"}`,
			status:   http.StatusNotFound,
			want: CovalentError{
				Code: 404,
				Msg:  "url not found",
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				w.Header().Set("Content-Type", "application/json")
				_, err := io.WriteString(w, tc.response)
				if err != nil {
					t.Fatalf("while write response got err: %v", err)
				}
			}))
			defer s.Close()
			client := s.Client()
			api := New(s.URL, "ckey_test", client)

			err := api.Request("GET", "/v1/", nil, &out{})

			if err != tc.want {
				t.Errorf("%v.Request(params) got err: %v want %v", api, err, tc.want)
			}
		})
	}
}
