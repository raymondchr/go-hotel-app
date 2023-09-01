package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
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

// Valid checked if form input is empty or not
func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields{
		value := f.Get(field)
			if strings.TrimSpace(value) == ""{
				f.Errors.Add(field, "This cannot be blank")
			}
	}
}

func (f *Form) MinLength(fields string, length int, r *http.Request) bool {
	x := r.Form.Get(fields)
	if len(x) < length {
		f.Errors.Add(fields, fmt.Sprintf("Minimal length of input is %d", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string){
	if !govalidator.IsEmail(f.Get(field)){
		f.Errors.Add(field, "Please enter a valid email")
	}
}