package shared

import "mlp/catalog/db"

type Site struct{
	AppVersion string
	Url UrlResolver
	Fmt Format
}



var version = "dev"



func SetVersion(ver string){
	version = ver
}

func LoadSite(tx *db.Tx) *Site{
	s := &Site{
		AppVersion:version,
	}
	return s
}




