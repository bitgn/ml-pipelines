package db

import "log"

func mustUid(uid []byte){
	if len(uid)==0{
		log.Panicln("UID can't be nil")
	}
}


func PutDataset(tx *Tx, val *DatasetData){
	tx.PutProto(CreateKey(Range_DATASETS, val.Uid), val)
}

func GetDataset(tx *Tx, uid []byte) *DatasetData{
	mustUid(uid)


	val := &DatasetData{}
	if tx.GetProto(CreateKey(Range_DATASETS, uid), val){
		return val
	}
	return nil
}


func ListDatasets(tx *Tx) []*DatasetData{
	var vals []*DatasetData
	tx.MustScanRange(CreateKey(Range_DATASETS), func(k,v []byte){
		val := &DatasetData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}
