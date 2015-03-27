package swagger

import "fmt"

type PARAMETER_FIELD int64

const (
	PARAMETER_FIELD_NAME              PARAMETER_FIELD = 0
	PARAMETER_FIELD_IN                PARAMETER_FIELD = 1
	PARAMETER_FIELD_DESCRIPTION       PARAMETER_FIELD = 2
	PARAMETER_FIELD_REQUIRED          PARAMETER_FIELD = 3
	PARAMETER_FIELD_SCHEMA            PARAMETER_FIELD = 4
	PARAMETER_FIELD_TYPE              PARAMETER_FIELD = 5
	PARAMETER_FIELD_FORMAT            PARAMETER_FIELD = 6
	PARAMETER_FIELD_ITEMS             PARAMETER_FIELD = 7
	PARAMETER_FIELD_COLLECTION_FORMAT PARAMETER_FIELD = 8
	PARAMETER_FIELD_DEFAULT           PARAMETER_FIELD = 9
	PARAMETER_FIELD_MAXIMUM           PARAMETER_FIELD = 10
	PARAMETER_FIELD_EXCLUSIVE_MAXIMUM PARAMETER_FIELD = 11
	PARAMETER_FIELD_MINIMUM           PARAMETER_FIELD = 12
	PARAMETER_FIELD_EXCLUSIVE_MINIMUM PARAMETER_FIELD = 13
	PARAMETER_FIELD_MAX_LENGTH        PARAMETER_FIELD = 14
	PARAMETER_FIELD_MIN_LENGTH        PARAMETER_FIELD = 15
	PARAMETER_FIELD_PATTERN           PARAMETER_FIELD = 16
	PARAMETER_FIELD_MAX_ITEMS         PARAMETER_FIELD = 17
	PARAMETER_FIELD_MIN_ITEMS         PARAMETER_FIELD = 18
	PARAMETER_FIELD_UNIQUE_ITEMS      PARAMETER_FIELD = 19
	PARAMETER_FIELD_ENUM              PARAMETER_FIELD = 20
	PARAMETER_FIELD_MULTIPLE_OF       PARAMETER_FIELD = 21
)

func (p PARAMETER_FIELD) String() string {
	switch p {
	case PARAMETER_FIELD_NAME:
		return "name"
	case PARAMETER_FIELD_IN:
		return "in"
	case PARAMETER_FIELD_DESCRIPTION:
		return "description"
	case PARAMETER_FIELD_REQUIRED:
		return "required"
	case PARAMETER_FIELD_SCHEMA:
		return "schema"
	case PARAMETER_FIELD_TYPE:
		return "type"
	case PARAMETER_FIELD_FORMAT:
		return "format"
	case PARAMETER_FIELD_ITEMS:
		return "items"
	case PARAMETER_FIELD_COLLECTION_FORMAT:
		return "collectionFormat"
	case PARAMETER_FIELD_DEFAULT:
		return "default"
	case PARAMETER_FIELD_MAXIMUM:
		return "maximum"
	case PARAMETER_FIELD_EXCLUSIVE_MAXIMUM:
		return "exclusiveMaximum"
	case PARAMETER_FIELD_MINIMUM:
		return "minimum"
	case PARAMETER_FIELD_EXCLUSIVE_MINIMUM:
		return "exclusiveMinimum"
	case PARAMETER_FIELD_MAX_LENGTH:
		return "maxLength"
	case PARAMETER_FIELD_MIN_LENGTH:
		return "minLength"
	case PARAMETER_FIELD_PATTERN:
		return "pattern"
	case PARAMETER_FIELD_MAX_ITEMS:
		return "maxItems"
	case PARAMETER_FIELD_MIN_ITEMS:
		return "minItems"
	case PARAMETER_FIELD_UNIQUE_ITEMS:
		return "uniqueItems"
	case PARAMETER_FIELD_ENUM:
		return "enum"
	case PARAMETER_FIELD_MULTIPLE_OF:
		return "multipleOf"
	}
	return "<UNSET>"
}

func PARAMETER_FIELDFromString(s string) (PARAMETER_FIELD, error) {
	switch s {
	case "name":
		return PARAMETER_FIELD_NAME, nil
	case "in":
		return PARAMETER_FIELD_IN, nil
	case "description":
		return PARAMETER_FIELD_DESCRIPTION, nil
	case "required":
		return PARAMETER_FIELD_REQUIRED, nil
	case "schema":
		return PARAMETER_FIELD_SCHEMA, nil
	case "type":
		return PARAMETER_FIELD_TYPE, nil
	case "format":
		return PARAMETER_FIELD_FORMAT, nil
	case "items":
		return PARAMETER_FIELD_ITEMS, nil
	case "collectionFormat":
		return PARAMETER_FIELD_COLLECTION_FORMAT, nil
	case "default":
		return PARAMETER_FIELD_DEFAULT, nil
	case "maximum":
		return PARAMETER_FIELD_MAXIMUM, nil
	case "exclusiveMaximum":
		return PARAMETER_FIELD_EXCLUSIVE_MAXIMUM, nil
	case "minimum":
		return PARAMETER_FIELD_MINIMUM, nil
	case "exclusiveMinimum":
		return PARAMETER_FIELD_EXCLUSIVE_MINIMUM, nil
	case "maxLength":
		return PARAMETER_FIELD_MAX_LENGTH, nil
	case "minLength":
		return PARAMETER_FIELD_MIN_LENGTH, nil
	case "pattern":
		return PARAMETER_FIELD_PATTERN, nil
	case "maxItems":
		return PARAMETER_FIELD_MAX_ITEMS, nil
	case "minItems":
		return PARAMETER_FIELD_MIN_ITEMS, nil
	case "uniqueItems":
		return PARAMETER_FIELD_UNIQUE_ITEMS, nil
	case "enum":
		return PARAMETER_FIELD_ENUM, nil
	case "multipleOf":
		return PARAMETER_FIELD_MULTIPLE_OF, nil
	}
	return PARAMETER_FIELD(0), fmt.Errorf("not a valid PARAMETER_FIELD string")
}
