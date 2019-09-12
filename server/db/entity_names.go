package db

import "mlp/catalog/vo"

func Name(tx *Tx, project_uid []byte, name string, kind vo.ENTITY, uid []byte){
	tx.PutProto(CreateKey(Range_NAMES_SCOPED, project_uid, name), &EntityRef{
		Uid:uid,
		Kind:kind,
	})
}

func Lookup(tx *Tx, project_uid []byte, name string) *EntityRef{

	r := &EntityRef{}
	if tx.GetProto(CreateKey(Range_NAMES_SCOPED, project_uid, name), r){
		return r
	}
	return nil
}

