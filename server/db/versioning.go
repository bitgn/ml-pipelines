package db

func GetVersion(tx *Tx) *AppVersionData {
	ver := &AppVersionData{}

	tx.GetProto(CreateKey(Range_VERSION), ver)
	return ver
}

func PutVersion(tx *Tx, ver *AppVersionData){
	tx.PutProto(CreateKey(Range_VERSION), ver)
}


// WipeViewTables cleans up view tables from the DB and returns number of values removed
func WipeViewTables(tx *Tx) int{

	count := 0
	for id, _ := range Range_name{

		r := Range(id)
		switch r {
		case Range_VERSION:
		case Range_Events:
			continue
		default:
			count += tx.MustDelRangeByPrefix(CreateKey(r))
		}

	}
	return count
}
