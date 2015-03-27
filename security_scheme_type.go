package swagger

import "fmt"

type SECURITY_SCHEME_TYPE int64

const (
	SECURITY_SCHEME_TYPE_BASIC   SECURITY_SCHEME_TYPE = 0
	SECURITY_SCHEME_TYPE_API_KEY SECURITY_SCHEME_TYPE = 1
	SECURITY_SCHEME_TYPE_OAUTH2  SECURITY_SCHEME_TYPE = 2
)

func (p SECURITY_SCHEME_TYPE) String() string {
	switch p {
	case SECURITY_SCHEME_TYPE_BASIC:
		return "basic"
	case SECURITY_SCHEME_TYPE_API_KEY:
		return "apiKey"
	case SECURITY_SCHEME_TYPE_OAUTH2:
		return "oauth2"
	}
	return "<UNSET>"
}

func SECURITY_SCHEME_TYPEFromString(s string) (SECURITY_SCHEME_TYPE, error) {
	switch s {
	case "basic":
		return SECURITY_SCHEME_TYPE_BASIC, nil
	case "apiKey":
		return SECURITY_SCHEME_TYPE_API_KEY, nil
	case "oauth2":
		return SECURITY_SCHEME_TYPE_OAUTH2, nil
	}
	return SECURITY_SCHEME_TYPE(0), fmt.Errorf("not a valid SECURITY_SCHEME_TYPE string: " + s)
}
