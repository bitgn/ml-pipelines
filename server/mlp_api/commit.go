package mlp_api

import (
	"bytes"
	"golang.org/x/net/context"
	"log"
	"mlp/catalog/db"
	"mlp/catalog/events"
	"mlp/catalog/sim"
)

func (s *server) Commit(c context.Context, r *CommitRequest) (*CommitResponse, error) {
	tx := s.db.MustWrite()
	defer tx.MustCleanup()


	resp := &CommitResponse{}



	for _, changeSet := range r.Datasets{
		ds := db.GetDataset(tx, changeSet.DatasetUid)

		if !bytes.Equal(ds.HeadVersion, changeSet.ParentVersionUid) {
			return resp.err(versionMismatch(changeSet.ParentVersionUid, ds.HeadVersion))
		}

		if changeSet.Timestamp == 0{
			changeSet.Timestamp = sim.Unix()
		}


		e := &events.DatasetVersionAdded{
			DatasetUid:ds.Uid,
			ProjectUid:ds.ProjectUid,
			Title:r.Title,
			Uid:sim.NewID(),
			Timestamp:changeSet.Timestamp,
			ProjectName:ds.ProjectName,
			ParentUid:ds.HeadVersion,
			Inputs:changeSet.Inputs,
			Items:changeSet.Items,
			CleanSlate:changeSet.CleanSlate,
			VersionNum:ds.VersionCount+1,
		}

		if len(changeSet.Remove)>0{
			log.Panicln("Implement changeset removal")
		}

		s.publish(tx, e)

		resp.DatasetVersions = append(resp.DatasetVersions, e.Uid)
	}


	tx.MustCommit()
	return resp, nil

}
