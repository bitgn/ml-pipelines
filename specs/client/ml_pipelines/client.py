from dataclasses import dataclass
from typing import Optional, Callable, List

from google.protobuf.message import Message

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
from . import vo_pb2 as vo

from . import errors
import grpc
import google.protobuf.message as pb


class Client:

    def __init__(self, catalog: rpc.CatalogStub):
        self.catalog = catalog

    def _rpc(self, callable: Callable[[], pb.Message]):
        try:

            resp = callable()

            if resp.error.code != 0:
                raise errors.from_error(resp.error)

            return resp
        except grpc.RpcError as e:
            raise errors.from_exception(e)

    def create_project(self, name: str, title: Optional[str] = None) -> 'ProjectCreated':
        delta = vo.ProjectMetadataDelta()

        if title:
            delta.title = title
            delta.title_set = True

        request = api.CreateProjectRequest(
            name=name,
            meta=delta)

        response: api.CreateProjectResponse = self._rpc(lambda: self.catalog.CreateProject(request))
        return ProjectCreated(uid=response.uid)

    def add_dataset_version(self,
                            project_name: str,
                            dataset_name: str,
                            items:List['DatasetItem'],
                            title: Optional[str]=None,
                            inputs:List[bytes]=(),
                            parent_uid:Optional[bytes]=None):

        prj = self._lookup_dataset(project_name, dataset_name)

        vo.DatasetItem()


        req = api.AddDatasetVersionRequest(
            dataset_uid=prj.dataset_uid,
            title=title,
            inputs=inputs,
            parent_uid=parent_uid,
        )

        response: api.AddDatasetVersionResponse = self._rpc(lambda: self.catalog.AddDatasetVersion(req))

        return response

    def create_job(
            self, project_name: str, name: str,
            title: Optional[str] = None,
            inputs: List[bytes] = (),
            outputs: List[bytes] = ()
    ) -> 'JobCreated':

        prj = self._lookup_project(project_name)

        d = vo.JobMetadataDelta()

        if title:
            d.title = title
            d.title_set = True


        for o in outputs:
            d.outputs.append(vo.JobOutput(target_id=o, type=vo.JobOutput.Type.Dataset))

        for i in inputs:
            d.inputs.append(vo.JobInput(source_id=i, type=vo.JobInput.Type.Dataset))


        r = api.CreateJobRequest(
            project_uid=prj.uid,
            name=name,
            meta=d
        )



        response: api.CreateJobResponse = self._rpc(lambda: self.catalog.CreateJob(r))
        return JobCreated(uid=response.uid)

    def _lookup_project(self, project_name: str):
        prj: api.LookupProjectResponse = self._rpc(
            lambda: self.catalog.LookupProject(api.LookupProjectRequest(name=project_name)))
        return prj

    def _lookup_dataset(self, project_name:str, dataset_name: str) -> api.LookupDatasetResponse:
        request = api.LookupDatasetRequest(
            project_name=project_name,
            dataset_name=dataset_name)

        prj: api.LookupDatasetResponse = self._rpc(lambda: self.catalog.LookupDataset(request))
        return prj

    def create_dataset(self, project_name: str,
                       name: str,
                       title: Optional[str] = None,
                       data_format: Optional[str] = None,
                       file_count: Optional[int] = None,
                       storage_bytes: Optional[int] = None,
                       record_count: Optional[int] = None,
                       description:Optional[str] = None,
                       location_id:Optional[str] = None,
                       location_uri: Optional[str] = None,
                       update_timestamp: Optional[int] = None) -> 'DatasetCreated':




        prj = self._lookup_project(project_name=project_name)

        d = vo.DatasetMetadataDelta()

        if title:
            d.title = title
            d.title_set = True

        if data_format:
            d.data_format = data_format
            d.data_format_set = True

        if file_count:
            d.file_count = file_count
            d.file_count_set = True

        if storage_bytes:
            d.storage_bytes = storage_bytes
            d.storage_bytes_set = True

        if record_count:
            d.record_count = record_count
            d.record_count_set = True

        if description:
            d.description = description
            d.description_set = True

        if location_id:
            d.location_id = location_id
            d.location_id_set = True

        if location_uri:
            d.location_uri = location_uri
            d.location_id_set = True

        if update_timestamp:
            d.update_timestamp = update_timestamp
            d.update_timestamp_set = True



        request = api.CreateDatasetRequest(
            name=name,
            meta=d,
            project_uid=prj.uid,
        )

        response: api.CreateDatasetResponse = self._rpc(lambda: self.catalog.CreateDataset(request))
        return DatasetCreated(uid=response.uid)

    def stats(self):
        request = api.StatRequest()
        return self.catalog.Stat(request)


@dataclass
class ProjectCreated:
    uid: bytes

    def __str__(self):
        return f'Project {self.uid.hex()}'

@dataclass
class DatasetCreated:
    uid: bytes

    def __str__(self):
        return f'Dataset {self.uid.hex()}'

@dataclass
class DatasetItem:
    name: str
    storage_bytes: int
    records: int
    def __init__(self, name:str, storage_bytes:int, records:Optional[int]=0):
        self.name=name
        self.storage_bytes=storage_bytes
        if records:
            self.records=records





@dataclass
class JobCreated:
    uid: bytes

    def __str__(self):
        return f'Job {self.uid.hex()}'


def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    return Client(catalog)
