package db

import (
	"github.com/abdullin/lex-go/tuple"
	"github.com/pkg/errors"
)

func AddDataset(tx *Tx, val *DatasetData){
	tx.PutProto(CreateKey(Range_DATASETS, val.DatasetId), val)
	tx.Put(CreateKey(Range_PROJECT_DATASETS, val.ProjectId, val.DatasetId), nil)

}

func GetDataset(tx *Tx, id string) *DatasetData{
	val := &DatasetData{}
	if tx.GetProto(CreateKey(Range_DATASETS, id), val){
		return val
	}
	return nil
}



func ListDatasets(tx *Tx, projectId string) []*DatasetData{
	var vals []*DatasetData
	tx.ScanRange(CreateKey(Range_PROJECT_DATASETS, projectId), func(k,v []byte){

		if tpl, err := tuple.Unpack(k); err != nil {
			panic(errors.Wrap(err, "Failed to unpack"))
		} else {
			datasetId := tpl[2]
			val := &DatasetData{}
			tx.GetProto(CreateKey(Range_DATASETS, datasetId), val)
			vals = append(vals, val)
		}
	})
	return vals
}

func ListAllDatasets(tx *Tx) []*DatasetData{
	var vals []*DatasetData
	tx.ScanRange(CreateKey(Range_DATASETS), func(k,v []byte){
		val := &DatasetData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}
