from typing import Optional, Union, List

from .cl_bases import Context
from . import vo_pb2 as vo


from .cl_dataset import *


class JobRun(JobRunId):
    def __init__(self, ctx: Context, project_uid, job_uid, uid):
        super().__init__(uid)
        self.uid = uid
        self.job_uid = job_uid
        self.project_uid = project_uid
        self.ctx = ctx

    def log(self, details: str, title: Optional[str] = None):

        req = api.LogJobRunRequest(
            uid=self.uid,
            details=details,
            log_title=title
        )

        self.ctx.log_job_run(req)

    def complete(self):
        req = api.CompleteJobRunRequest(
            uid=self.uid
        )
        self.ctx.complete_job_run(req)

    def fail(self, e: Exception):
        req = api.FailJobRunRequest(
            uid=self.uid,
            details=str(e)
        )
        self.ctx.fail_job_run(req)




class Job:
    def __init__(self, ctx:Context, project_uid, uid, name):
        self.name = name
        self.uid = uid
        self.project_uid = project_uid
        self.ctx = ctx

    def start_run(self, inputs:List[Union[DatasetVersionId, ServiceId]],title: Optional[str]=None) -> JobRun:

        req = api.StartJobRunRequest(
            project_uid=self.project_uid,
            job_uid=self.uid,
            title=title,

        )

        for i in inputs:
            if isinstance(i, DatasetVersionId):
                req.inputs.append(vo.JobRunInput(uid=i.uid, type=vo.JobRunInput.Type.DatasetVer))
                continue
            if isinstance(i, ServiceId):
                req.inputs.append(vo.JobRunInput(uid=i.uid, type=vo.JobRunInput.Type.Service))
                continue
            raise ValueError(f"Unexpected job run input {i.__class__}")
        resp = self.ctx.start_job_run(req)


        return JobRun(self.ctx, self.project_uid, self.uid, resp.uid)


