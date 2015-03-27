package swagger

import (
	"fmt"
	"reflect"
)

type EXTERNAL_DOC_FIELD int64

const (
	EXTERNAL_DOC_FIELD_DESCRIPTION EXTERNAL_DOC_FIELD = 0
	EXTERNAL_DOC_FIELD_URL         EXTERNAL_DOC_FIELD = 1
	EXTERNAL_DOC_FIELD_EMAIL       EXTERNAL_DOC_FIELD = 2
)

func (p EXTERNAL_DOC_FIELD) Required() bool {
	switch p {
	case EXTERNAL_DOC_FIELD_URL:
		return true
	}
	return false
}

func (p EXTERNAL_DOC_FIELD) String() string {
	switch p {
	case EXTERNAL_DOC_FIELD_DESCRIPTION:
		return "description"
	case EXTERNAL_DOC_FIELD_URL:
		return "url"
	}
	return "<UNSET>"
}

func (p EXTERNAL_DOC_FIELD) IsTypeOf(object interface{}) bool {
	return reflect.TypeOf(object) == p.Type()
}

func (p EXTERNAL_DOC_FIELD) Type() reflect.Type {
	return reflect.TypeOf("")
}

func EXTERNAL_DOC_FIELDFromString(s string) (EXTERNAL_DOC_FIELD, error) {
	switch s {
	case "description":
		return EXTERNAL_DOC_FIELD_DESCRIPTION, nil
	case "url":
		return EXTERNAL_DOC_FIELD_URL, nil
	}
	return EXTERNAL_DOC_FIELD(0), fmt.Errorf("not a valid EXTERNAL_DOC_FIELD string: " + s)
}
