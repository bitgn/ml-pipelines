from dataclasses import dataclass
from typing import Optional, Callable, List, Union

from google.protobuf.message import Message

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
from . import vo_pb2 as vo
from .cl_project import Project
from .cl_dataset import Dataset
from .cl_bases import Context

from . import errors
import grpc
import google.protobuf.message as pb


class Server():
    def __init__(self, context: Context):
        self.context = context

    def stats(self):

        request = api.StatRequest()
        return self.context.stat()

    def reset_db(self):
        self.context.reset()

class Client():

    def __init__(self, context: Context):
        self.context = context
        self.server = Server(context)





    def get_project(self, project_id) -> Project:

        r = api.GetProjectRequest(project_id=project_id)
        resp = self.context.get_project(r)
        return Project(self.context, project_id=project_id)


    def __getitem__(self, key:str):
        return self.get_project(key)


    def add_project(self, project_id: str) -> Project:


        request = api.CreateProjectRequest(
            project_id=project_id,)

        resp = self.context.create_project(request)
        return Project(self.context, project_id=project_id)





def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    ctx = Context(catalog)

    return Client(ctx)
