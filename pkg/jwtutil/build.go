package jwtutil

import (
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

const symKey = "hoge"

type BuildJWTError struct {
	msg string
	err error
}

func (e *BuildJWTError) Error() string {
	return fmt.Sprintf("cannto build JWT: %s (%s)", e.msg, e.err.Error())
}

func (e *BuildJWTError) Unwrap() error {
	return e.err
}

func BuildJWT(scope string) (string, error) {
	tok, err := jwt.NewBuilder().
		Issuer(`github.com/osemisan/osemisan-authorization-server`).
		IssuedAt(time.Now()).
		Claim("scope", scope).
		Build()
	if err != nil {
		return "", &BuildJWTError{msg: "build JWT", err: err}
	}
	key, err := jwk.FromRaw([]byte(symKey))
	if err != nil {
		return "", &BuildJWTError{msg: "create key", err: err}
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, key))
	if err != nil {
		return "", &BuildJWTError{msg: "sign", err: err}
	}
	return fmt.Sprintf("%s", signed), nil
}
