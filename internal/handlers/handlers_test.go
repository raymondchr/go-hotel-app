package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name                 string
	url                  string
	method               string
	params               []postData
	expectedResponseCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarter", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-avail", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-02"},
		{key: "start", value: "2020-01-03"},
	}, http.StatusOK},
	{"search-avail-post-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-02"},
		{key: "start", value: "2020-01-03"},
	}, http.StatusOK},
	{"reservation summary", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Ray"},
		{key: "last_name", value: "Chris"},
		{key: "email", value: "a@a.com"},
		{key: "phone_number", value: "123123"},
	}, http.StatusOK},
	{"reserve summary","/reservation-summary","GET", []postData{
		
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				log.Fatal(err)
			}

			if resp.StatusCode != e.expectedResponseCode {
				t.Errorf("for %s, expected %d but got %d instead", e.name, e.expectedResponseCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x :=range e.params{
				values.Add(x.key,x.value)
			}

			resp,err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				log.Fatal(err)
			}

			if resp.StatusCode != e.expectedResponseCode {
				t.Errorf("for %s, expected %d but got %d instead", e.name, e.expectedResponseCode, resp.StatusCode)
			}
		}
	}
}
