from typing import Optional, List, Any

from .cl_bases import *
from . import vo_pb2 as vo



class Dataset:
    def __init__(self, ctx: Context, project_id:str, dataset_id:str):
        self.project_id  = project_id
        self.dataset_id = dataset_id
        self.ctx = ctx
