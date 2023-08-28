package forms

import (
	"net/http"
	"net/url"
)

// Form creates a custom form struct
type Form struct {
	url.Values
	Errors errors
}

// New create a new form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checked if form field is not empty
func (f *Form) Has(field string, r *http.Request) bool{
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field,"This field cannot be empty")
		return false
	}

	return true
}

func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}