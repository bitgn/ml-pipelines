package db



func PutDataset(tx *Tx, val *DatasetData){
	tx.PutProto(CreateKey(Range_DATASETS, val.ProjectId, val.DatasetId), val)
}

func GetDataset(tx *Tx, projectId, datasetId string) *DatasetData{
	val := &DatasetData{}
	if tx.GetProto(CreateKey(Range_DATASETS, projectId, datasetId), val){
		return val
	}
	return nil
}


func ListDatasets(tx *Tx, projectId string) []*DatasetData{
	var vals []*DatasetData
	tx.MustScanRange(CreateKey(Range_DATASETS, projectId), func(k,v []byte){
		val := &DatasetData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}

func ListAllDatasets(tx *Tx) []*DatasetData{
	var vals []*DatasetData
	tx.MustScanRange(CreateKey(Range_DATASETS), func(k,v []byte){
		val := &DatasetData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}
