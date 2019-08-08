package web

type Site struct{
	ProjectCount int
	DatasetCount int
	JobCount int
	ModelCount int
	ReportCount int
	ExpertCount int
	AppVersion string
	ActiveMenu string

	Url UrlResolver
	Fmt Format
}


