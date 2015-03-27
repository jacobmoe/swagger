package swagger

import "fmt"

type REFERENCE_FIELD int64

const (
	REFERENCE_FIELD_REF REFERENCE_FIELD = 0
)

func (p REFERENCE_FIELD) String() string {
	switch p {
	case REFERENCE_FIELD_REF:
		return "$ref"
	}
	return "<UNSET>"
}

func REFERENCE_FIELDFromString(s string) (REFERENCE_FIELD, error) {
	switch s {
	case "$ref":
		return REFERENCE_FIELD_REF, nil
	}
	return REFERENCE_FIELD(0), fmt.Errorf("not a valid REFERENCE_FIELD string")
}
