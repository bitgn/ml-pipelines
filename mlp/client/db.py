import json
from collections import OrderedDict
from pathlib import Path
from typing import Any, Union, List, Optional, Callable, Tuple

from google.protobuf.message import Message

from proto import dto_pb2 as dto

import base64
from google.protobuf import message


class DB:

    def __init__(self, loc: Path):
        self.loc = loc

    def _load(self):
        if not self.loc.exists():
            return {}

        d = OrderedDict()

        with self.loc.open('r') as f:

            for line in f:
                line = line.strip().split('->', 1)
                d[line[0]] = line[1]
        return d

    def _save(self, d):
        ordered = sorted(d.items(), key=lambda t: t[0])
        with self.loc.open('w') as f:
            for k, v in ordered:
                f.write(f'{k}->{v}\n')


    def add_str(self, key: str, value):

        if len(value)==0:
            raise ValueError(f"Empty string {value} for key {key}")
        d = self._load()
        d[key] = value
        self._save(d)

    def add(self, key: str, value: message.Message):
        val = base64.b32encode(value.SerializeToString()).decode()

        #print(f"Save '{key}' with '{value}'({type(value)}) to '{val}'")
        self.add_str(key,val)

    def get(self, key: str, factory: Callable[...,Message]):
        val = self._load().get(key, None)

        #print(f'Get {key} got {val}')

        if val:
            bts = base64.b32decode(val.encode())
            d = factory()

            d.ParseFromString(bts)
            #print(f'Deserialize {type(d)} into "{d}"')
            return d
        return None


    def get_str(self, key: str, d = None) -> str:
        val = self._load().get(key, d)
        return val

    def key_exists(self, key: str):
        return key in self._load()

    def __get_range_str(self, prefix: str):
        d = self._load()

        def f(t):
            return t[0].startswith(prefix)

        return OrderedDict(filter(f, d.items())).items()

    def __get_range(self, key:str, factory: Callable[...,Message]) -> List[Tuple[str,Message]]:

        list = []
        for k, val in self.__get_range_str(key):
            bts = base64.b32decode(val.encode())
            d = factory()
            d.ParseFromString(bts)
            list.append((k, d))
        return list


    def add_project(self, id: str, val: dto.ProjectData):
        self.add(f'project:{id}:', val)
        self.add_str(f'project-by-name:{val.name}:', id)

    def get_project_by_name(self, name: str) -> Union[dto.ProjectData, None]:
        id = self.get_str(f'project-by-name:{name}:')
        print(f"Found id: '{id}'")
        if id:
            return self.get_project(id)
        return None

    def project_by_name_exists(self, name: str)-> bool:
        return self.key_exists(f'project-by-name:{name}:')


    def get_project(self, id: str) -> Union[dto.ProjectData,None]:
        return self.get(f'project:{id}:', dto.ProjectData)

    def list_projects(self) -> List[dto.ProjectData]:
        return [v for k,v in  self.__get_range('project:', dto.ProjectData)]



    def add_dataset(self, val: dto.DatasetData):
        if not val.id:
            raise ValueError("Null dataset id")
        self.add(f'dataset:{val.id}:', val)
        self.add_str(f'project-datasets:{val.project_id}:{val.id}:', val.id)

    def get_datasets_by_project_id(self, project_id: str) -> List[dto.DatasetData]:
        return [self.get_dataset(val) for k, val in self.__get_range_str(f'project-datasets:{project_id}:')]

    def get_datasets(self) -> List[dto.DatasetData]:
        return [self.get_dataset(val) for k, val in self.__get_range_str(f'project-datasets:')]


    def get_dataset(self, key:str) -> Optional[dto.DatasetData]:
        return self.get(f'dataset:{key}:', dto.DatasetData)


    def add_dataset_version(self, val: dto.VersionData):
        self.add(f'dataset-versions:{val.dataset_id}:{val.version:04d}', val)

    def list_dataset_versions(self, dataset_id: str) -> List[dto.VersionData]:
        search = self.__get_range(f'dataset-versions:{dataset_id}:', dto.VersionData)

        list = []
        for k,v in search:
            v: dto.VersionData
            list.append(v)
        return list




