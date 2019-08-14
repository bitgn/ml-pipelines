package web

import (
	"html/template"
	"io"
	"log"
)

var reloadTemplates = false


func EnableTemplateReloading(){
	reloadTemplates = true
}


type LoadedTemplate struct {

	files    []string
	template *template.Template
}


func DefineTemplate(files ...string) *LoadedTemplate{


	t := template.New("layout")

	t, err := t.ParseFiles(files...)

	if err != nil {
		log.Fatal("Failed to load template", files, err)
	}

	return &LoadedTemplate{
		template: t,
		files:    files,
	}
}

func (loaded *LoadedTemplate) ExecuteTemplate(wr io.Writer, data interface{}) error {

	var t *template.Template

	if reloadTemplates{
		var err error
		parsed := template.New("layout")
		t, err = parsed.ParseFiles(loaded.files...)
		if err != nil {
			return err
		}
	} else {
		t = loaded.template
	}

	return t.ExecuteTemplate(wr, "layout", data)
}
