package swagger

import "fmt"

type DATA_TYPE_FORMAT int64

const (
	DATA_TYPE_FORMAT_INT32     DATA_TYPE_FORMAT = 0
	DATA_TYPE_FORMAT_INT64     DATA_TYPE_FORMAT = 1
	DATA_TYPE_FORMAT_FLOAT     DATA_TYPE_FORMAT = 2
	DATA_TYPE_FORMAT_DOUBLE    DATA_TYPE_FORMAT = 3
	DATA_TYPE_FORMAT_BYTE      DATA_TYPE_FORMAT = 4
	DATA_TYPE_FORMAT_DATE      DATA_TYPE_FORMAT = 5
	DATA_TYPE_FORMAT_DATE_TIME DATA_TYPE_FORMAT = 6
)

func (p DATA_TYPE_FORMAT) String() string {
	switch p {
	case DATA_TYPE_FORMAT_INT32:
		return "int32"
	case DATA_TYPE_FORMAT_INT64:
		return "int64"
	case DATA_TYPE_FORMAT_FLOAT:
		return "float"
	case DATA_TYPE_FORMAT_DOUBLE:
		return "double"
	case DATA_TYPE_FORMAT_BYTE:
		return "byte"
	case DATA_TYPE_FORMAT_DATE:
		return "date"
	case DATA_TYPE_FORMAT_DATE_TIME:
		return "date-time"
	}
	return "<UNSET>"
}

func DATA_TYPE_FORMATFromString(s string) (DATA_TYPE_FORMAT, error) {
	switch s {
	case "int32":
		return DATA_TYPE_FORMAT_INT32, nil
	case "int64":
		return DATA_TYPE_FORMAT_INT64, nil
	case "float":
		return DATA_TYPE_FORMAT_FLOAT, nil
	case "double":
		return DATA_TYPE_FORMAT_DOUBLE, nil
	case "byte":
		return DATA_TYPE_FORMAT_BYTE, nil
	case "date":
		return DATA_TYPE_FORMAT_DATE, nil
	case "date-time":
		return DATA_TYPE_FORMAT_DATE_TIME, nil
	}
	return DATA_TYPE_FORMAT(0), fmt.Errorf("not a valid DATA_TYPE_FORMAT string")
}
