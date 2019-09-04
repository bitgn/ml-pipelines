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


func PutJobRun(tx *Tx, j *JobRunData){
	tx.PutProto(CreateKey(Range_JOB_RUNS, j.Uid), j)
}


func GetJobRun(tx *Tx, uid []byte) *JobRunData {
	j := &JobRunData{}

	if tx.GetProto(CreateKey(Range_JOB_RUNS, uid), j){
		return j
	}
	return nil
}


