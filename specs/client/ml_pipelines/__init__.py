name = "MLP"


import grpc


from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc

class Client:
    def __init__(self, catalog: rpc.CatalogStub):
        self.catalog = catalog



    def create_project(self, project_id: str, project_name: str):
        request = api.CreateProjectRequest(ProjectId=project_id, ProjectName=project_name)
        resp = self.catalog.CreateProject(request)
        return resp


    def stats(self):
        request = api.StatRequest()
        return self.catalog.Stat(request)



def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    return Client(catalog)
