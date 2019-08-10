package web

import "mlp/catalog/db"

type Site struct{
	Stats *db.TenantStats
	AppVersion string
	ActiveMenu string

	Url UrlResolver
	Fmt Format
}


func LoadSite(tx *db.Tx) *Site{
	stats := db.GetStats(tx)
	s := &Site{
		Stats:stats,
	}
	return s

}


