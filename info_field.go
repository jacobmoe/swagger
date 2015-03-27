package swagger

import "fmt"

type INFO_FIELD int64

const (
	INFO_FIELD_TITLE            INFO_FIELD = 0
	INFO_FIELD_DESCRIPTION      INFO_FIELD = 1
	INFO_FIELD_TERMS_OF_SERVICE INFO_FIELD = 2
	INFO_FIELD_CONTACT          INFO_FIELD = 3
	INFO_FIELD_LICENSE          INFO_FIELD = 4
	INFO_FIELD_VERSION          INFO_FIELD = 5
)

func (p INFO_FIELD) Required() bool {
	switch p {
	case INFO_FIELD_TITLE, INFO_FIELD_VERSION:
		return true
	default:
		return false
	}
}

func (p INFO_FIELD) String() string {
	switch p {
	case INFO_FIELD_TITLE:
		return "title"
	case INFO_FIELD_DESCRIPTION:
		return "description"
	case INFO_FIELD_TERMS_OF_SERVICE:
		return "termsOfService"
	case INFO_FIELD_CONTACT:
		return "contact"
	case INFO_FIELD_LICENSE:
		return "license"
	case INFO_FIELD_VERSION:
		return "version"
	}
	return "<UNSET>"
}

func INFO_FIELDFromString(s string) (INFO_FIELD, error) {
	switch s {
	case "title":
		return INFO_FIELD_TITLE, nil
	case "description":
		return INFO_FIELD_DESCRIPTION, nil
	case "termsOfService":
		return INFO_FIELD_TERMS_OF_SERVICE, nil
	case "contact":
		return INFO_FIELD_CONTACT, nil
	case "license":
		return INFO_FIELD_LICENSE, nil
	case "version":
		return INFO_FIELD_VERSION, nil
	}
	return INFO_FIELD(0), fmt.Errorf("not a valid INFO_FIELD string: " + s)
}
