package shared

import (
	"html/template"
	"io"
	"log"
	"path"
)

type TemplateLoader struct {
	reloadTemplates bool
	templatePath string
}


func NewTemplateLoader(reloadTemplates bool, templatePath string) *TemplateLoader{
	return &TemplateLoader{templatePath:templatePath,reloadTemplates:reloadTemplates}
}


type LoadedTemplate struct {

	files    []string
	template *template.Template
	reload   bool
}



func resolveTemplates(templatePath string, files []string) []string {
	if len(templatePath) == 0 {
		return files
	}

	result := make([]string , len(files))

	for i, f := range files {
		result[i]= path.Join(templatePath, f)
	}
	return result
}

// DefineTemplate - parses a new template file, failing fast if error is found
func (tl *TemplateLoader) DefineTemplate(files ...string) *LoadedTemplate{
	t := template.New("layout")

	resolved := resolveTemplates(tl.templatePath, files)

	t, err := t.ParseFiles(resolved...)

	if err != nil {
		log.Fatal("Failed to load template", resolved, err)
	}

	return &LoadedTemplate{
		template: t,
		files:    resolved,
		reload:   tl.reloadTemplates,
	}
}

// ExecuteTemplate - applies the template to the data model and writes result to the output.
// if AlwaysReloadTemplates was set, then this will trigger template reload
func (loaded *LoadedTemplate) Exec(wr io.Writer, data interface{}) error {

	var t *template.Template

	if loaded.reload {
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