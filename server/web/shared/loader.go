package shared

import (
	"bufio"
	"bytes"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"log"
	"net/http"
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


func (loaded *LoadedTemplate) Render(w http.ResponseWriter, model interface{}){


	var b bytes.Buffer
	foo := bufio.NewWriter(&b)


	if err := loaded.exec(foo, model); err != nil {
		http.Error(w, err.Error(), 408)
	} else {
		foo.Flush()
		b.WriteTo(w)

	}


}
// ExecuteTemplate - applies the template to the data model and writes result to the output.
// if AlwaysReloadTemplates was set, then this will trigger template reload
func (loaded *LoadedTemplate) exec(wr io.Writer, data interface{}) error {

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

	err2 := t.ExecuteTemplate(wr, "layout", data)
	if err2 != nil {
		return errors.WithMessagef(err2, "%v", loaded.files)
	}
	return nil
}