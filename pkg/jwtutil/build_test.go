package jwtutil_test

import (
	"testing"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/osemisan-authorization-server/pkg/jwtutil"
)

const symKey = "hoge"

func TestBuildJWT(t *testing.T) {
	tests := []struct {
		name  string
		scope string
		wantScope string
	}{
		{
			"JWTにサインしてから検証しても同じスコープになる",
			"abura kuma",
			"abura kuma",
		},
	}
	key, err := jwk.FromRaw([]byte(symKey))
	if err != nil {
		t.Error("Failed to create key")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := jwtutil.BuildJWT(tt.scope)
			if err != nil {
				t.Error("Failed to build JWT", err)
				return
			}
			verifiedTok, err := jwt.Parse([]byte(got), jwt.WithKey(jwa.HS256, key))
			if err != nil {
				t.Error("Failed to parse got token bytes")
				return
			}
			scope, exists := verifiedTok.Get("scope")
			if !exists {
				t.Error(`"scope" does not exists in the verified token`)
				return
			}
			if scope != tt.wantScope {
				t.Errorf("Unexpected scope, expected: %s, actual: %s", tt.wantScope, scope)
			}
		})
	}
}