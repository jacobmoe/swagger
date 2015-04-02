package swagger

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Common struct {
	errors []error
	data   map[string]reflect.Value
}

type CommonInterface interface {
	Errors() []error
}

func NewCommon() *Common {
	return &Common{data: make(map[string]reflect.Value, 0), errors: make([]error, 0)}
}

func (s *Common) HasErrors() bool {
	return (len(s.errors) > 0 || len(s.Errors()) > 0)
}

func (s *Common) Error() string {
	var errors []string
	for _, err := range s.Errors() {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, "\n")
}

func (s *Common) Errors() []error {
	var errors []error
	for _, err := range s.errors {
		errors = append(errors, err)
	}
	for _, v := range s.data {
		if o, ok := v.Interface().(CommonInterface); ok {
			errs := o.Errors()
			for _, err := range errs {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func (p *Common) Set(name interface{}, object interface{}) {
	p.data[fmt.Sprintf("%v", name)] = reflect.ValueOf(object)
}

func (s *Common) ExtendSlice(i interface{}, with interface{}) {
	name := fmt.Sprintf("%v", i)
	m, ok := s.data[name]
	if !ok {
		m = reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(with)), 0, 0)
	}
	s.data[name] = reflect.Append(m, reflect.ValueOf(with))
}

func (s *Common) ExtendMap(i interface{}, with interface{}) {
	name := fmt.Sprintf("%v", i)
	m, ok := s.data[name]
	if !ok {
		m = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(with)))
	}
	m.SetMapIndex(reflect.ValueOf(name), reflect.ValueOf(with))
	s.data[name] = m
}

func (s *Common) addError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	s.errors = append(s.errors, fmt.Errorf("[error] %s:%d %v", fn, line, err))
}

func (s *Common) MarshalJSON() ([]byte, error) {
	data := make(map[string]interface{}, 0)
	for k, v := range s.data {
		data[k] = v.Interface()
	}
	return json.Marshal(data)
}
