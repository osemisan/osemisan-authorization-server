package handlers_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/osemisan-authorization-server/pkg/handlers"
)

func TestAuthorizeHandler(t *testing.T) {
	tests := []struct {
		name string
		query url.Values
		wantTextInResHTML string
		wantStatusCode int
	}{
		{
			name: "不正なクライアントIDを渡すと、エラーページが表示される",
			query: url.Values{
				"client_id": {"invalid_client_id"},
			},
			wantTextInResHTML: "Unknown client",
			wantStatusCode: http.StatusOK,
		},
		{
			name: "不正なリダイレクトURIを渡すと、エラーページが表示される",
			query: url.Values{
				"client_id": {"osemisan-client-id-1"},
				"redirect_uri": {"invalid_redirect_uri"},
			},
			wantTextInResHTML: "Invalid redirect URI",
			wantStatusCode: http.StatusOK,
		},
		{
			name: "生成なクエリパラメータを渡すと、許可ページが表示される",
			query: url.Values{
				"client_id": {"osemisan-client-id-1"},
				"redirect_uri": {"http://localhost:9000/callback"},
			},
			wantTextInResHTML: "このクライアントを許可しますか？",
			wantStatusCode: http.StatusOK,
		},
	}

	c := new(http.Client)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(handlers.AuthorizeHandler))
			u, err := url.Parse(s.URL)
			if err != nil {
				t.Error("Failed to parse URL", err)
			}
			u.RawQuery = tt.query.Encode()

			req, err := http.NewRequest(http.MethodGet, u.String(), nil)
			if err != nil {
				t.Error("Failed to create new request", err)
				return
			}

			res, err := c.Do(req)
			if err != nil {
				t.Error("Failed to request")
				return
			}
			defer res.Body.Close()

			if tt.wantStatusCode != res.StatusCode {
				t.Errorf("Unexpected statuc code, expected: %d, actual: %d", tt.wantStatusCode, res.StatusCode)
				return
			}

			body, _ := ioutil.ReadAll(res.Body)
			buf := bytes.NewBuffer(body)
			html := buf.String()

			if !strings.Contains(html, tt.wantTextInResHTML) {
				t.Errorf(`Response HTML does not contain "%s"`, tt.wantTextInResHTML)
				return
			}
		})
	}
}
