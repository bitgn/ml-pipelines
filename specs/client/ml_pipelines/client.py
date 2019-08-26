from typing import Optional, Callable, List

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

    def create_project(self, name: str, title: Optional[str] = None) -> api.CreateProjectResponse:
        delta = vo.ProjectMetadataDelta()

        if title:
            delta.title = title
            delta.title_set = True

        request = api.CreateProjectRequest(
            name=name,
            meta=delta)

        return self._rpc(lambda: self.catalog.CreateProject(request))

    def create_job(
            self, project_uid: bytes, name: str,
            title: Optional[str],
            inputs: List[bytes] = None,
            outputs: List[bytes] = None
    ):

        d = vo.JobMetadataDelta()

        if title:
            d.title = title
            d.title_set = True


        for o in outputs:
            d.outputs.append(vo.JobOutput(target_id=o, type=vo.JobOutput.Type.Dataset))

        for i in inputs:
            d.inputs.append(vo.JobInput(source_id=i, type=vo.JobInput.Type.Dataset))


        r = api.CreateJobRequest(
            project_uid=project_uid,
            name=name,
            meta=d
        )



        return self._rpc(lambda: self.catalog.CreateJob(r))

    def create_dataset(self, project_uid: bytes, name: str,
                       title: Optional[str] = None,
                       data_format: Optional[str] = None,
                       file_count: Optional[int] = None,
                       storage_bytes: Optional[int] = None,
                       record_count: Optional[int] = None,
                       description:Optional[str] = None,
                       location_id:Optional[str] = None,
                       location_uri: Optional[str] = None,
                       update_timestamp: Optional[int] = None) -> api.CreateDatasetResponse:

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
            project_uid=project_uid,
        )

        return self._rpc(lambda: self.catalog.CreateDataset(request))

    def stats(self):
        request = api.StatRequest()
        return self.catalog.Stat(request)


def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    return Client(catalog)
