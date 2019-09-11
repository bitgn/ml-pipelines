import sys
import platform
import traceback
import locale
import struct
from typing import Optional, Union, List
from .deps import get_package_info
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

    # TODO: cache this thing per run
    def log_version_info(self):
        (sysname, nodename, release, version, machine, processor) = platform.uname()

        bits = struct.calcsize("P") * 8
        host = [
            ("python", "%d.%d.%d.%s.%s" % sys.version_info[:] + " {0} bit".format(bits)),
            ("OS", "{0} {1}".format(sysname, release)),
            ("machine", machine),
            ("node", platform.node()),
            ("processor", processor),
            ("byteorder", sys.byteorder)
        ]

        host.extend(get_package_info(['pandas', 'numpy', 'bokeh', 'lz4','grpc','google.protobuf']))

        log = "\n".join([k + ": " + str(v) for k,v in host])
        self.log(log)


    def log(self, details: str, title: Optional[str] = None, timestamp: int = 0):

        req = api.LogJobRunRequest(
            uid=self.uid,
            details=details,
            log_title=title,
            timestamp=timestamp,
        )

        self.ctx.log_job_run(req)

    def complete(self, timestamp: int = 0):
        req = api.CompleteJobRunRequest(
            uid=self.uid,
            timestamp=timestamp
        )
        self.ctx.complete_job_run(req)

    def fail(self, message:Optional[str]=None, timestamp: int = 0):

        req = api.FailJobRunRequest(
            uid=self.uid,
            message=message,
            details=traceback.format_exc(),
            timestamp=timestamp
        )
        self.ctx.fail_job_run(req)




class Job:
    def __init__(self, ctx:Context, project_uid, uid, name):
        self.name = name
        self.uid = uid
        self.project_uid = project_uid
        self.ctx = ctx

    def start_run(
            self,
            inputs:List[Union[DatasetVersionId, SystemId]],
            title: Optional[str]=None,
            timestamp: int = 0
    ) -> JobRun:

        req = api.StartJobRunRequest(
            project_uid=self.project_uid,
            job_uid=self.uid,
            title=title,
            timestamp=timestamp,
        )

        for i in inputs:
            if isinstance(i, DatasetVersionId):
                req.inputs.append(vo.JobRunInput(uid=i.uid, type=vo.JobRunInput.Type.DatasetVer))
                continue
            if isinstance(i, SystemId):
                req.inputs.append(vo.JobRunInput(uid=i.uid, type=vo.JobRunInput.Type.System))
                continue
            raise ValueError(f"Unexpected job run input {i.__class__}")
        resp = self.ctx.start_job_run(req)


        return JobRun(self.ctx, self.project_uid, self.uid, resp.uid)


