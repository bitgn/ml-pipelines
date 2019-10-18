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


    def __init__(self, status_code:api.StatusCode, message:str, detail_code:str=None,
                 detail_args:List[str]=(), details:List[str]=()):
        super(ClientError, self).__init__(message)
        self.message = message
        self.details = details
        self.status_code = status_code
        self.detail_code=detail_code
        self.detail_args=detail_args


    def __str__(self):
        msg =  f'{self.status_code} {self.message}'
        if self.details:
            msg += f" {self.details}"
        if self.detail_code:
            msg += f"{self.detail_code}({self.detail_args})"
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
        e.detail_code,
        e.detail_args,
        e.details
        )


