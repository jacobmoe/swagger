package swagger

import (
	"encoding/json"
	"fmt"

	swaggererrors "github.com/casualjim/go-swagger/errors"
	"github.com/casualjim/go-swagger/spec"
	"github.com/casualjim/go-swagger/strfmt"
	"github.com/casualjim/go-swagger/validate"
)

type Schema struct {
	*Common
}

func NewSchema() *Schema {
	return &Schema{NewCommon()}
}

func (s *Schema) Generate(indent bool) []byte {
	var (
		j   []byte
		err error
	)
	if indent {
		j, err = json.MarshalIndent(s, "", "    ")
	} else {
		j, err = json.Marshal(s)
	}
	if err != nil {
		s.addError(err)
		return nil
	}
	doc, err := spec.New(j, "")
	if err != nil {
		s.addError(err)
		return nil
	}
	result := validate.Spec(doc, strfmt.Default)
	if result != nil {
		for _, desc := range result.(*swaggererrors.CompositeError).Errors {
			s.addError(fmt.Errorf("The swagger spec is invalid against swagger specification %s. %s", doc.Version(), desc.Error()))
		}
		return nil
	}
	return j
}
