from typing import List, Sequence

import grpc

from . import mlp_api_pb2 as api

def from_exception(e: grpc.RpcError):
    code, msg = e.code().value

    ctor = MAP.get(code, ClientError)
    return ctor(code, str(e))




class ClientError(Exception):
    """Base class for all exceptions raised by this client"""

    status_code = None


    def __init__(self, status_code:api.StatusCode, message:str, subject_uid:bytes=None,
                 subject_name:str=None, details:List[str]=(), project_name: str=None,
                 project_uid: str=None,
                 method_name: str=None):
        super(ClientError, self).__init__(message)
        self.method_name = method_name
        self.message = message
        self.details = details
        self.status_code = status_code
        self.subject_uid=subject_uid
        self.subject_name=subject_name
        self.project_name = project_name
        self.project_uid = project_uid


    def __str__(self):
        msg =  f'{self.status_code} {self.message}'
        if self.details:
            msg += f" {self.details}"
        if self.method_name:
            msg += f" method {self.method_name}"
        return msg




class InvalidArgument(ClientError):
    pass
class AlreadyExists(ClientError):
    pass
class BadName(InvalidArgument):
    pass

class NotFound(InvalidArgument):
    pass

class Unavailable(ClientError):
    pass

MAP= {
    api.StatusCode.INVALID_ARGUMENT: InvalidArgument,
    api.StatusCode.ALREADY_EXISTS:AlreadyExists,
    api.StatusCode.BAD_NAME:BadName,
    api.StatusCode.NOT_FOUND:NotFound,
    api.StatusCode.UNAVAILABLE:Unavailable,
}

MAP = {int(k): v for k,v in MAP.items()}

def from_error(e: api.ApiError) -> ClientError:
    ctor = MAP.get(e.code, ClientError)
    return ctor(
        e.code,
        e.message,
        e.subject_uid,
        e.subject_name,
        list(e.details),
        e.project_name,
        e.project_uid,
        e.method_name)


