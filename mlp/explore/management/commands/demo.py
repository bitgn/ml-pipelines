import lmdb
from django.conf import settings

from django.core.management import BaseCommand

from agg import projection
from explore.specs import preset, social
import test


from proto import events_pb2 as evt
import random



random.seed(1)

class Command(BaseCommand):
    help = 'Fill a demo db'

    def add_arguments(self, parser):
        parser.add_argument("--setup", default="random")
        pass


    def handle(self, *args, **options):
        with lmdb.open(settings.LMDB, subdir=False) as env:
            main = env.open_db()

            with env.begin(write=True) as tx:
                tx.drop(main)


                if options['setup'] == "empty":
                    return

                self.random_setup(tx)


    def random_setup(self, tx):
        t = test.Env()

        datasets = set()
        for i in range(random.randint(5, 10)):
            project = preset.project_created(t)
            projection.apply(project, tx)

            for d in range(random.randint(2, 7)):
                ds = preset.dataset_created(t, project)
                datasets.add(ds.dataset_id)

                projection.apply(ds, tx)

        # split datasets into pipelines

        while datasets:
            pipeline_size = min(len(datasets), random.randint(1, 10))
            pipeline = set(random.choices(list(datasets), k=pipeline_size))

            datasets.difference_update(pipeline)

            prev_step = None
            while pipeline:
                step = set(random.choices(list(pipeline), k=random.randint(1,3)))


                if prev_step:
                    link_id = f"process-{t.next_id()}"
                    link_name = link_id

                    if len(prev_step)>1:
                        link_name=random.choice(['merge', 'join', 'enrich'])
                    else:
                        link_name=random.choice(['clean', 'filter', 'upload', 'format'])

                    e = evt.DatasetLinkAdded(link_id=link_id, link_name=link_name, inputs=prev_step, outputs=step, )
                    projection.apply(e, tx)


                    print(e)
                prev_step = step

                pipeline.difference_update(step)

