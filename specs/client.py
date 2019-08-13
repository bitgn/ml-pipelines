import grpc

from api import api_pb2_grpc as rpc
from api import api_pb2 as api
name = "client"


class Client:
    def __init__(self, test: rpc.TestStub, catalog: rpc.CatalogStub):
        self.test = test
        self.catalog = catalog


    def ping(self):
        self.test.Ping(api.PingRequest())

    def create_project(self, project_id: str, project_name: str):

        request = api.CreateProjectRequest(ProjectId=project_id, ProjectName=project_name)
        resp = self.catalog.CreateProject(request)



def connect(url):
    channel = grpc.insecure_channel(url)
    stub = rpc.TestStub(channel)
    catalog = rpc.CatalogStub(channel)

    return Client(stub, catalog)
