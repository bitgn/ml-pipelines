package mlp_api



func (r *DatasetInfoResponse) err(e *ApiError) (*DatasetInfoResponse, error){
	r.Error = e
	return r, nil
}


func (r *ListDatasetsResponse) err(e *ApiError) (*ListDatasetsResponse, error){
	r.Error = e
	return r, nil
}



func (r *ListProjectsResponse) err(e *ApiError) (*ListProjectsResponse, error){
	r.Error = e
	return r, nil
}

func (r *EmptyResponse) err(e *ApiError) (*EmptyResponse, error){

	r.Error = e
	return r, nil

}
