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

	has := form.Has("name", r)
	if has {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("name", "raymond")

	r = httptest.NewRequest("POST", "/anything", nil)
	form = New(postedData)

	has = form.Has("name", r)
	if !has {
		t.Error("Data has been populated, but does not validated")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	isMin := form.MinLength("name", 1, r)
	if isMin {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("name", "raymond")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	isMin = form.MinLength("name", 1, r)
	if !isMin {
		t.Error("Data has been populated, but does not validated")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/anything", nil)
	form := New(r.PostForm)

	isEmail := form.IsEmail("email")
	if isEmail {
		t.Error("Data has not been populated, but still it validated")
	}

	postedData := url.Values{}
	postedData.Add("email", "a@a.com")

	r = httptest.NewRequest("POST", "/anything", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	isEmail = form.IsEmail("email")
	if !isEmail {
		t.Error("Data has been populated, but does not validated")
	}
}