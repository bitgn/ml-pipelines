package mlp_api

import (
	"encoding/hex"
	"fmt"
	"mlp/catalog/vo"
)




func notFound(kind vo.ENTITY, uid []byte) *ApiError{

	name := vo.ENTITY_name[int32(kind)]
	return &ApiError{
		Code:StatusCode_NOT_FOUND,
		Message:fmt.Sprintf("%s '%s' not found", name, hex.EncodeToString(uid)),
		SubjectUid:uid,

	}
}

func unknownProjectName(name string) *ApiError{

	typ := vo.ENTITY_name[int32(vo.ENTITY_PROJECT)]
	return &ApiError{
		Code:StatusCode_NOT_FOUND,
		Message:fmt.Sprintf("%s '%s' not found", typ, name),
		SubjectName:name,

	}
}




func badName(entity vo.ENTITY, name string, err error) *ApiError{
	entityName := vo.ENTITY_name[int32(entity)]
	return &ApiError{
		Code:StatusCode_BAD_NAME,
		Message:fmt.Sprintf("Bad %s name '%s'", entityName, name),
		Details:[]string{err.Error()},
		SubjectName:name,
		FieldName:"name",

	}
}







func alreadyExists(entity vo.ENTITY, project_name, name string, id []byte) *ApiError{
	entityName := vo.ENTITY_name[int32(entity)]
	return &ApiError{
		Code:StatusCode_ALREADY_EXISTS,
		Message:fmt.Sprintf("%s named '%s' already exists in %s", entityName, name, project_name),

		SubjectName:name,
		SubjectUid:id,
		ProjectName:project_name,

	}
}


func nameNotFound(entity vo.ENTITY, project_name, name string) *ApiError{
	entityName := vo.ENTITY_name[int32(entity)]
	return &ApiError{
		Code:StatusCode_NOT_FOUND,
		Message:fmt.Sprintf("%s named '%s' not found in %s", entityName, name, project_name),

		SubjectName:name,
		ProjectName:project_name,

	}
}
