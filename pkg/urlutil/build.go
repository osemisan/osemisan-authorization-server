package urlutil

import (
	"fmt"
	"net/url"
)

type BuildURLError struct {
	msg string
	err error
}

func (e *BuildURLError) Error() string {
	return fmt.Sprintf("cannot build URL: %s (%s)", e.msg, e.err.Error())
}

func (e *BuildURLError) Unwrap() error {
	return e.err
}

func BuildURL(b string, qMap map[string]string) (string, error) {
	u, err := url.Parse(b)
	if err != nil {
		return "", &BuildURLError{msg: "parse URL", err: err}
	}
	q := u.Query()
	for k, v := range qMap {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
