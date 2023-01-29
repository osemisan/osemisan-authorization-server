package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/osemisan-authorization-server/pkg/handlers"
)

func GetTestHandler() http.HandlerFunc {
	fn := handlers.RootHandler
	return fn
}

func TestRootHandler(t *testing.T) {
	s := httptest.NewServer(GetTestHandler())
	defer s.Close()

	c := new(http.Client)

	tests := []struct {
		name string
		wantStatusCode int
	}{
		{
			"リクエストしたらステータスコード200",
			http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, s.URL, nil)
			if err != nil {
				t.Error("Failed to create new request", err)
				return
			}

			res, err := c.Do(req)
			if err != nil {
				t.Error("Failed to request", err)
				return
			}

			if res.StatusCode != tt.wantStatusCode {
				t.Errorf("Unexpected status code, expected %d, actual %d", tt.wantStatusCode, res.StatusCode)
			}
		})
	}
}
