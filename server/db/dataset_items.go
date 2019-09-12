package db

import "mlp/catalog/vo"

func PutDatasetHeadItem(tx *Tx, ds_uid []byte, data *DatasetItemData){

	mustUid(ds_uid)
	tx.PutProto(CreateKey(Range_DATASET_HEAD_ITEMS, ds_uid, data.Name), data)
}

func DelDatasetHeadItem(tx *Tx, dataset_uid []byte, name string){
	mustUid(dataset_uid)
	mustName(name)

	tx.MustDel(CreateKey(Range_DATASET_HEAD_ITEMS, dataset_uid, name))
}


func DeleteDatasetHead(tx *Tx, dataset_uid []byte) int{
	mustUid(dataset_uid)

	return tx.MustDelRangeByPrefix(CreateKey(Range_DATASET_HEAD_ITEMS, dataset_uid))
}



func ListDatasetHeadItems(tx *Tx, dataset_uid []byte) []*DatasetItemData{
	var vals []*DatasetItemData
	tx.MustScanRange(CreateKey(Range_DATASET_HEAD_ITEMS,dataset_uid), func(k,v []byte){
		val := &DatasetItemData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}



func PutDatasetItem(tx *Tx, dataset_uid []byte, i *vo.DatasetItem){
	tx.PutProto(CreateKey(Range_DATASET_BINARY_TREE, dataset_uid, i.Name), i)
}
