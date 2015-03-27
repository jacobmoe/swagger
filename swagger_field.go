package swagger

import "fmt"

type SWAGGER_FIELD int64

const (
	SWAGGER_FIELD_VERSION              SWAGGER_FIELD = 0
	SWAGGER_FIELD_INFO                 SWAGGER_FIELD = 1
	SWAGGER_FIELD_HOST                 SWAGGER_FIELD = 2
	SWAGGER_FIELD_BASE_PATH            SWAGGER_FIELD = 3
	SWAGGER_FIELD_SCHEMES              SWAGGER_FIELD = 4
	SWAGGER_FIELD_CONSUMES             SWAGGER_FIELD = 5
	SWAGGER_FIELD_PRODUCES             SWAGGER_FIELD = 6
	SWAGGER_FIELD_PATHS                SWAGGER_FIELD = 7
	SWAGGER_FIELD_DEFINITIONS          SWAGGER_FIELD = 8
	SWAGGER_FIELD_PARAMETERS           SWAGGER_FIELD = 9
	SWAGGER_FIELD_RESPONSES            SWAGGER_FIELD = 10
	SWAGGER_FIELD_SECURITY_DEFINITIONS SWAGGER_FIELD = 11
	SWAGGER_FIELD_SECURITY             SWAGGER_FIELD = 12
	SWAGGER_FIELD_TAGS                 SWAGGER_FIELD = 13
	SWAGGER_FIELD_EXTERNAL_DOCS        SWAGGER_FIELD = 14
)

func (p SWAGGER_FIELD) Required() bool {
	switch p {
	case
		SWAGGER_FIELD_VERSION,
		SWAGGER_FIELD_INFO,
		SWAGGER_FIELD_PATHS:
		return true
	default:
		return false
	}
}

func (p SWAGGER_FIELD) String() string {
	switch p {
	case SWAGGER_FIELD_VERSION:
		return "swagger"
	case SWAGGER_FIELD_INFO:
		return "info"
	case SWAGGER_FIELD_HOST:
		return "host"
	case SWAGGER_FIELD_BASE_PATH:
		return "basePath"
	case SWAGGER_FIELD_SCHEMES:
		return "schemes"
	case SWAGGER_FIELD_CONSUMES:
		return "consumes"
	case SWAGGER_FIELD_PRODUCES:
		return "produces"
	case SWAGGER_FIELD_PATHS:
		return "paths"
	case SWAGGER_FIELD_DEFINITIONS:
		return "definitions"
	case SWAGGER_FIELD_PARAMETERS:
		return "parameters"
	case SWAGGER_FIELD_RESPONSES:
		return "responses"
	case SWAGGER_FIELD_SECURITY_DEFINITIONS:
		return "securityDefinitions"
	case SWAGGER_FIELD_SECURITY:
		return "security"
	case SWAGGER_FIELD_TAGS:
		return "tags"
	case SWAGGER_FIELD_EXTERNAL_DOCS:
		return "externalDocs"
	}
	return "<UNSET>"
}

func SWAGGER_FIELDFromString(s string) (SWAGGER_FIELD, error) {
	switch s {
	case "swagger":
		return SWAGGER_FIELD_VERSION, nil
	case "info":
		return SWAGGER_FIELD_INFO, nil
	case "host":
		return SWAGGER_FIELD_HOST, nil
	case "basePath":
		return SWAGGER_FIELD_BASE_PATH, nil
	case "schemes":
		return SWAGGER_FIELD_SCHEMES, nil
	case "consumes":
		return SWAGGER_FIELD_CONSUMES, nil
	case "produces":
		return SWAGGER_FIELD_PRODUCES, nil
	case "paths":
		return SWAGGER_FIELD_PATHS, nil
	case "definitions":
		return SWAGGER_FIELD_DEFINITIONS, nil
	case "parameters":
		return SWAGGER_FIELD_PARAMETERS, nil
	case "responses":
		return SWAGGER_FIELD_RESPONSES, nil
	case "securityDefinitions":
		return SWAGGER_FIELD_SECURITY_DEFINITIONS, nil
	case "security":
		return SWAGGER_FIELD_SECURITY, nil
	case "tags":
		return SWAGGER_FIELD_TAGS, nil
	case "externalDocs":
		return SWAGGER_FIELD_EXTERNAL_DOCS, nil
	}
	return SWAGGER_FIELD(0), fmt.Errorf("not a valid SWAGGER_FIELD string: " + s)
}
