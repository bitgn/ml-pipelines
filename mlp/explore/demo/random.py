"""
Generates events to setup a random demo

"""

import random


from explore.specs import preset
from proto import events_pb2 as evt
import test

def setup_random_demo():
    t = test.Env()

    datasets = set()
    for i in range(random.randint(5, 10)):
        project = preset.project_created(t)
        yield project


        for d in range(random.randint(2, 7)):
            ds = preset.dataset_created(t, project)
            datasets.add(ds.dataset_id)

            yield ds

    # split datasets into pipelines

    while datasets:
        pipeline_size = min(len(datasets), random.randint(1, 10))
        pipeline = set(random.choices(list(datasets), k=pipeline_size))

        datasets.difference_update(pipeline)

        prev_step = None
        while pipeline:
            step = set(random.choices(list(pipeline), k=random.randint(1, 3)))

            if prev_step:
                link_id = f"random-{t.next_id()}"
                link_name = link_id

                if len(prev_step) > 1:
                    link_name = random.choice(['merge', 'join', 'enrich'])
                else:
                    link_name = random.choice(['clean', 'filter', 'upload', 'format'])

                e = evt.DatasetLinkAdded(link_id=link_id, link_name=link_name, inputs=prev_step, outputs=step, )
                yield e

            prev_step = step

            pipeline.difference_update(step)
