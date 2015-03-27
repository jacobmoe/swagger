package swagger

import "fmt"

type SCHEMA_FIELD int64

const (
	SCHEMA_FIELD_REF               SCHEMA_FIELD = 0
	SCHEMA_FIELD_FORMAT            SCHEMA_FIELD = 1
	SCHEMA_FIELD_TITLE             SCHEMA_FIELD = 2
	SCHEMA_FIELD_DESCRIPTION       SCHEMA_FIELD = 3
	SCHEMA_FIELD_DEFAULT           SCHEMA_FIELD = 4
	SCHEMA_FIELD_MULTIPLE_OF       SCHEMA_FIELD = 5
	SCHEMA_FIELD_MAXIMUM           SCHEMA_FIELD = 6
	SCHEMA_FIELD_EXCLUSIVE_MAXIMUM SCHEMA_FIELD = 7
	SCHEMA_FIELD_MINIMUM           SCHEMA_FIELD = 8
	SCHEMA_FIELD_EXCLUSIVE_MINIMUM SCHEMA_FIELD = 9
	SCHEMA_FIELD_MAX_LENGTH        SCHEMA_FIELD = 10
	SCHEMA_FIELD_MIN_LENGTH        SCHEMA_FIELD = 11
	SCHEMA_FIELD_PATTERN           SCHEMA_FIELD = 12
	SCHEMA_FIELD_MAX_ITEMS         SCHEMA_FIELD = 13
	SCHEMA_FIELD_MIN_ITEMS         SCHEMA_FIELD = 14
	SCHEMA_FIELD_UNIQUE_ITEMS      SCHEMA_FIELD = 15
	SCHEMA_FIELD_MAX_PROPERTIES    SCHEMA_FIELD = 16
	SCHEMA_FIELD_MIN_PROPERTIES    SCHEMA_FIELD = 17
	SCHEMA_FIELD_REQUIRED          SCHEMA_FIELD = 18
	SCHEMA_FIELD_ENUM              SCHEMA_FIELD = 19
	SCHEMA_FIELD_TYPE              SCHEMA_FIELD = 20
	SCHEMA_FIELD_ITEMS             SCHEMA_FIELD = 21
	SCHEMA_FIELD_ALLOF             SCHEMA_FIELD = 22
	SCHEMA_FIELD_PROPERTIES        SCHEMA_FIELD = 23
	SCHEMA_FIELD_DISCRIMINATOR     SCHEMA_FIELD = 24
	SCHEMA_FIELD_READ_ONLY         SCHEMA_FIELD = 25
	SCHEMA_FIELD_XML               SCHEMA_FIELD = 26
	SCHEMA_FIELD_EXTERNAL_DOCS     SCHEMA_FIELD = 27
	SCHEMA_FIELD_EXAMPLE           SCHEMA_FIELD = 28
	SCHEMA_FIELD_OBJECT            SCHEMA_FIELD = 29
)

