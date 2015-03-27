package swagger

import (
	"fmt"
	"reflect"
)

type CONTACT_FIELD int64

const (
	CONTACT_FIELD_NAME  CONTACT_FIELD = 0
	CONTACT_FIELD_URL   CONTACT_FIELD = 1
	CONTACT_FIELD_EMAIL CONTACT_FIELD = 2
)

func (p CONTACT_FIELD) Required() bool {
	return false
}

func (p CONTACT_FIELD) String() string {
	switch p {
	case CONTACT_FIELD_NAME:
		return "name"
	case CONTACT_FIELD_URL:
		return "url"
	case CONTACT_FIELD_EMAIL:
		return "email"
	}
	return "<UNSET>"
}

func (p CONTACT_FIELD) IsTypeOf(object interface{}) bool {
	return reflect.TypeOf(object) == p.Type()
}

func (p CONTACT_FIELD) Type() reflect.Type {
	return reflect.TypeOf("")
}

func CONTACT_FIELDFromString(s string) (CONTACT_FIELD, error) {
	switch s {
	case "name":
		return CONTACT_FIELD_NAME, nil
	case "url":
		return CONTACT_FIELD_URL, nil
	case "email":
		return CONTACT_FIELD_EMAIL, nil
	}
	return CONTACT_FIELD(0), fmt.Errorf("not a valid CONTACT_FIELD string: " + s)
}
