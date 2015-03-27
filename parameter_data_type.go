package swagger

import (
	"encoding/json"
	"fmt"
)

type PARAMETER_DATA_TYPE int64

const (
	PARAMETER_DATA_TYPE_INTEGER PARAMETER_DATA_TYPE = 0
	PARAMETER_DATA_TYPE_STRING  PARAMETER_DATA_TYPE = 1
	PARAMETER_DATA_TYPE_NUMBER  PARAMETER_DATA_TYPE = 2
	PARAMETER_DATA_TYPE_BOOLEAN PARAMETER_DATA_TYPE = 3
	PARAMETER_DATA_TYPE_ARRAY   PARAMETER_DATA_TYPE = 4
	PARAMETER_DATA_TYPE_FILE    PARAMETER_DATA_TYPE = 5
)

func (p PARAMETER_DATA_TYPE) String() string {
	switch p {
	case PARAMETER_DATA_TYPE_INTEGER:
		return "integer"
	case PARAMETER_DATA_TYPE_STRING:
		return "string"
	case PARAMETER_DATA_TYPE_NUMBER:
		return "number"
	case PARAMETER_DATA_TYPE_BOOLEAN:
		return "boolean"
	case PARAMETER_DATA_TYPE_ARRAY:
		return "array"
	case PARAMETER_DATA_TYPE_FILE:
		return "file"
	}
	return "<UNSET>"
}

func PARAMETER_DATA_TYPEFromString(s string) (PARAMETER_DATA_TYPE, error) {
	switch s {
	case "integer":
		return PARAMETER_DATA_TYPE_INTEGER, nil
	case "string":
		return PARAMETER_DATA_TYPE_STRING, nil
	case "number":
		return PARAMETER_DATA_TYPE_NUMBER, nil
	case "boolean":
		return PARAMETER_DATA_TYPE_BOOLEAN, nil
	case "array":
		return PARAMETER_DATA_TYPE_ARRAY, nil
	case "file":
		return PARAMETER_DATA_TYPE_FILE, nil
	}
	return PARAMETER_DATA_TYPE(0), fmt.Errorf("not a valid PARAMETER_DATA_TYPE string: " + s)
}

func (p PARAMETER_DATA_TYPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
