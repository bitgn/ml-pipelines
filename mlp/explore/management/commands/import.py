"""
Import interop JSON format into the database
"""
import lmdb
from django.conf import settings
from django.core.management import BaseCommand
from agg import projection, db
from explore import demo
from proto import events_pb2 as evt
import random
import json

random.seed(1)


class Command(BaseCommand):
    help = 'import a demo format'

    def add_arguments(self, parser):
        parser.add_argument("--json")
        pass

    def handle(self, *args, **options):
        with lmdb.open(settings.LMDB, subdir=False) as env:
            main = env.open_db()
            file_name = options['json']
            print(file_name)
            with open(file_name, 'r') as f:
                meta = json.load(f)


            with env.begin(write=True) as tx:


                for i in meta:
                    if i['type'] == 'expert':
                        e = evt.ExpertAdded(expert_id=i['expert_id'], expert_name=i['expert_name'])
                        projection.apply(e, tx)
                        continue
                    if i['type'] == 'project':
                        self.create_project(i, tx)
                    if i['type'] == 'dataset':

                        project_id=i['project_id']
                        name=i['dataset_name']
                        dataset_id = i['dataset_id']

                        mtd = evt.DatasetMetadata(
                            file_count=i['file_count'],
                            record_count=i['record_count'],
                            data_format=i['data_format'],
                            zip_bytes=i['zip_bytes'],
                            update_timestamp=i['update_timestamp'],
                            sample_body=i['sample'].encode(),
                            sample_kind=evt.SAMPLE_KIND.JSON
                        )


                        mtd.set_fields.extend([
                            evt.FIELD_UPDATE_TIMESTAMP,
                            evt.FIELD_RECORD_COUNT,
                            evt.FIELD_DATA_FORMAT,
                            evt.FIELD_ZIP_BYTES,
                            evt.FIELD_FILE_COUNT,
                            evt.FIELD_SAMPLE
                        ])

                        if 'storage_id' in i:
                            mtd.storage_id = i['storage_id']
                            mtd.set_fields.append(evt.FIELD_STORAGE_ID)

                        if 'storage_location' in i:
                            mtd.storage_location = i['storage_location']
                            mtd.set_fields.append(evt.FIELD_STORAGE_LOCATION)



                        e = evt.DatasetCreated(project_id=project_id,
                                               name=name,
                                               metadata=mtd,
                                               dataset_id=dataset_id)

                        if 'experts' in i:
                            e.experts.extend(i['experts'])

                        projection.apply(e, tx)

                    if i['type']=='job':
                        e = evt.JobAdded(
                            project_id=i['project_id'],
                            job_name=i['job_name'],
                            job_id=i['job_id'],
                            inputs=i['inputs'],
                            outputs=i['outputs']
                        )
                        projection.apply(e, tx)

                        #db.dataset_get()










            print(len(meta))

    def create_project(self, i, tx):
        prj = db.project_get(tx, project_id=i['project_id'])
        if not prj:
            e = evt.ProjectCreated(project_id=i['project_id'], name=i['project_name'])
            projection.apply(e, tx)

