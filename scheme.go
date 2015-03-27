package swagger

import "fmt"

type SCHEME int64

const (
	SCHEME_HTTP  SCHEME = 0
	SCHEME_HTTPS SCHEME = 1
)

func (p SCHEME) String() string {
	switch p {
	case SCHEME_HTTP:
		return "http"
	case SCHEME_HTTPS:
		return "https"
	}
	return "<UNSET>"
}

func SCHEMEFromString(s string) (SCHEME, error) {
	switch s {
	case "http":
		return SCHEME_HTTP, nil
	case "https":
		return SCHEME_HTTPS, nil
	}
	return SCHEME(0), fmt.Errorf("not a valid SCHEME string: " + s)
}
