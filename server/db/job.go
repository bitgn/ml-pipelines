package db

func GetJob(tx *Tx, job_uid []byte) *Job {
	j := &Job{}
	if tx.GetProto(CreateKey(Range_JOBS, job_uid), j) {
		return j
	}
	return nil
}

func PutJob(tx *Tx, j *Job){
	tx.PutProto(CreateKey(Range_JOBS, j.Uid), j)
}