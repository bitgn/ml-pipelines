import argparse
import json
import os

import grpc

from api import events_pb2 as evt





def gen_import_events(file_name):



    folder = os.path.dirname(file_name)
    print(file_name)
    with open(file_name, 'r') as f:
        meta = json.load(f)



        for i in meta:
            if i['type'] == 'expert':
                yield evt.ExpertAdded(expert_id=i['expert_id'], expert_name=i['expert_name'])

            if i['type'] == 'project':
                yield evt.ProjectCreated(project_id=i['project_id'], name=i['project_name'])
            if i['type'] == 'dataset':

                project_id =i['project_id']
                name =i['dataset_name']
                dataset_id = i['dataset_id']

                sample = evt.DatasetSample(
                    body=i['sample'].encode(),
                    format=evt.DatasetSample.FORMAT.JSON
                )

                mtd = evt.DatasetMetadata(
                    file_count=i['file_count'],
                    file_count_set=True,
                    record_count=i['record_count'],
                    record_count_set=True,
                    data_format=i['data_format'],
                    data_format_set=True,
                    storage_bytes=i['zip_bytes'],
                    storage_bytes_set=True,
                    update_timestamp=i['update_timestamp'],
                    update_timestamp_set=True,
                    sample=sample,
                    sample_set=True,
                )



                if 'location_id' in i:
                    mtd.location_id = i['location_id']
                    mtd.location_id_set = True

                if 'location_uri' in i:
                    mtd.location_uri = i['location_uri']
                    mtd.location_uri_set = True

                description = os.path.join(folder, dataset_id + ".md")
                if os.path.exists(description):
                    print(f"Import {description}")
                    with open(description, 'r') as f:
                        mtd.description = "\n".join(f.readlines())
                        mtd.description_set = True





                e = evt.DatasetCreated(project_id=project_id,
                                       name=name,
                                       meta=mtd,
                                       dataset_id=dataset_id)

                if 'experts' in i:
                    e.meta.experts.extend(i['experts'])
                    e.meta.experts_set = True

                yield e

            if i['type' ]=='job':
                yield evt.JobAdded(
                    project_id=i['project_id'],
                    job_name=i['job_name'],
                    job_id=i['job_id'],
                    inputs=i['inputs'],
                    outputs=i['outputs']
                )
from api import api_pb2_grpc as rpc
from api import marshal
from api import api_pb2 as api

parser = argparse.ArgumentParser(description='Process some integers.')

parser.add_argument("--json", action="store", dest="json")
parser.add_argument("--grpc", action="store", dest="grpc", default="localhost:9111")

args = parser.parse_args()


channel = grpc.insecure_channel(args.grpc)
stub = rpc.TestStub(channel)
catalog = rpc.CatalogStub(channel)

stub.Wipe(api.WipeDatabase())


if args.json:
    events = gen_import_events(args.json)
    catalog.Apply(api.ApplyRequest(Events=marshal.serialize(list(events))))
