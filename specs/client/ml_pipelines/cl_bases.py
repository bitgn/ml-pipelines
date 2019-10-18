from dataclasses import dataclass
from typing import Callable, Optional

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
import grpc
from . import errors
import os
import google.protobuf.message as pb



class EntityId:
    def __init__(self, uid: bytes):
        self.uid= uid

class JobRunId (EntityId):
    pass


class DatasetVersionId(EntityId):
    pass

class SystemId(EntityId):
    pass



class Context:

    def __init__(self, catalog: rpc.CatalogStub):
        self.catalog = catalog

    def create_project(self, req: api.CreateProjectRequest) -> api.ProjectInfoResponse:
        resp: api.ProjectInfoResponse = self._rpc(lambda: self.catalog.CreateProject(req))
        return resp

    def get_project(self, req: api.GetProjectRequest) -> api.ProjectInfoResponse:
        resp: api.ProjectInfoResponse = self._rpc(lambda: self.catalog.GetProject(req))
        return resp

    def list_projects(self, req: api.ListProjectsRequest) -> api.ListProjectsResponse:
        resp: api.ListProjectsResponse = self._rpc(lambda : self.catalog.ListProjects(req))
        return resp


    def create_dataset(self, req: api.CreateDatasetRequest) -> api.DatasetInfoResponse:
        s: api.DatasetInfoResponse = self._rpc(lambda: self.catalog.CreateDataset(req))
        return s

    def get_dataset(self, req: api.GetDatasetRequest) -> api.DatasetInfoResponse:
        s: api.DatasetInfoResponse = self._rpc(lambda: self.catalog.GetDataset(req))
        return s

    def list_datasets(self, req:api.ListDatasetsRequest) -> api.ListDatasetsResponse:
        s:api.ListDatasetsResponse = self._rpc(lambda : self.catalog.ListDatasets(req))
        return s

    def update_dataset(self, req: api.UpdateDatasetRequest) -> api.UpdateDatasetRequest:
        s:api.UpdateDatasetRequest = self._rpc(lambda: self.catalog.UpdateDataset(req))
        return s

    def reset(self):
        self._rpc(lambda : self.catalog.Reset(api.ResetRequest()))

    def stat(self):
        s:api.StatResponse = self._rpc(lambda : self.catalog.Stat(api.StatRequest()))
        return s

    def _rpc(self, callable: Callable[[], pb.Message]):
        try:

            resp = callable()

            if not hasattr(resp,"error"):

                raise ValueError(print(f'{resp.__class__} has no error attribute'))

            if resp.error.code != 0:
                raise errors.from_error(resp.error)

            return resp
        except grpc.RpcError as e:
            raise errors.from_exception(e)
