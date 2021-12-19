package handlers

import (
	"net/http"

	"github.com/jcaraballo113/go-bookings/pkg/config"
	"github.com/jcaraballo113/go-bookings/pkg/models"
	"github.com/jcaraballo113/go-bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// TemplateData holds data sent to templates


var Repo * Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewRepo
func NewHandlers(r *Repository) {
	Repo = r
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

