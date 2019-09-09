package graph

import (
	"encoding/hex"
	"fmt"
	"mlp/catalog/db"
	"mlp/catalog/web/shared"
	"strings"
)


func NewRender(tx *db.Tx, url shared.UrlResolver, fmt shared.Format) *SvgRender{
	return &SvgRender{
		tx:tx,

		url:url,
	}
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


func (s *SvgRender) String() string {
	return s.sb.String()
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
	s.Line("  \"%s\" [label=\"%s\" href=\"%s\"];// DatasetVer\n", hx(uid), ds.Title, link)
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
