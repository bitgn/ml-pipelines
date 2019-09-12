package db

import (
	"github.com/abdullin/lex-go/tuple"
	"github.com/pkg/errors"
)

func PutDatasetVersion(tx *Tx, val *DatasetVersionData){
	tx.PutProto(CreateKey(Range_DATASET_VERSIONS, val.Uid), val)
}

func GetDatasetVersion(tx *Tx, uid []byte) *DatasetVersionData{
	val := &DatasetVersionData{}

	if tx.GetProto(CreateKey(Range_DATASET_VERSIONS, uid), val){
		return val
	}

	return nil
}


func AddVersionToDataset(tx *Tx, datasetUid, versionUid []byte){
	tx.Put(CreateKey(Range_IDX_DATASET_VERSIONS, datasetUid, versionUid), nil)
}


func GetLastDatasetVersion(tx *Tx, datasetUid []byte) *DatasetVersionData {



	// slow

	vers := ListVersionsFromDataset(tx, datasetUid)
	if len(vers)>0 {
		return vers[len(vers)-1]
	}
	return nil

}

func ListVersionsFromDataset(tx *Tx, datasetUid []byte) []*DatasetVersionData{
	var ids [][]byte
	tx.MustScanRange(CreateKey(Range_IDX_DATASET_VERSIONS, datasetUid), func(k,v []byte){

		if tpl, err := tuple.Unpack(k); err != nil {
			panic(errors.Wrap(err, "Failed to unpack"))
		} else {
			datasetId := tpl[2].([]byte)
			ids = append(ids, datasetId)
		}
	})

	vals := make([]*DatasetVersionData, len(ids))

	for i, id := range ids {
		vals[i] =GetDatasetVersion(tx, id)


	}

	return vals
}






