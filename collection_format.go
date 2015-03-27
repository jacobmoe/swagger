package swagger

import "fmt"

type COLLECTION_FORMAT int64

const (
	COLLECTION_FORMAT_CSV   COLLECTION_FORMAT = 0
	COLLECTION_FORMAT_SSV   COLLECTION_FORMAT = 1
	COLLECTION_FORMAT_TSV   COLLECTION_FORMAT = 2
	COLLECTION_FORMAT_PIPES COLLECTION_FORMAT = 3
	COLLECTION_FORMAT_MULTI COLLECTION_FORMAT = 4
)

func (p COLLECTION_FORMAT) String() string {
	switch p {
	case COLLECTION_FORMAT_CSV:
		return "csv"
	case COLLECTION_FORMAT_SSV:
		return "ssv"
	case COLLECTION_FORMAT_TSV:
		return "tsv"
	case COLLECTION_FORMAT_PIPES:
		return "pipes"
	case COLLECTION_FORMAT_MULTI:
		return "multi"
	}
	return "<UNSET>"
}

func COLLECTION_FORMATFromString(s string) (COLLECTION_FORMAT, error) {
	switch s {
	case "csv":
		return COLLECTION_FORMAT_CSV, nil
	case "ssv":
		return COLLECTION_FORMAT_SSV, nil
	case "tsv":
		return COLLECTION_FORMAT_TSV, nil
	case "pipes":
		return COLLECTION_FORMAT_PIPES, nil
	case "multi":
		return COLLECTION_FORMAT_MULTI, nil
	}
	return COLLECTION_FORMAT(0), fmt.Errorf("not a valid COLLECTION_FORMAT string: " + s)
}
