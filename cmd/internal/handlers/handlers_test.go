package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
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
			//POST request Test here
		}
	}
}

// func TestRepository_Home(t *testing.T) {
// 	routes := getRoutes()
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		m    *Repository
// 		args args
// 	}{
// 		{
// 			//success
// 			name: "home",
// 			m: rep
// 			// {
// 			// 	UseCache:      true,
// 			// 	TemplateCache: CreateTestTemplateCache(),
// 			// 	InProduction:  false,
// 			// 	Session:       scs.New(),
// 			// },
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.m.Home(tt.args.w, tt.args.r)
// 		})
// 	}
// }
