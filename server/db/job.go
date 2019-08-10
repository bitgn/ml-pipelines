package db

func GetJob(tx *Tx, job_id string) *Job {
	j := &Job{}
	if tx.GetProto(CreateKey(Range_JOBS, job_id), j) {
		return j
	}
	return nil
}

func PutJob(tx *Tx, j *Job){
	tx.PutProto(CreateKey(Range_JOBS, j.JobId), j)
}