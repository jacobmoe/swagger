package swagger

import (
	"fmt"
	"reflect"
)

type LICENSE_FIELD int64

const (
	LICENSE_FIELD_NAME LICENSE_FIELD = 0
	LICENSE_FIELD_URL  LICENSE_FIELD = 1
)

func (p LICENSE_FIELD) Required() bool {
	switch p {
	case LICENSE_FIELD_NAME:
		return true
	}
	return false
}

func (p LICENSE_FIELD) String() string {
	switch p {
	case LICENSE_FIELD_NAME:
		return "name"
	case LICENSE_FIELD_URL:
		return "url"
	}
	return "<UNSET>"
}

func (p LICENSE_FIELD) IsTypeOf(object interface{}) bool {
	return reflect.TypeOf(object) == p.Type()
}

func (p LICENSE_FIELD) Type() reflect.Type {
	return reflect.TypeOf("")
}

func LICENSE_FIELDFromString(s string) (LICENSE_FIELD, error) {
	switch s {
	case "name":
		return LICENSE_FIELD_NAME, nil
	case "url":
		return LICENSE_FIELD_URL, nil
	}
	return LICENSE_FIELD(0), fmt.Errorf("not a valid LICENSE_FIELD string: " + s)
}
