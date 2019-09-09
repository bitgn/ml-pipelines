from typing import Optional, List, Any, Union

from .cl_bases import *
from . import vo_pb2 as vo


class System(SystemId):
    def __init__(self, ctx: Context, project_uid: bytes, uid: bytes, name: str):
        super().__init__(uid)
        self.name = name

        self.project_uid = project_uid
        self.ctx = ctx

    def add_version(self, title:Optional[str]=None, inputs:List[Union[SystemId, JobRunId]]=(), outputs:List[SystemId]=()):


        api_inputs = []

        for x in inputs:
            if isinstance(x, SystemId):
                api_inputs.append(vo.SystemVersionInput(uid=x.uid, type=vo.SystemVersionInput.Type.System))
                continue
            if isinstance(x, JobRunId):
                api_inputs.append(vo.SystemVersionInput(uid=x.uid, type=vo.SystemVersionInput.Type.JobRun))
                continue

            raise ValueError(f'Unexpected input {x.__class__}')

        api_outputs = []
        for x in outputs:
            if isinstance(x, SystemId):
                api_outputs.append(vo.SystemVersionOutput(uid=x.uid, type=vo.SystemVersionOutput.Type.System))
                continue

            raise ValueError(f'Unexpected outputs {x.__class__}')

        r = api.AddSystemVersionRequest(
            system_uid=self.uid,
            title=title,
            inputs=api_inputs,
            outputs=api_outputs,
        )
        self.ctx.add_system_version(r)

