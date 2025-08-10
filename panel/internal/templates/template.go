package templates

import (
	"embed"
	"html/template"
	"io"
)

var templates embed.FS

type TemplateService struct {
	templates *template.Template
}

func NewTemplateService() (*TemplateService, error) {
	tmpl, err := template.ParseFS(templates, "*.html")
	if err != nil {
		return nil, err
	}

	return &TemplateService{
		templates: tmpl,
	}, nil
}

func (ts *TemplateService) RenderTemplate(w io.Writer, templateName string, data interface{}) error {
	return ts.templates.ExecuteTemplate(w, templateName, data)
}

type TemplateData struct {
	Title   string
	Message string
	Error   string
}
