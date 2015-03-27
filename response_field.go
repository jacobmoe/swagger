package swagger

import "fmt"

type RESPONSE_FIELD int64

const (
	RESPONSE_FIELD_DESCRIPTION RESPONSE_FIELD = 0
	RESPONSE_FIELD_SCHEMA      RESPONSE_FIELD = 1
	RESPONSE_FIELD_HEADERS     RESPONSE_FIELD = 2
	RESPONSE_FIELD_EXAMPLES    RESPONSE_FIELD = 3
)

func (p RESPONSE_FIELD) String() string {
	switch p {
	case RESPONSE_FIELD_DESCRIPTION:
		return "description"
	case RESPONSE_FIELD_SCHEMA:
		return "schema"
	case RESPONSE_FIELD_HEADERS:
		return "headers"
	case RESPONSE_FIELD_EXAMPLES:
		return "examples"
	}
	return "<UNSET>"
}

func RESPONSE_FIELDFromString(s string) (RESPONSE_FIELD, error) {
	switch s {
	case "description":
		return RESPONSE_FIELD_DESCRIPTION, nil
	case "schema":
		return RESPONSE_FIELD_SCHEMA, nil
	case "headers":
		return RESPONSE_FIELD_HEADERS, nil
	case "examples":
		return RESPONSE_FIELD_EXAMPLES, nil
	}
	return RESPONSE_FIELD(0), fmt.Errorf("not a valid RESPONSE_FIELD string")
}
