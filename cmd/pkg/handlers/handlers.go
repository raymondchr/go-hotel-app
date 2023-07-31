package handlers

import (
	"net/http"

	"github.com/raymondchr/go-hotel-app/cmd/pkg/config"
	"github.com/raymondchr/go-hotel-app/cmd/pkg/models"
	"github.com/raymondchr/go-hotel-app/cmd/pkg/render"
)

var Repo *Repository



type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, this a string taken from backend"
		
	
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
			StringMap: stringMap,
	})
}
