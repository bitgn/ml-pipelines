package db


func GetStats(tx *Tx) *TenantStats{
	d := &TenantStats{}
	tx.GetProto(CreateKey(Range_STATS), d)
	return d
}

func SetStats(tx *Tx, d *TenantStats){
	tx.PutProto(CreateKey(Range_STATS), d)
}



