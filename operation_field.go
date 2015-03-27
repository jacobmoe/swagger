package swagger

import "fmt"

type OPERATION_FIELD int64

const (
	OPERATION_FIELD_TAGS          OPERATION_FIELD = 0
	OPERATION_FIELD_SUMMARY       OPERATION_FIELD = 1
	OPERATION_FIELD_DESCRIPTION   OPERATION_FIELD = 2
	OPERATION_FIELD_EXTERNAL_DOCS OPERATION_FIELD = 3
	OPERATION_FIELD_OPERATION_ID  OPERATION_FIELD = 4
	OPERATION_FIELD_CONSUMES      OPERATION_FIELD = 5
	OPERATION_FIELD_PRODUCES      OPERATION_FIELD = 6
	OPERATION_FIELD_PARAMETERS    OPERATION_FIELD = 7
	OPERATION_FIELD_RESPONSES     OPERATION_FIELD = 8
	OPERATION_FIELD_SCHEMES       OPERATION_FIELD = 9
	OPERATION_FIELD_DEPRECATED    OPERATION_FIELD = 10
	OPERATION_FIELD_SECURITY      OPERATION_FIELD = 11
)

func (p OPERATION_FIELD) String() string {
	switch p {
	case OPERATION_FIELD_TAGS:
		return "tags"
	case OPERATION_FIELD_SUMMARY:
		return "summary"
	case OPERATION_FIELD_DESCRIPTION:
		return "description"
	case OPERATION_FIELD_EXTERNAL_DOCS:
		return "externalDocs"
	case OPERATION_FIELD_OPERATION_ID:
		return "operationId"
	case OPERATION_FIELD_CONSUMES:
		return "consumes"
	case OPERATION_FIELD_PRODUCES:
		return "produces"
	case OPERATION_FIELD_PARAMETERS:
		return "parameters"
	case OPERATION_FIELD_RESPONSES:
		return "responses"
	case OPERATION_FIELD_SCHEMES:
		return "schemes"
	case OPERATION_FIELD_DEPRECATED:
		return "deprecated"
	case OPERATION_FIELD_SECURITY:
		return "security"
	}
	return "<UNSET>"
}

func OPERATION_FIELDFromString(s string) (OPERATION_FIELD, error) {
	switch s {
	case "tags":
		return OPERATION_FIELD_TAGS, nil
	case "summary":
		return OPERATION_FIELD_SUMMARY, nil
	case "description":
		return OPERATION_FIELD_DESCRIPTION, nil
	case "externalDocs":
		return OPERATION_FIELD_EXTERNAL_DOCS, nil
	case "operationId":
		return OPERATION_FIELD_OPERATION_ID, nil
	case "consumes":
		return OPERATION_FIELD_CONSUMES, nil
	case "produces":
		return OPERATION_FIELD_PRODUCES, nil
	case "parameters":
		return OPERATION_FIELD_PARAMETERS, nil
	case "responses":
		return OPERATION_FIELD_RESPONSES, nil
	case "schemes":
		return OPERATION_FIELD_SCHEMES, nil
	case "deprecated":
		return OPERATION_FIELD_DEPRECATED, nil
	case "security":
		return OPERATION_FIELD_SECURITY, nil
	}
	return OPERATION_FIELD(0), fmt.Errorf("not a valid OPERATION_FIELD string: " + s)
}
