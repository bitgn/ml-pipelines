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
	if j.RunNum == 0 {
		panic("RunNUM must be positive")
	}

	tx.PutProto(CreateKey(Range_JOB_RUNS, j.Uid), j)
	tx.Put(CreateKey(Range_IDX_JOB_RUNS, j.JobUid, int(j.RunNum)), j.Uid)
}


func LookupJobRun(tx *Tx, uid []byte, num int32) *JobRunData{
	runUid := tx.Get(CreateKey(Range_IDX_JOB_RUNS, uid, int(num)))
	return GetJobRun(tx, runUid)
}

func GetJobRun(tx *Tx, uid []byte) *JobRunData {
	j := &JobRunData{}

	if tx.GetProto(CreateKey(Range_JOB_RUNS, uid), j){
		return j
	}
	return nil
}


func AddJobRunEvent(tx *Tx, run []byte, num int32,  e *JobRunEvent){
	tx.PutProto(CreateKey(Range_JOB_RUN_HISTORY, run, int(num)), e)
}

func ListJobRunHistory(tx *Tx, run []byte) []*JobRunEvent{
	var vals []*JobRunEvent
	tx.MustScanRange(CreateKey(Range_JOB_RUN_HISTORY, run), func(k,v[]byte){
		val := &JobRunEvent{}
		mustUnmarshal(v, val)
		vals = append(vals, val)
	})
	return vals
}






