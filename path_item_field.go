package swagger

import "fmt"

type PATH_ITEM_FIELD int64

const (
	PATH_ITEM_FIELD_REF        PATH_ITEM_FIELD = 0
	PATH_ITEM_FIELD_GET        PATH_ITEM_FIELD = 1
	PATH_ITEM_FIELD_PUT        PATH_ITEM_FIELD = 2
	PATH_ITEM_FIELD_POST       PATH_ITEM_FIELD = 3
	PATH_ITEM_FIELD_DELETE     PATH_ITEM_FIELD = 4
	PATH_ITEM_FIELD_OPTIONS    PATH_ITEM_FIELD = 5
	PATH_ITEM_FIELD_HEAD       PATH_ITEM_FIELD = 6
	PATH_ITEM_FIELD_PATCH      PATH_ITEM_FIELD = 7
	PATH_ITEM_FIELD_PARAMETERS PATH_ITEM_FIELD = 8
)

func (p PATH_ITEM_FIELD) String() string {
	switch p {
	case PATH_ITEM_FIELD_REF:
		return "$ref"
	case PATH_ITEM_FIELD_GET:
		return "get"
	case PATH_ITEM_FIELD_PUT:
		return "put"
	case PATH_ITEM_FIELD_POST:
		return "post"
	case PATH_ITEM_FIELD_DELETE:
		return "delete"
	case PATH_ITEM_FIELD_OPTIONS:
		return "options"
	case PATH_ITEM_FIELD_HEAD:
		return "head"
	case PATH_ITEM_FIELD_PATCH:
		return "patch"
	case PATH_ITEM_FIELD_PARAMETERS:
		return "parameters"
	}
	return "<UNSET>"
}

func PATH_ITEM_FIELDFromString(s string) (PATH_ITEM_FIELD, error) {
	switch s {
	case "$ref":
		return PATH_ITEM_FIELD_REF, nil
	case "get":
		return PATH_ITEM_FIELD_GET, nil
	case "put":
		return PATH_ITEM_FIELD_PUT, nil
	case "post":
		return PATH_ITEM_FIELD_POST, nil
	case "delete":
		return PATH_ITEM_FIELD_DELETE, nil
	case "options":
		return PATH_ITEM_FIELD_OPTIONS, nil
	case "head":
		return PATH_ITEM_FIELD_HEAD, nil
	case "patch":
		return PATH_ITEM_FIELD_PATCH, nil
	case "parameters":
		return PATH_ITEM_FIELD_PARAMETERS, nil
	}
	return PATH_ITEM_FIELD(0), fmt.Errorf("not a valid PATH_ITEM_FIELD string:" + s)
}
