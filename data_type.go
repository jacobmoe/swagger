package swagger

import "fmt"

type DATA_TYPE int64

const (
	DATA_TYPE_INTEGER DATA_TYPE = 0
	DATA_TYPE_STRING  DATA_TYPE = 1
	DATA_TYPE_NUMBER  DATA_TYPE = 2
	DATA_TYPE_BOOLEAN DATA_TYPE = 3
	DATA_TYPE_ARRAY   DATA_TYPE = 4
)

func (p DATA_TYPE) String() string {
	switch p {
	case DATA_TYPE_INTEGER:
		return "integer"
	case DATA_TYPE_STRING:
		return "string"
	case DATA_TYPE_NUMBER:
		return "number"
	case DATA_TYPE_BOOLEAN:
		return "boolean"
	case DATA_TYPE_ARRAY:
		return "array"
	}
	return "<UNSET>"
}

func DATA_TYPEFromString(s string) (DATA_TYPE, error) {
	switch s {
	case "integer":
		return DATA_TYPE_INTEGER, nil
	case "string":
		return DATA_TYPE_STRING, nil
	case "number":
		return DATA_TYPE_NUMBER, nil
	case "boolean":
		return DATA_TYPE_BOOLEAN, nil
	case "array":
		return DATA_TYPE_ARRAY, nil
	}
	return DATA_TYPE(0), fmt.Errorf("not a valid DATA_TYPE string: " + s)
}