func (p SCHEMA_FIELD) String() string {
	switch p {
	case SCHEMA_FIELD_REF:
		return "$ref"
	case SCHEMA_FIELD_FORMAT:
		return "format"
	case SCHEMA_FIELD_TITLE:
		return "title"
	case SCHEMA_FIELD_DESCRIPTION:
		return "description"
	case SCHEMA_FIELD_DEFAULT:
		return "default"
	case SCHEMA_FIELD_MULTIPLE_OF:
		return "multipleOf"
	case SCHEMA_FIELD_MAXIMUM:
		return "maximum"
	case SCHEMA_FIELD_EXCLUSIVE_MAXIMUM:
		return "exclusiveMaximum"
	case SCHEMA_FIELD_MINIMUM:
		return "minimum"
	case SCHEMA_FIELD_EXCLUSIVE_MINIMUM:
		return "exclusiveMinimum"
	case SCHEMA_FIELD_MAX_LENGTH:
		return "maxLength"
	case SCHEMA_FIELD_MIN_LENGTH:
		return "minLength"
	case SCHEMA_FIELD_PATTERN:
		return "pattern"
	case SCHEMA_FIELD_MAX_ITEMS:
		return "maxItems"
	case SCHEMA_FIELD_MIN_ITEMS:
		return "minItems"
	case SCHEMA_FIELD_UNIQUE_ITEMS:
		return "uniqueItems"
	case SCHEMA_FIELD_MAX_PROPERTIES:
		return "maxProperties"
	case SCHEMA_FIELD_MIN_PROPERTIES:
		return "minProperties"
	case SCHEMA_FIELD_REQUIRED:
		return "required"
	case SCHEMA_FIELD_ENUM:
		return "enum"
	case SCHEMA_FIELD_TYPE:
		return "type"
	case SCHEMA_FIELD_ITEMS:
		return "items"
	case SCHEMA_FIELD_ALLOF:
		return "allOf"
	case SCHEMA_FIELD_PROPERTIES:
		return "properties"
	case SCHEMA_FIELD_DISCRIMINATOR:
		return "discriminator"
	case SCHEMA_FIELD_READ_ONLY:
		return "readOnly"
	case SCHEMA_FIELD_XML:
		return "xml"
	case SCHEMA_FIELD_EXTERNAL_DOCS:
		return "externalDocs"
	case SCHEMA_FIELD_EXAMPLE:
		return "example"
	case SCHEMA_FIELD_OBJECT:
		return "object"
	}
	return "<UNSET>"
}

func SCHEMA_FIELDFromString(s string) (SCHEMA_FIELD, error) {
	switch s {
	case "$ref":
		return SCHEMA_FIELD_REF, nil
	case "format":
		return SCHEMA_FIELD_FORMAT, nil
	case "title":
		return SCHEMA_FIELD_TITLE, nil
	case "description":
		return SCHEMA_FIELD_DESCRIPTION, nil
	case "default":
		return SCHEMA_FIELD_DEFAULT, nil
	case "multipleOf":
		return SCHEMA_FIELD_MULTIPLE_OF, nil
	case "maximum":
		return SCHEMA_FIELD_MAXIMUM, nil
	case "exclusiveMaximum":
		return SCHEMA_FIELD_EXCLUSIVE_MAXIMUM, nil
	case "minimum":
		return SCHEMA_FIELD_MINIMUM, nil
	case "exclusiveMinimum":
		return SCHEMA_FIELD_EXCLUSIVE_MINIMUM, nil
	case "maxLength":
		return SCHEMA_FIELD_MAX_LENGTH, nil
	case "minLength":
		return SCHEMA_FIELD_MIN_LENGTH, nil
	case "pattern":
		return SCHEMA_FIELD_PATTERN, nil
	case "maxItems":
		return SCHEMA_FIELD_MAX_ITEMS, nil
	case "minItems":
		return SCHEMA_FIELD_MIN_ITEMS, nil
	case "uniqueItems":
		return SCHEMA_FIELD_UNIQUE_ITEMS, nil
	case "maxProperties":
		return SCHEMA_FIELD_MAX_PROPERTIES, nil
	case "minProperties":
		return SCHEMA_FIELD_MIN_PROPERTIES, nil
	case "required":
		return SCHEMA_FIELD_REQUIRED, nil
	case "enum":
		return SCHEMA_FIELD_ENUM, nil
	case "type":
		return SCHEMA_FIELD_TYPE, nil
	case "items":
		return SCHEMA_FIELD_ITEMS, nil
	case "allOf":
		return SCHEMA_FIELD_ALLOF, nil
	case "properties":
		return SCHEMA_FIELD_PROPERTIES, nil
	case "discriminator":
		return SCHEMA_FIELD_DISCRIMINATOR, nil
	case "readOnly":
		return SCHEMA_FIELD_READ_ONLY, nil
	case "xml":
		return SCHEMA_FIELD_XML, nil
	case "externalDocs":
		return SCHEMA_FIELD_EXTERNAL_DOCS, nil
	case "example":
		return SCHEMA_FIELD_EXAMPLE, nil
	case "object":
		return SCHEMA_FIELD_OBJECT, nil
	}
	return SCHEMA_FIELD(0), fmt.Errorf("not a valid SCHEMA_FIELD string")
}
