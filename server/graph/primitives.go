package graph

import (
	"encoding/hex"
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/web/shared"
	"os/exec"
	"strings"
)


func NewRender(tx *db.Tx, url shared.UrlResolver, fmt shared.Format) *SvgRender{
	render := &SvgRender{
		tx: tx,

		url: url,
	}

	render.Line("digraph{")
	render.Line("rankdir=LR;")
	render.Line("fontname=\"Arial\";")
	render.Line("node[shape=\"rectangle\" color=\"#343a40\" penwidth=\"1.5\" fontname=\"Arial\"];")
	render.Line("edge[color=\"#343a40\" penwidth=\"1.0\"];")



	return render
}



type SvgRender struct {
	tx *db.Tx
	sb strings.Builder
	url shared.UrlResolver
	fmt shared.Format
}


func (s *SvgRender) Line(l string, a ...interface{}){

	if len(a)==0{
		s.sb.WriteString(l + "\n")
		return
	}
	s.sb.WriteString(fmt.Sprintf(l, a...) + "\n")
}

func (s *SvgRender) JobRun(uid []byte) {

	run := db.GetJobRun(s.tx, uid)
	job := db.GetJob(s.tx, run.JobUid)
	title := s.fmt.Timestamp(run.Timestamp)
	runTitle := fmt.Sprintf("<%s<BR/><I>%s</I>>", job.Caption(), title)
	s.sb.WriteString(fmt.Sprintf("  \"%s\" [label=%s style=\"rounded\"]; // JobRun \n", hx(run.Uid), runTitle))
}



func (s *SvgRender) Dash(from, to []byte){
	s.sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowhead=\"none\"];", hx(from), hx(to)))
}

func (s *SvgRender) Arrow(from, to []byte){
	s.sb.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\" [arrowtail=\"none\"];", hx(from), hx(to)))
}




func(s *SvgRender) DatasetVer(uid []byte){
	ver := db.GetDatasetVersion(s.tx,uid)
	ds := db.GetDataset(s.tx, ver.DatasetUid)
	link := s.url.ViewDatasetVersion(ds.ProjectName, ds.Name, ver.Uid)

	title := fmt.Sprintf("<%s<BR/><I>%s</I>>", ds.Title, s.fmt.Timestamp(ver.Timestamp))
	s.Line("  \"%s\" [label=%s href=\"%s\"];// DatasetVer\n", hx(uid), title, link)
}


func (s *SvgRender) ServiceVer(uid []byte){
	ver := db.GetServiceVersion(s.tx, uid)
	svc := db.GetService(s.tx, ver.ServiceUid)

	link := s.url.ViewServiceVer(svc.ProjectName, svc.Name, ver.VersionNum)

	title := fmt.Sprintf("<%s<BR/><I>%s</I>>", svc.Caption(), s.fmt.Timestamp(ver.Timestamp))
	s.sb.WriteString(fmt.Sprintf("  \"%s\" [label=%s href=\"%s\"]; // ServiceVer\n", hx(ver.Uid), title, link))
}


func (s *SvgRender) ColorGreen(uid []byte){
	s.Line("\"%s\"[color=\"#28a745\"]", hx(uid))
}


func (s *SvgRender) Service(uid []byte){
	svc := db.GetService(s.tx, uid)

	link := s.url.ViewService(svc.ProjectName, svc.Name)
	this := hx(svc.Uid)
	s.sb.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\" href=\"%s\"];\n", this, svc.Caption(), link))
}


func hx(b []byte) string{
	return hex.EncodeToString(b)
}


func renderDot(line string) ([]byte, error){
	cmd := exec.Command("dot", "-Tsvg")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		_, err := io.WriteString(stdin, line)
		if err != nil {
			log.Fatal(errors.Wrap(err, "Failed to write to stdin"))
		}

	}()

	return cmd.CombinedOutput()
}


func (s *SvgRender) ToHtml() template.HTML {

	s.Line("}")

	source := s.sb.String()
	result, err := renderDot(source)



	if err != nil {
		log.Println("Problem with SVG:", string(result))
		log.Println("Original input:")
		log.Println(source)
		return template.HTML(err.Error())
	}

	str := string(result)

	str = strings.Replace(str, "<svg ", `<svg class="img-fluid" `, 1)
	return template.HTML(str)
}