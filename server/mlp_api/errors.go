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





func alreadyExists(entity vo.ENTITY, name string, id []byte) *ApiError{
	entityName := vo.ENTITY_name[int32(entity)]
	return &ApiError{
		Code:StatusCode_ALREADY_EXISTS,
		Message:fmt.Sprintf("%s named '%s' already exists", entityName, name),

		SubjectName:name,
		SubjectUid:id,

	}
}

