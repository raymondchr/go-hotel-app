package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	isValid := form.Valid()

	if !isValid {
		t.Error("Should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Data has been populated, but does not validated")
	}

}

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	has := form.Has("name")
	if has {
		t.Error("Data has not been populated, but still it validated")
	}

	isErr := form.Errors.Get("name")
	if isErr == ""{
		t.Error("It should return error, but does not")
	}

	postedData := url.Values{}
	postedData.Add("name", "raymond")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	has = form.Has("name")
	if !has {
		t.Error("Data has been populated, but does not validated")
	}

	isErr = form.Errors.Get("name")
	if isErr != ""{
		t.Error("It should not return error, but does")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	isMin := form.MinLength("name", 1)
	if isMin {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("name", "raymond")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	isMin = form.MinLength("name", 1)
	if !isMin {
		t.Error("Data has been populated, but does not validated")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("email", "a@a.com")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("email format is correct, but does not validated")
	}

	postedData = url.Values{}
	postedData.Add("email", "raymond")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Key is not an email format, but validated")
	}
}