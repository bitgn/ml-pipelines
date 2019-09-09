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


        api_inputs = []

        for x in inputs:
            if isinstance(x, ServiceId):
                api_inputs.append(vo.ServiceVersionInput(uid=x.uid, type=vo.ServiceVersionInput.Type.Service))
                continue
            if isinstance(x, JobRunId):
                api_inputs.append(vo.ServiceVersionInput(uid=x.uid, type=vo.ServiceVersionInput.Type.JobRun))
                continue

            raise ValueError(f'Unexpected input {x.__class__}')

        api_outputs = []
        for x in outputs:
            if isinstance(x, ServiceId):
                api_outputs.append(vo.ServiceVersionOutput(uid=x.uid, type=vo.ServiceVersionOutput.Type.Service))
                continue

            raise ValueError(f'Unexpected outputs {x.__class__}')

        r = api.AddServiceVersionRequest(
            service_uid=self.uid,
            title=title,
            inputs=api_inputs,
            outputs=api_outputs,
        )
        self.ctx.add_service_version(r)

