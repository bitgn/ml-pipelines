from django import urls
from django.http import HttpResponse
from django.shortcuts import render

from explore.config import env
from explore.domain import stale
from agg import db

import markdown as md
import hashlib
import graphviz as g

import proto.dto_pb2 as dto


def view_dataset(request, dataset_id):
    with env.begin() as tx:
        ds = db.dataset_get(tx, dataset_id)
        if not ds:
            raise KeyError(f"Dataset {dataset_id} not found")
        prj = db.project_get(tx, ds.project_id)
        lineage = build_lineage(dataset_id, tx)
        svg = get_or_cache(tx, dataset_id, lineage).decode('utf-8').replace('<svg ', '<svg class="img-fluid" ')



    context = {
        'dataset': ds,
        'project': prj,
        'stale': stale.is_stale(ds),
        'lineage': svg
    }

    if ds.description_set:
        context['description'] = md.markdown(ds.description, extensions=['fenced_code', 'codehilite']).replace('h3>','h5>').replace('h2>','h4>').replace('h1>','h3>')

    return render(request, "explore/view-dataset.html", context)


def view_dataset_lineage(request, dataset_id):
    with env.begin() as tx:
        dot = build_lineage(dataset_id, tx)

    print(dot.source)

    # set the body
    r = HttpResponse(dot.pipe(format='svg'))

    r['Content-Type'] = 'image/svg+xml'

    return r


def build_lineage(dataset_id, tx):
    ds = db.dataset_get(tx, dataset_id)
    dot = g.Digraph(comment='lineage',
                    node_attr={'shape': 'rectangle', 'color': '#343a40',
                               'penwidth': "1.5", 'fontname': 'Arial'},
                    edge_attr={'penwidth': '1.0', 'color': '#343a40'})
    dot.attr(rankdir='LR', fontname='Arial')
    dot.node("this", label=ds.name, color="#28a745", bgcolor="#28a745")
    edge_node = {'style': 'rounded'}
    for l in ds.upstream_jobs:
        job = db.job_get(tx, l)

        if not job:
            print(f'no link for id {l}')

        dot.node(job.job_id, label=job.job_name, **edge_node)
        dot.edge(job.job_id, "this")

        for input_id in job.inputs:
            input_ds = db.dataset_get(tx, input_id)
            dot.node(input_id, label=input_ds.name, href=urls.reverse("explore:view_dataset", args=[input_id]))

            dot.edge(input_id, job.job_id, arrowhead='none')
    for l in ds.downstream_jobs:
        job = db.job_get(tx, l)
        dot.node(job.job_id, label=job.job_name, **edge_node)
        dot.edge("this", job.job_id, arrowhead='none')

        for output_id in job.outputs:
            output_ds = db.dataset_get(tx, output_id)
            dot.node(output_id, label=output_ds.name, href=urls.reverse("explore:view_dataset", args=[output_id]))
            dot.edge(job.job_id, output_id, arrowtail='none')
    return dot


def get_or_cache(tx, dataset_id, d: g.Digraph):
    m = hashlib.sha1()
    for l in d.body:
        m.update(l.encode())
    digest = m.hexdigest()

    cache = db.cache_get(tx, dataset_id)
    if cache and cache.digest == digest:
        return cache.body

    result = d.pipe(format='svg')

    with env.begin(write=True) as wr:
        cache = dto.AssetCache(digest=digest, body=result)
        db.cache_put(wr, dataset_id, cache)

    with env.begin() as r:
        print(db.cache_get(r, dataset_id) is not None)

    return result