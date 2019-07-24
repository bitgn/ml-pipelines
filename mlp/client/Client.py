import gzip
import hashlib
import io
import pandas as pd
import shutil
from datetime import datetime

from typing import List, Dict, Union, Any

from attr import dataclass

from client.db import DB
from client.pretty import date_fmt, sizeof_fmt

from client.file_store import *

from client.Table import Table
from proto import dto_pb2 as dto


class Client:
    """
    Provide access to the BitGN-MLP platform
    """

    def __init__(self, key: str):

        # temporary hack to look good on the demo
        import base64
        key = base64.b64decode(key[2:].encode('ascii')).decode('ascii')

        self._folder = Path(key).resolve()
        self._db = DB(self._folder / 'db.json')

    def _new_id(self):
        return hex(int(datetime.utcnow().timestamp() * 100))[2:]

    def project_exists(self, name: str):
        return self._db.project_by_name_exists(name)

    def get_project_by_name(self, name: str) -> Union[dto.ProjectData, None]:
        return self._db.get_project_by_name(name)

    def create_project(self, name: str) -> 'ProjectID':
        id = 'prj-' + self._new_id()
        self._db.add_project(id, dto.ProjectData(name=name, id=id))
        return ProjectID(name=name, id=id)



    def read_dataset(self, dataset_id: Union['DatasetID',str], version=-1, chunk=-1):
        if isinstance(dataset_id, DatasetID):
            dataset_id = dataset_id.id
        dataset = self._db.get_dataset(dataset_id)

        versions = self._db.list_dataset_versions(dataset_id)

        ver = versions[version]

        chunk = ver.chunks[chunk]
        sha = chunk.hash

        file = self._folder / sha


        reader = file.open('rb')

        unzip = gzip.GzipFile(fileobj=reader, mode='rb')
        return unzip



    def extract_stats(self, sha) -> dto.DatasetDescriptor:

        file = self._folder / sha
        with file.open('rb') as f:
            with gzip.GzipFile(fileobj=f, mode='rb') as gz:
                df = pd.read_csv(gz, sep='\t')


        schema = []

        for name, type in df.dtypes.iteritems():
            schema.append(dto.FieldSchema(name=name, type=str(type)))

        return dto.DatasetDescriptor(type='table', record_count=len(df), schema=schema, sample_string=df.to_string(max_rows=10))



    def upload_new_version_from_file(self, dataset_id: Union['DatasetID',str], name):
        # compute new version

        if isinstance(dataset_id, DatasetID):
            dataset_id = dataset_id.id

        dataset = self._db.get_dataset(dataset_id)

        if not dataset:
            raise KeyError(f"Dataset {dataset_id} not found" )

        # print(f'Dataset: {dataset}')

        dataset.max_version += 1


        ver = dataset.max_version

        raw_size = 0
        zip_size = 0
        sha1 = hashlib.sha1()
        basename = os.path.basename(name)

        with open(name, 'rb') as source:
            with io.BytesIO() as target:
                with gzip.GzipFile(fileobj=target, mode='w') as zip:
                    while 1:
                        buf = source.read(64 * 1024)
                        if not buf:
                            break
                        raw_size += len(buf)
                        sha1.update(buf)
                        zip.write(buf)



                hex = sha1.hexdigest()
                store_file = self._folder / hex



                if store_file.exists():
                    print(f'Data from {basename} already exists, deduplicating it.')
                else:
                    target.seek(0)
                    with store_file.open('wb') as stor:
                        while 1:
                            buf = target.read(64*1024)
                            if not buf:
                                break
                            zip_size += len(buf)
                            stor.write(buf)

        chunk = dto.ChunkData(hash=hex, name=basename, raw_size=raw_size, zip_size=zip_size)
        source = dto.SourceLink(kind=dto.UPLOAD, source_name=basename, user_name='abdullin')

        version = dto.VersionData(version=ver, chunks=[chunk], sources=[source],
                                  raw_size=raw_size, zip_size=zip_size,dataset_id=dataset_id,
                                  descriptor=self.extract_stats(hex))

        self._db.add_dataset_version(version)
        self._db.add_dataset(dataset)


        return version

    def create_dataset(self, project: 'ProjectID', name: str, tags=None) -> 'DatasetID':
        "Dataset names are non-unique"

        prj = self._db.get_project(project.id)
        if not prj:
            raise ValueError(f"Project '{project.id}' not found")

        project_id = project.id
        dataset_id = 'ds-' + self._new_id()

        tags = tags or {}



        ds = dto.DatasetData(id=dataset_id,
                             project_id=project_id,
                             metadata=tags,
                             name=name
        )


        self._db.add_dataset(ds)
        return DatasetID(id=dataset_id, name=name)

    def list_versions(self, dataset_id: Union['DatasetID', str]):



        @dataclass
        class VerList:
            items:List[dto.VersionData]

            def _repr_html_(self):
                t = Table('#','Size', 'Archive', 'Sources')



                for i in self.items:
                    size = f'{i.chunks}&nbsp;/&nbsp;{sizeof_fmt(i.raw_size)}'
                    t.row(i.version, size, sizeof_fmt(i.zip_size), i.sources)
                return t._repr_html_()

        if isinstance(dataset_id,DatasetID):
            dataset_id = dataset_id.id

        versions = self._db.list_dataset_versions(dataset_id)

        return VerList(list(versions))


    def list_datasets(self, project: 'ProjectID' = None, **kwargs):

        @dataclass
        class DatasetItem:
            version: int
            name: str
            meta: Dict[str, str]
            cells: List[str]
            id: str
            master_size: int
            chunks:int
            stor_size: int

            @property
            def master_size_pretty(self):
                return f'{self.chunks} / {sizeof_fmt(self.master_size)}'
            @property
            def stor_size_pretty(self):
                return sizeof_fmt(self.stor_size)



        @dataclass
        class DatasetList:
            meta: List[str]
            items: List[DatasetItem]

            def _repr_html_(self):

                meta = self.meta

                cols = ['Dataset', 'ID', 'Ver', 'Size', 'Archive']
                cols.extend(meta)

                t = Table(*cols)

                t.monospace(1)

                for i in self.items:
                    master_size = i.master_size_pretty
                    cells = [ i.name, i.id, i.version, master_size, sizeof_fmt(i.stor_size)]
                    cells.extend(i.cells)
                    t.row(*cells)

                return t._repr_html_()

        r = []

        if project:
            datasets = self._db.get_datasets_by_project_id(project.id)
        else:
            datasets = self._db.get_datasets()


        if kwargs:
            print(f"Filtering by {kwargs}")
            search = set(kwargs.items())
            matches = []
            for ds in datasets:
                target = ds.metadata

                match = True

                for k,v in kwargs.items():
                    if k not in target:
                        match = False
                        break

                    if v == '*':
                        continue

                    if v == target[k]:
                        continue
                    match = False
                    break




                if match:
                    matches.append(ds)
            datasets = matches

        meta = set()

        for ds in datasets:
            for k in ds.metadata.keys():
                meta.add(k)

        meta = sorted(meta)

        for dataset in datasets:
            dataset_id = dataset.id
            versions = self._db.list_dataset_versions(dataset_id)

            master = list(versions)[-1]

            raw_size = master.raw_size

            chunks = len(master.chunks)

            cells = []
            for m in meta:
                cells.append(dataset.metadata.get(m, ''))



            sizes = {}
            for ver in versions:
                for c in ver.chunks:
                    sha = c.hash
                    curr = sizes.get(sha, 0)
                    curr += c.zip_size
                    sizes[sha] = curr
            stor_size=sum(sizes.values())

            r.append(DatasetItem(
                version=dataset.max_version,
                name=dataset.name,
                id=dataset_id,
                cells=cells,
                meta=dict(dataset.metadata),
                master_size=raw_size,
                chunks=chunks,
                stor_size=stor_size
            ))

        return DatasetList(meta, r)

    def describe_dataset(self, dataset_id: Union['DatasetID', str], version_num=-1) -> 'DatasetInfo':

        if isinstance(dataset_id, DatasetID):
            dataset = dataset_id.id
        dataset = self._db.get_dataset(dataset_id)
        versions = self._db.list_dataset_versions(dataset_id)

        version = versions[version_num]

        project = self._db.get_project(dataset.project_id)

        print(dataset)
        print(version)

        stats = version.descriptor
        records = stats.record_count


        max_version = len(versions)
        current_version = version_num
        if current_version == -1:
            current_version = max_version



        sample = stats.sample_string
        sources = version.sources

        return DatasetInfo(
            name=dataset.name,
            meta=dict(dataset.metadata),
            sample=sample,
            records=records,
            schema=stats.schema,
            max_version=max_version,
            current_version=current_version,
            version = version,
            project_id=dataset.project_id,
            project_name=project.name

        )

    def list_projects(self) -> 'ProjectList':
        items = []
        for v in self._db.list_projects():

            name = v.name
            project_id = v.id

            datasets = self._db.get_datasets_by_project_id(project_id)

            size = 0

            for ds in datasets:


                versions = self._db.list_dataset_versions(ds.id)
                for ver in versions:
                    size += ver.zip_size



            items.append(ProjectListItem(name,
                                         size,
                                         len(datasets)),
                         )
        return ProjectList(items)


@dataclass
class DatasetInfo:
    name: str
    meta: Dict[str,str]
    sample: str
    records: int
    schema: List[Dict[str,str]]
    current_version: int
    max_version: int
    version: dto.VersionData
    project_name: str
    project_id: str


@dataclass
class VersionID:
    id: str
    num: int


@dataclass
class ProjectID:
    id: str
    name: str


@dataclass
class DatasetID:
    id: str
    name: str

    def _repr_pretty_(self, p, cycle):
        p.text(f'Dataset({self.name}/{self.id})')


@dataclass
class DatasetVersionID:
    dataset_id: str
    version: int


@dataclass
class ProjectView:
    name: str
    id: str


@dataclass
class ProjectListItem:
    name: str
    size: int

    dataset_count: int

    @property
    def size_pretty(self):
        return sizeof_fmt(self.size)


@dataclass
class ProjectList:
    """
    A list of projects returned by the client
    """
    projects: List[ProjectListItem]

    def _repr_html_(self):
        t = Table('Project', 'Datasets', 'Notebooks', 'Models', 'Size')

        for p in self.projects:
            t.row(p.name, p.dataset_count, 0, 0, sizeof_fmt(p.size))

        return t._repr_html_()


@dataclass
class ProjectListDetail:
    name: str
