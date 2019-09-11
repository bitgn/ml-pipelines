package domain

import (
	"mlp/catalog/db"
	"mlp/catalog/sim"
	"time"
)

/*
   if not dm.update_timestamp_set:
       return False

   was = dt.datetime.utcfromtimestamp(dm.update_timestamp)

   now = config.utcnow()

   stale = (now - was).days >= 3
   return stale
 */

func IsStale(ds *db.DatasetData) bool {
	if ds.UpdateTimestamp == 0 {
		return false
	}

	delta := sim.UTC().Sub(time.Unix(ds.UpdateTimestamp,0))

	days := delta.Hours() / 24

	return days >= 3

}



func IsJobStale(run *db.JobRunData) bool {
	if run == nil {
		return false
	}

	delta := sim.UTC().Sub(time.Unix(run.UpdateTimestamp,0))

	days := delta.Hours() / 24

	return days >= 3

}

