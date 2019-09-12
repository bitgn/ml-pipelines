from dataclasses import dataclass
from typing import Optional, Callable, List, Union

from google.protobuf.message import Message

from . import mlp_api_pb2 as api
from . import mlp_api_pb2_grpc as rpc
from . import vo_pb2 as vo
from .cl_project import Project
from .cl_job import Job
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
        return self.context.stat(request)

    def reset_db(self):
        self.context.reset()

class Client():

    def __init__(self, context: Context):
        self.context = context
        self.server = Server(context)





    def get_project(self, name) -> Project:

        r = api.GetProjectRequest(name=name)
        resp = self.context.get_project(r)
        return Project(self.context, resp.uid)


    def __getitem__(self, key:str):
        return self.get_project(key)


    def add_project(self, name: Optional[str], title: Optional[str] = None) -> Project:

        delta = vo.ProjectMetadataDelta()

        if title:
            delta.title = title
            delta.title_set = True

        request = api.CreateProjectRequest(
            name=name,
            meta=delta)

        resp = self.context.create_project(request)
        return Project(self.context, resp.uid)





def connect(url):
    channel = grpc.insecure_channel(url)
    catalog = rpc.CatalogStub(channel)

    ctx = Context(catalog)

    return Client(ctx)
