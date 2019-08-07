package web

import "mlp/catalog/db"

type ViewDatsetModel struct {
	*Site
	Dataset *db.DatasetData
	Stale bool
}

