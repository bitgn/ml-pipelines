package db

import "encoding/binary"


func NextActivityCounter(tx *Tx) (counter uint64) {
	counter_bytes := tx.Get(CreateKey(Range_ACTIVITY_COUNT))
	if counter_bytes != nil {
		counter = binary.BigEndian.Uint64(counter_bytes)

	} else {
		counter_bytes = make([]byte, 8)
	}

	counter +=1

	binary.BigEndian.PutUint64(counter_bytes, counter)

	tx.Put(CreateKey(Range_ACTIVITY_COUNT), counter_bytes)
	return counter
}

func AddDatasetActivity(tx *Tx, val *DatasetActivity){
	// reverse timestamp
	tx.PutProto(CreateKey(Range_DATASET_ACTIVITY, val.ProjectId, val.DatasetId, -val.UpdateTimestamp, val.Id), val)
}


func AddGlobalActivity(tx *Tx, val *DatasetActivity){
	// reverse timestamp
	tx.PutProto(CreateKey(Range_GLOBAL_ACTIVITIES, -val.UpdateTimestamp, val.Id), val)
}
func AddGlobalProblem(tx *Tx, val *DatasetActivity){
	// reverse timestamp
	tx.PutProto(CreateKey(Range_GLOBAL_PROBLEMS, -val.UpdateTimestamp, val.Id), val)
}

func ListDatasetActivities(tx *Tx, projectId, datasetId string)(vals []*DatasetActivity) {
	tx.MustScanRange(CreateKey(Range_DATASET_ACTIVITY,projectId, datasetId), func(k, v []byte) {
		val := &DatasetActivity{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}

func ListGlobalActivities(tx *Tx) (vals []*DatasetActivity){
	tx.MustScanRange(CreateKey(Range_GLOBAL_ACTIVITIES), func(k, v []byte) {
		val := &DatasetActivity{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}
