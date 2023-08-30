package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/raymondchr/go-hotel-app/cmd/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf(fmt.Sprintf("Type is %T instead of *chi.Mux", v))
	}

}
