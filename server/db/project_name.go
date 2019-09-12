package db



func NameProject(tx *Tx, uid []byte, name string){
	tx.Put(CreateKey(Range_PROJECT_NAMES, name), uid)
}

func LookupProject(tx *Tx, name string) []byte{
	return tx.Get(CreateKey(Range_PROJECT_NAMES, name))
}

