from typing import List, Sequence

import grpc

from . import mlp_api_pb2 as api





class ClientError(Exception):
    """Base class for all exceptions raised by this client"""

    status_code = None


    def __init__(self, status_code:api.StatusCode, message:str, subject_uid:bytes,
                 subject_name:str,details:List[str]):
        super(ClientError, self).__init__(message)
        self.message = message
        self.details = details
        self.status_code = status_code
        self.subject_uid=subject_uid
        self.subject_name=subject_name


    def __str__(self):
        return f'{self.status_code} {self.message} {self.details}'




class InvalidArgument(ClientError):
    pass
class AlreadyExists(ClientError):
    pass
class BadName(InvalidArgument):
    pass

class NotFound(InvalidArgument):
    pass



MAP= {
    api.StatusCode.INVALID_ARGUMENT: InvalidArgument,
    api.StatusCode.ALREADY_EXISTS:AlreadyExists,
    api.StatusCode.BAD_NAME:BadName,
    api.StatusCode.NOT_FOUND:NotFound,

}

def from_error(e: api.ApiError) -> ClientError:
    ctor = MAP.get(e.code, ClientError)
    return ctor(e.code, e.message, e.subject_uid, e.subject_name, list(e.details))


def from_exception(e: grpc.RpcError):
    ctor = MAP.get(e.code(), ClientError)
    return ctor(e.code(), str(e), None, None)
