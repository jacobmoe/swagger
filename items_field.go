package swagger

import "fmt"

type ITEMS_FIELD int64

const (
	ITEMS_FIELD_TYPE              ITEMS_FIELD = 0
	ITEMS_FIELD_FORMAT            ITEMS_FIELD = 1
	ITEMS_FIELD_ITEMS             ITEMS_FIELD = 2
	ITEMS_FIELD_COLLECTION_FORMAT ITEMS_FIELD = 3
	ITEMS_FIELD_DEFAULT           ITEMS_FIELD = 4
	ITEMS_FIELD_MAXIMUM           ITEMS_FIELD = 5
	ITEMS_FIELD_EXCLUSIVE_MAXIMUM ITEMS_FIELD = 6
	ITEMS_FIELD_MINIMUM           ITEMS_FIELD = 7
	ITEMS_FIELD_EXCLUSIVE_MINIMUM ITEMS_FIELD = 8
	ITEMS_FIELD_MAX_LENGTH        ITEMS_FIELD = 9
	ITEMS_FIELD_MIN_LENGTH        ITEMS_FIELD = 10
	ITEMS_FIELD_PATTERN           ITEMS_FIELD = 11
	ITEMS_FIELD_MAX_ITEMS         ITEMS_FIELD = 12
	ITEMS_FIELD_MIN_ITEMS         ITEMS_FIELD = 13
	ITEMS_FIELD_UNIQUE_ITEMS      ITEMS_FIELD = 14
	ITEMS_FIELD_ENUM              ITEMS_FIELD = 15
	ITEMS_FIELD_MULTIPLE_OF       ITEMS_FIELD = 16
)

func (p ITEMS_FIELD) String() string {
	switch p {
	case ITEMS_FIELD_TYPE:
		return "type"
	case ITEMS_FIELD_FORMAT:
		return "format"
	case ITEMS_FIELD_ITEMS:
		return "items"
	case ITEMS_FIELD_COLLECTION_FORMAT:
		return "collectionFormat"
	case ITEMS_FIELD_DEFAULT:
		return "default"
	case ITEMS_FIELD_MAXIMUM:
		return "maximum"
	case ITEMS_FIELD_EXCLUSIVE_MAXIMUM:
		return "exclusiveMaximum"
	case ITEMS_FIELD_MINIMUM:
		return "minimum"
	case ITEMS_FIELD_EXCLUSIVE_MINIMUM:
		return "exclusiveMinimum"
	case ITEMS_FIELD_MAX_LENGTH:
		return "maxLength"
	case ITEMS_FIELD_MIN_LENGTH:
		return "minLength"
	case ITEMS_FIELD_PATTERN:
		return "pattern"
	case ITEMS_FIELD_MAX_ITEMS:
		return "maxItems"
	case ITEMS_FIELD_MIN_ITEMS:
		return "minItems"
	case ITEMS_FIELD_UNIQUE_ITEMS:
		return "uniqueItems"
	case ITEMS_FIELD_ENUM:
		return "enum"
	case ITEMS_FIELD_MULTIPLE_OF:
		return "multipleOf"
	}
	return "<UNSET>"
}

func ITEMS_FIELDFromString(s string) (ITEMS_FIELD, error) {
	switch s {
	case "type":
		return ITEMS_FIELD_TYPE, nil
	case "format":
		return ITEMS_FIELD_FORMAT, nil
	case "items":
		return ITEMS_FIELD_ITEMS, nil
	case "collectionFormat":
		return ITEMS_FIELD_COLLECTION_FORMAT, nil
	case "default":
		return ITEMS_FIELD_DEFAULT, nil
	case "maximum":
		return ITEMS_FIELD_MAXIMUM, nil
	case "exclusiveMaximum":
		return ITEMS_FIELD_EXCLUSIVE_MAXIMUM, nil
	case "minimum":
		return ITEMS_FIELD_MINIMUM, nil
	case "exclusiveMinimum":
		return ITEMS_FIELD_EXCLUSIVE_MINIMUM, nil
	case "maxLength":
		return ITEMS_FIELD_MAX_LENGTH, nil
	case "minLength":
		return ITEMS_FIELD_MIN_LENGTH, nil
	case "pattern":
		return ITEMS_FIELD_PATTERN, nil
	case "maxItems":
		return ITEMS_FIELD_MAX_ITEMS, nil
	case "minItems":
		return ITEMS_FIELD_MIN_ITEMS, nil
	case "uniqueItems":
		return ITEMS_FIELD_UNIQUE_ITEMS, nil
	case "enum":
		return ITEMS_FIELD_ENUM, nil
	case "multipleOf":
		return ITEMS_FIELD_MULTIPLE_OF, nil
	}
	return ITEMS_FIELD(0), fmt.Errorf("not a valid ITEMS_FIELD string")
}
