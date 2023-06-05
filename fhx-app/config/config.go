package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the applications config
type AppConfig struct {
	UseCache       bool
	TemplateCatche map[string]*template.Template
	InfoLog        *log.Logger
	InProduction   bool
	Session        *scs.SessionManager
}
