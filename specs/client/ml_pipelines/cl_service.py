from typing import Optional, List, Any, Union

from .cl_bases import *
from . import vo_pb2 as vo


class Service(ServiceId):
    def __init__(self, ctx: Context, project_uid: bytes, uid: bytes, name: str):
        super().__init__(uid)
        self.name = name

        self.project_uid = project_uid
        self.ctx = ctx

    def add_version(self, title:Optional[str]=None, inputs:List[Union[ServiceId, JobRunId]]=(), outputs:List[ServiceId]=()):
        pass
