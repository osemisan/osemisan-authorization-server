package url_test

import (
	"testing"

	"github.com/osemisan-authorization-server/pkg/url"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		name string
		base string
		qMap map[string]string;
		wantURL string
	}{
		{
			name: "baseだけ渡したらそのまま帰ってくる",
			base: "http://example.com",
			qMap: map[string]string{},
			wantURL:"http://example.com",
		},
		{
			name: "クエリをmapで渡すとその通りに付与される",
			base: "http://example.com",
			qMap: map[string]string{
				"foo": "1",
				"bar": "2",
			},
			wantURL: "http://example.com?bar=2&foo=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := url.BuildURL(tt.base, tt.qMap)
			if err != nil {
				t.Error("Failed to build URL", err)
			}
			if a != tt.wantURL {
				t.Errorf("Unexpected built URL, expected: %s, actual: %s", tt.wantURL, a)
			}
		})
	}
}