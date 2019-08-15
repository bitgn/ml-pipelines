package core

import (
	"log"
	"mlp/catalog/db"
	"mlp/catalog/projection"
	"strings"
)

type UpgradePolicy int

const (
	UpgradePolicy_None UpgradePolicy = iota;
	UpgradePolicy_Auto;
	UpgradePolicy_Force;
)

func UpgradeDB(env *db.DB, version string, mode UpgradePolicy) {
	tx := env.MustWrite()
	defer tx.MustCleanup()





	ver := db.GetVersion(tx)
	if !upgradeNeeded(version, ver, mode){
		return
	}

	log.Printf("Upgrade is needed from '%s' to '%s' under policy %v\n", version, ver.ProjectionVersion, string(mode))
	deleted := db.WipeViewTables(tx)
	replayed := db.ReplayEvents(tx, projection.Handle)

	ver.ProjectionVersion = version
	db.PutVersion(tx, ver)

	log.Printf("Purged %d values and replayed %d events", deleted, replayed)

	tx.MustCommit()
}


func upgradeNeeded(version string, data *db.AppVersionData, mode UpgradePolicy) bool {
	switch mode {
	case UpgradePolicy_None:
		return false
	case UpgradePolicy_Force:
		return true
	case UpgradePolicy_Auto:
		return version == "dev" || !strings.EqualFold(version, data.ProjectionVersion)
	default:
		log.Panicln("Unexpected upgrade mode", mode)
		return false
	}

}
