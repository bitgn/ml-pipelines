package shared

import "mlp/catalog/db"

type Site struct{
	Stats *db.TenantStats
	AppVersion string
	ActiveMenu string
	Url UrlResolver
	Fmt Format
}



var version = "dev"



func SetVersion(ver string){
	version = ver
}

func LoadSite(tx *db.Tx) *Site{
	stats := db.GetStats(tx)
	s := &Site{
		Stats:stats,
		AppVersion:version,
	}
	return s
}




