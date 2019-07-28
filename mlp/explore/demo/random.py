"""
Generates events to setup a random demo

"""

import random

import faker
from explore.specs import preset
from proto import events_pb2 as evt
import test

def setup_random_demo():
    t = test.Env()

    f = faker.Faker()

    datasets = set()

    mapping = {}
    experts = []

    for i in range(random.randint(5, 10)):
        name = f.name()

        expert_id = f'expert-{i}'
        expert_name = f.name()

        experts.append(expert_id)

        yield evt.ExpertAdded(expert_id=expert_id, expert_name=expert_name)


    for i in range(random.randint(5, 10)):
        project = preset.project_created(t)
        yield project


        for d in range(random.randint(2, 7)):
            ds = preset.dataset_created(t, project)
            datasets.add(ds.dataset_id)

            mapping[ds.dataset_id] = project.project_id

            ds.experts.extend(random.choices(experts, k=random.randint(0,2)))

            yield ds

    # split datasets into pipelines

    while datasets:
        pipeline_size = min(len(datasets), random.randint(1, 10))
        pipeline = set(random.choices(list(datasets), k=pipeline_size))

        datasets.difference_update(pipeline)

        prev_step = None
        while pipeline:
            step = set(random.choices(list(pipeline), k=random.randint(0, 3)))

            if prev_step:
                job_id = f"random-{t.next_id()}"
                job_name = job_id

                if len(prev_step) > 1:
                    job_name = random.choice(['merge', 'join', 'enrich'])
                else:
                    job_name = random.choice(['clean', 'filter', 'upload', 'format'])

                if step:
                    project_id = mapping[list(step)[0]]
                else:
                    project_id = mapping[list(prev_step)[0]]
                e = evt.JobAdded(job_id=job_id, job_name=job_name, inputs=prev_step, outputs=step, project_id=project_id)
                e.experts.extend(random.choices(experts, k=random.randint(0, 2)))
                yield e

            prev_step = step

            pipeline.difference_update(step)
