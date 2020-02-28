package forms

import (
	"fmt"
	"net/url"
)

type Form struct {
	Values url.Values
	Errors map[string][]string
}

func NewForm(values url.Values) *Form {
	return &Form{
		Values: values,
		Errors: map[string][]string{},
	}
}

func (f *Form) ValidateRequired(field string) {
	value := f.Values.Get(field)
	if value == "" {
		f.Errors[field] = append(f.Errors[field], "This field is required.")
	}
}

func (f *Form) ValidateURL(field string) {
	value := f.Values.Get(field)
	if value == "" {
		return
	}

	_, err := url.ParseRequestURI(value)
	if err != nil {
		f.Errors[field] = append(f.Errors[field], "A valid URL is required.")
	}
}

func (f *Form) ValidateMinLength(field string, min int) {
	value := f.Values.Get(field)
	if value == "" {
		return
	}

	if len(value) < min {
		f.Errors[field] = append(f.Errors[field], fmt.Sprintf("A minimum length of %d is required.", min))
	}
}

func (f *Form) ValidateMatch(field1, field2 string) {
	value1 := f.Values.Get(field1)
	if value1 == "" {
		return
	}

	value2 := f.Values.Get(field2)
	if value2 == "" {
		return
	}

	if value1 != value2 {
		f.Errors[field1] = append(f.Errors[field1], fmt.Sprintf("Value must match %s.", field2))
	}
}

func (f *Form) ContainsErrors() bool {
	return len(f.Errors) != 0
}
