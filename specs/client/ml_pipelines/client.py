from typing import Optional

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
from . import vo_pb2 as vo

import grpc



class ArgumentError(ValueError):
    pass

class Client:
    def __init__(self, catalog: rpc.CatalogStub):
        self.catalog = catalog

    def create_project(self, project_id: str,
                       project_name: Optional[str] = None):
        delta = vo.ProjectMetadataDelta()

        if project_name:
            delta.name = project_name
            delta.name_set = True

        request = api.CreateProjectRequest(
            project_id=project_id,
            meta=delta)
        try:

            resp = self.catalog.CreateProject(request)
            return resp
        except grpc.RpcError as e:
            code: grpc.StatusCode = e.code()

            if code == grpc.StatusCode.INVALID_ARGUMENT:
                print(f"Got error {e}")

                raise ArgumentError() from e

            raise RuntimeError() from e
            



    def stats(self):
        request = api.StatRequest()
        return self.catalog.Stat(request)


def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    return Client(catalog)
