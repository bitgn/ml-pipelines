package mlp_api



func (r *DatasetInfoResponse) err(e *ApiError) (*DatasetInfoResponse, error){

	r.Error = e
	return r, nil

}
func (r *ServiceInfoResponse) err(e *ApiError) (*ServiceInfoResponse, error){

	r.Error = e
	return r, nil

}
func (r *DatasetVersionResponse) err(e *ApiError) (*DatasetVersionResponse, error){

	r.Error = e
	return r, nil

}

func (r *JobRunInfoResponse) err(e *ApiError) (*JobRunInfoResponse, error){

	r.Error = e
	return r, nil

}


func (r *EmptyResponse) err(e *ApiError) (*EmptyResponse, error){

	r.Error = e
	return r, nil

}




func (r *CommitResponse) err(e *ApiError) (*CommitResponse, error){

	r.Error = e
	return r, nil

}