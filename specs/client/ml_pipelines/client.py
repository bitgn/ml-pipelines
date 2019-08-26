from typing import Optional, Callable

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

    def create_project(self, name: str, title: Optional[str] = None):
        delta = vo.ProjectMetadataDelta()

        if title:
            delta.title = title
            delta.title_set = True

        request = api.CreateProjectRequest(
            name=name,
            meta=delta)

        return self._rpc(lambda: self.catalog.CreateProject(request))



    def stats(self):
        request = api.StatRequest()
        return self.catalog.Stat(request)


def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    return Client(catalog)
