package swagger

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type SWAGGER_VERSION int64

const (
	SWAGGER_VERSION_2_0 SWAGGER_VERSION = 0
)

func (v SWAGGER_VERSION) String() string {
	switch v {
	case SWAGGER_VERSION_2_0:
		return "2.0"
	}
	return "<UNSET>"
}

func (v SWAGGER_VERSION) Type() reflect.Type {
	return reflect.TypeOf(v)
}

func (v SWAGGER_VERSION) Slice() reflect.Type {
	return reflect.SliceOf(v.Type())
}

func SWAGGER_VERSIONFromString(s string) (SWAGGER_VERSION, error) {
	switch s {
	case "2.0":
		return SWAGGER_VERSION_2_0, nil
	}
	return SWAGGER_VERSION(0), fmt.Errorf("not a valid SWAGGER_VERSION string: " + s)
}

func (v SWAGGER_VERSION) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.String())
}
