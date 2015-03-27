package swagger

import (
	"encoding/json"
	"fmt"
)

type PARAMETER_TYPE int64

const (
	PARAMETER_TYPE_PATH      PARAMETER_TYPE = 0
	PARAMETER_TYPE_QUERY     PARAMETER_TYPE = 1
	PARAMETER_TYPE_HEADER    PARAMETER_TYPE = 2
	PARAMETER_TYPE_BODY      PARAMETER_TYPE = 3
	PARAMETER_TYPE_FORM_DATA PARAMETER_TYPE = 4
)

func (p PARAMETER_TYPE) String() string {
	switch p {
	case PARAMETER_TYPE_PATH:
		return "path"
	case PARAMETER_TYPE_QUERY:
		return "query"
	case PARAMETER_TYPE_HEADER:
		return "header"
	case PARAMETER_TYPE_BODY:
		return "body"
	case PARAMETER_TYPE_FORM_DATA:
		return "formData"
	}
	return "<UNSET>"
}

func PARAMETER_TYPEFromString(s string) (PARAMETER_TYPE, error) {
	switch s {
	case "path":
		return PARAMETER_TYPE_PATH, nil
	case "query":
		return PARAMETER_TYPE_QUERY, nil
	case "header":
		return PARAMETER_TYPE_HEADER, nil
	case "body":
		return PARAMETER_TYPE_BODY, nil
	case "formData":
		return PARAMETER_TYPE_FORM_DATA, nil
	}
	return PARAMETER_TYPE(0), fmt.Errorf("not a valid PARAMETER_TYPE string: " + s)
}

func (p PARAMETER_TYPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
