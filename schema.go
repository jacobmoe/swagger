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

func (schema *Schema) TypeOf(name string) *Schema {
	schema.Set(SCHEMA_FIELD_TYPE, name)
	return schema
}

func (schema *Schema) ArrayOf(name string) *Schema {
	schema.Set(SCHEMA_FIELD_TYPE, "array")
	items := NewSchema()
	items.Set(REFERENCE_FIELD_REF, "#/definitions/"+name)
	schema.Set(SCHEMA_FIELD_ITEMS, items)
	return schema
}

func (schema *Schema) RefOf(name string) *Schema {
	schema.Set("$ref", "#/definitions/"+name)
	return schema
}

func (s *Schema) Generate(indent bool, valid bool) []byte {
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
		s.AddError(err)
		return nil
	}
	if valid {
		doc, err := spec.New(j, "")
		if err != nil {
			s.AddError(err)
			return nil
		}
		result := validate.Spec(doc, strfmt.Default)
		if result != nil {
			for _, desc := range result.(*swaggererrors.CompositeError).Errors {
				s.AddError(fmt.Errorf("The swagger spec is invalid against swagger specification %s. %s", doc.Version(), desc.Error()))
			}
			return nil
		}
	}
	return j
}
