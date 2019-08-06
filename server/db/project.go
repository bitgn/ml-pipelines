package db

func AddProject(tx *Tx, val *ProjectData){
	tx.PutProto(CreateKey(Range_PROJECTS, val.Id), val)
}

func GetProject(tx *Tx, id string) (*ProjectData){
	d := &ProjectData{}
	if tx.GetProto(CreateKey(Range_PROJECTS, id), d) {
		return d
	}
	return nil
}


func ListProjects(tx *Tx) []*ProjectData{
	var vals []*ProjectData
	tx.ScanRange(CreateKey(Range_PROJECTS), func(k,v []byte){
		val := &ProjectData{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}
