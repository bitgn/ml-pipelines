package db


func PutSystem(tx *Tx, val *SystemData){
	tx.PutProto(CreateKey(Range_SYSTEMS, val.Uid), val)
}

func GetSystem(tx *Tx, uid []byte) *SystemData{
	mustUid(uid)


	val := &SystemData{}
	if tx.GetProto(CreateKey(Range_SYSTEMS, uid), val){
		return val
	}
	return nil
}



func PutSystemVersion(tx *Tx, val *SystemVersionData){
	tx.PutProto(CreateKey(Range_SYSTEM_VERSIONS, val.Uid), val)
}

func GetSystemVersion(tx *Tx, uid []byte) *SystemVersionData{
	mustUid(uid)
	val := &SystemVersionData{}

	if tx.GetProto(CreateKey(Range_SYSTEM_VERSIONS, uid), val){
		return val
	}
	return nil
}


func IndexSystemVersion(tx *Tx, service_uid, version_uid []byte, num int32){
	tx.Put(CreateKey(Range_IDX_SYSTEM_VERSIONS, service_uid, int(num)),version_uid)
}
func LookupSystemVersion(tx *Tx, service_uid []byte, version int32)[]byte{
	return tx.Get(CreateKey(Range_IDX_SYSTEM_VERSIONS, service_uid, int(version)))
}





func PutSystemLink(tx *Tx, service_uid []byte, ref *SystemLink){
	tx.PutProto(CreateKey(Range_SYSTEM_LINKS, service_uid, ref.ContainerUid), ref)
}

func ListSystemLinks(tx *Tx, service_uid []byte) []*SystemLink{

	var refs []*SystemLink
	tx.MustScanRange(CreateKey(Range_SYSTEM_LINKS, service_uid), func(k,v []byte){

		ref := &SystemLink{}
		mustUnmarshal(v, ref)
		refs = append(refs, ref)
	})

	return refs

}

