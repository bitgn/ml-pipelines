from dataclasses import dataclass
from typing import Callable, Optional

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
import grpc
from . import errors

import google.protobuf.message as pb



class PipelineInput:
    def __init__(self, uid: bytes):
        self.uid= uid

class Context:

    def __init__(self, catalog: rpc.CatalogStub):
        self.catalog = catalog

    def get_project(self, req: api.GetProjectRequest) -> api.ProjectInfoResponse:
        resp: api.ProjectInfoResponse = self._rpc(lambda: self.catalog.GetProject(req))
        return resp

    def create_project(self, req: api.CreateProjectRequest) -> api.ProjectInfoResponse:
        resp: api.ProjectInfoResponse = self._rpc(lambda: self.catalog.CreateProject(req))
        return resp

    def get_job(self, req: api.GetJobRequest) -> api.JobInfoResponse:
        s: api.JobInfoResponse = self._rpc(lambda: self.catalog.GetJob(req))
        return s

    def create_job(self, req: api.CreateJobRequest) -> api.JobInfoResponse:
        s: api.JobInfoResponse = self._rpc(lambda: self.catalog.CreateJob(req))
        return s

    def create_dataset(self, req: api.CreateDatasetRequest) -> api.DatasetInfoResponse:
        s: api.DatasetInfoResponse = self._rpc(lambda: self.catalog.CreateDataset(req))
        return s

    def get_dataset(self, req: api.GetDatasetRequest) -> api.DatasetInfoResponse:
        s: api.DatasetInfoResponse = self._rpc(lambda: self.catalog.GetDataset(req))
        return s

    def get_last_version(self, req: api.GetLastDatasetVersionRequest) -> api.DatasetVersionResponse:
        s: api.DatasetVersionResponse = self._rpc(lambda: self.catalog.GetLastDatasetVersion(req))
        return s

    def start_job_run(self, req: api.StartJobRunRequest) -> api.JobRunInfoResponse:
        s: api.JobRunInfoResponse = self._rpc(lambda: self.catalog.StartJobRun(req))
        return s

    def log_job_run(self, req: api.LogJobRunRequest) -> api.EmptyResponse:
        s: api.EmptyResponse = self._rpc(lambda: self.catalog.LogJobRun(req))
        return s

    def complete_job_run(self, req: api.CompleteJobRunRequest) -> api.EmptyResponse:
        s: api.EmptyResponse = self._rpc(lambda: self.catalog.CompleteJobRun(req))
        return s

    def fail_job_run(self, req: api.FailJobRunRequest) -> api.EmptyResponse:
        s: api.EmptyResponse = self._rpc(lambda: self.catalog.FailJobRun(req))
        return s

    def commit(self, req: api.CommitRequest) -> api.CommitResponse:
        s: api.CommitResponse = self._rpc(lambda: self.catalog.Commit(req))
        return s

    def _rpc(self, callable: Callable[[], pb.Message]):
        try:

            resp = callable()

            if resp.error.code != 0:
                raise errors.from_error(resp.error)

            return resp
        except grpc.RpcError as e:
            raise errors.from_exception(e)
