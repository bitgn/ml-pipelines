package mlp_api

import (
	"fmt"
)




func projectNotFound(name string) *ApiError{
	return &ApiError{
		Code:StatusCode_NOT_FOUND,
		Message:fmt.Sprintf("Project '%s' not found", name),
		DetailCode:"project_not_found",
		DetailArgs:[]string{name},
	}
}

func datasetNotFound(projectId, datasetId string) *ApiError{
	return &ApiError{
		Code:StatusCode_NOT_FOUND,
		Message:fmt.Sprintf("Dataset '%s/%s' not found", projectId, datasetId),
		DetailCode:"dataset_not_found",
		DetailArgs:[]string{projectId, datasetId},
	}
}






func badName(name string, err error) *ApiError{
	return &ApiError{
		Code:StatusCode_BAD_NAME,
		Message:fmt.Sprintf("Bad name '%s'",name),
		Details:err.Error(),
		DetailCode:"bad_name",
		DetailArgs:[]string{name},

	}
}



func projectAlreadyExists(projectId string) *ApiError{

	return &ApiError{
		Code:StatusCode_ALREADY_EXISTS,
		Message:fmt.Sprintf("Project %s already exists", projectId),
		DetailCode: "project_already_exists",
		DetailArgs:[]string{projectId},


	}
}





func datasetAlreadyExists(projectId, datasetId string) *ApiError{

	return &ApiError{
		Code:StatusCode_ALREADY_EXISTS,
		Message:fmt.Sprintf("Dataset %s/%s already exists", projectId, datasetId),
		DetailCode: "dataset_already_exists",
		DetailArgs:[]string{projectId, datasetId},


	}
}

