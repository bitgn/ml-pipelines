package web

import (
	"html/template"
	"io"
	"log"
)

var reloadTemplates = false

// AlwaysReloadTemplates - tells web server that we want to reload web templates
// on every page refresh. This is useful for the development
func AlwaysReloadTemplates(){
	reloadTemplates = true
}


type LoadedTemplate struct {

	files    []string
	template *template.Template
}

// DefineTemplate - parses a new template file, failing fast if error is found
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

// ExecuteTemplate - applies the template to the data model and writes result to the output.
// if AlwaysReloadTemplates was set, then this will trigger template reload
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
