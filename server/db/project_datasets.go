package db

import (
	"github.com/abdullin/lex-go/tuple"
	"github.com/pkg/errors"
)

func AddDatasetToProject(tx *Tx, projectId []byte, datasetId []byte){
	tx.Put(CreateKey(Range_IDX_PROJECT_DATASETS, projectId, datasetId), nil)
}


func ListDatasetsFromProject(tx *Tx, projectId []byte) []*DatasetData{
	var ids [][]byte
	tx.MustScanRange(CreateKey(Range_IDX_PROJECT_DATASETS, projectId), func(k,v []byte){

		if tpl, err := tuple.Unpack(k); err != nil {
			panic(errors.Wrap(err, "Failed to unpack"))
		} else {
			datasetId := tpl[2].([]byte)
			ids = append(ids, datasetId)
		}
	})

	vals := make([]*DatasetData, len(ids))

	for i, id := range ids {
		val := &DatasetData{}
		tx.GetProto(CreateKey(Range_DATASETS, id), val)
		vals[i] =val
	}

	return vals
}



func AddSystemToProject(tx *Tx, projectId []byte, serviceId []byte){
	tx.Put(CreateKey(Range_IDX_PROJECT_SYSTEMS, projectId, serviceId), nil)
}


func ListProjectSystems(tx *Tx, projectId []byte) [][]byte{
	var ids [][]byte
	tx.MustScanRange(CreateKey(Range_IDX_PROJECT_SYSTEMS, projectId), func(k,v []byte){

		if tpl, err := tuple.Unpack(k); err != nil {
			panic(errors.Wrap(err, "Failed to unpack"))
		} else {
			datasetId := tpl[2].([]byte)
			ids = append(ids, datasetId)
		}
	})

	return ids
}




func AddJobToProject(tx *Tx, projectId []byte, uid []byte){
	tx.Put(CreateKey(Range_IDX_PROJECT_JOBS, projectId, uid), nil)
}


func ListProjectJobs(tx *Tx, projectId []byte) [][]byte{
	var ids [][]byte
	tx.MustScanRange(CreateKey(Range_IDX_PROJECT_JOBS, projectId), func(k,v []byte){

		if tpl, err := tuple.Unpack(k); err != nil {
			panic(errors.Wrap(err, "Failed to unpack"))
		} else {
			datasetId := tpl[2].([]byte)
			ids = append(ids, datasetId)
		}
	})

	return ids
}
