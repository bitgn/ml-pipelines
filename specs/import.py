import argparse
import json
import os

import grpc

from test_api import events_pb2 as evt


from client import ml_pipelines as client





def gen_import_events(file_name, cl:client.Client):



    folder = os.path.dirname(file_name)
    print(file_name)
    with open(file_name, 'r') as f:
        meta = json.load(f)



        for i in meta:
            if i['type'] == 'expert':
                pass
                # yield evt.ExpertAdded(expert_id=i['expert_id'], expert_name=i['expert_name'])

            if i['type'] == 'project':
                cl.create_project(i['project_id'], title=i['project_name'])


            if i['type'] == 'dataset':

                project =i['project_id']
                title =i['dataset_name']
                name = i['dataset_id']


                prj = cl.get_project(project)


                location_uri = None
                location_id = None
                description = None



                if 'location_id' in i:
                    location_id = i['location_id']

                if 'location_uri' in i:
                    location_uri = i['location_uri']


                description_path = os.path.join(folder, name + ".md")
                if os.path.exists(description_path):
                    print(f"Import {description_path}")
                    with open(description_path, 'r') as f:
                        description = "\n".join(f.readlines())

                ds = prj.create_dataset(
                    name,
                    title=title,
                    data_format=i['data_format'],
                    location_uri=location_uri,
                    location_id=location_id,
                    description=description,

                )

                stage = ds.get_last_version().prepare_commit(clean_slate=True)
                stage.add_input()

                stage.commit(timestamp=i['update_timestamp'])





                mtd = evt.DatasetMetadata(
                    file_count=i['file_count'],
                    file_count_set=True,
                    record_count=i['record_count'],
                    record_count_set=True,
                    data_format=,
                    data_format_set=True,
                    storage_bytes=i['zip_bytes'],
                    storage_bytes_set=True,
                    update_timestamp=,
                    update_timestamp_set=True,
                    sample=sample,
                    sample_set=True,
                )








                e = evt.DatasetCreated(project_id=project,
                                       name=title,
                                       meta=mtd,
                                       dataset_id=name)

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
from test_api import api_pb2_grpc as rpc
from test_api import marshal
from test_api import api_pb2 as api

parser = argparse.ArgumentParser(description='Process some integers.')

parser.add_argument("--json", action="append", dest="json")
parser.add_argument("--grpc", action="store", dest="grpc", default="localhost:9111")

args = parser.parse_args()


channel = grpc.insecure_channel(args.grpc)
print("Connecting")

stub = rpc.TestStub(channel)
stub.Ping(api.PingRequest())

catalog = rpc.CatalogStub(channel)
print("Reset")
stub.Wipe(api.WipeDatabase())

print("Apply")
if args.json:
    for f in args.json:
        events = list(gen_import_events(f))
        catalog.Apply(api.ApplyRequest(Events=marshal.serialize(list(events))))
        print(f"{f}: {len(events)}")
print("Done")