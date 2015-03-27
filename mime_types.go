package swagger

import (
	"encoding/json"
	"fmt"
)

type MIME_TYPE int64

const (
	MIME_TYPE_TEXT_PLAIN_CHARSET_UTF_8            MIME_TYPE = 0
	MIME_TYPE_APPLICATION_JSON                    MIME_TYPE = 1
	MIME_TYPE_APPLICATION_VND_GITHUB_JSON         MIME_TYPE = 2
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_JSON      MIME_TYPE = 3
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_RAW_JSON  MIME_TYPE = 4
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_TEXT_JSON MIME_TYPE = 5
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_HTML_JSON MIME_TYPE = 6
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_FULL_JSON MIME_TYPE = 7
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_DIFF      MIME_TYPE = 8
	MIME_TYPE_APPLICATION_VND_GITHUB_V3_PATCH     MIME_TYPE = 9
	MIME_TYPE_APPLICATION_XML                     MIME_TYPE = 10
	MIME_TYPE_APPLICATION_X_WWW_FORM_URLENCODED   MIME_TYPE = 11
	MIME_TYPE_MULTIPART_FORM_DATA                 MIME_TYPE = 12
	// Add more mime types here
)

func (p MIME_TYPE) String() string {
	switch p {
	case MIME_TYPE_TEXT_PLAIN_CHARSET_UTF_8:
		return "text/plain; charset=utf-8"
	case MIME_TYPE_APPLICATION_JSON:
		return "application/json"
	case MIME_TYPE_APPLICATION_XML:
		return "application/xml"
	case MIME_TYPE_APPLICATION_VND_GITHUB_JSON:
		return "application/vnd.github+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_JSON:
		return "application/vnd.github.v3+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_RAW_JSON:
		return "application/vnd.github.v3.raw+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_TEXT_JSON:
		return "application/vnd.github.v3.text+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_HTML_JSON:
		return "application/vnd.github.v3.html+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_FULL_JSON:
		return "application/vnd.github.v3.full+json"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_DIFF:
		return "application/vnd.github.v3.diff"
	case MIME_TYPE_APPLICATION_VND_GITHUB_V3_PATCH:
		return "application/vnd.github.v3.patch"
	case MIME_TYPE_APPLICATION_X_WWW_FORM_URLENCODED:
		return "application/x-www-form-urlencoded"
	case MIME_TYPE_MULTIPART_FORM_DATA:
		return "multipart/form-data"
	}
	return "<UNSET>"
}

func MIME_TYPEFromString(s string) (MIME_TYPE, error) {
	switch s {
	case "text/plain; charset=utf-8":
		return MIME_TYPE_TEXT_PLAIN_CHARSET_UTF_8, nil
	case "application/json":
		return MIME_TYPE_APPLICATION_JSON, nil
	case "application/xml":
		return MIME_TYPE_APPLICATION_XML, nil
	case "application/vnd.github+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_JSON, nil
	case "application/vnd.github.v3+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_JSON, nil
	case "application/vnd.github.v3.raw+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_RAW_JSON, nil
	case "application/vnd.github.v3.text+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_TEXT_JSON, nil
	case "application/vnd.github.v3.html+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_HTML_JSON, nil
	case "application/vnd.github.v3.full+json":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_FULL_JSON, nil
	case "application/vnd.github.v3.diff":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_DIFF, nil
	case "application/vnd.github.v3.patch":
		return MIME_TYPE_APPLICATION_VND_GITHUB_V3_PATCH, nil
	case "application/x-www-form-urlencoded":
		return MIME_TYPE_APPLICATION_X_WWW_FORM_URLENCODED, nil
	case "multipart/form-data":
		return MIME_TYPE_MULTIPART_FORM_DATA, nil
	}
	return MIME_TYPE(0), fmt.Errorf("not a valid MIME_TYPE string: " + s)
}

func (p MIME_TYPE) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}
