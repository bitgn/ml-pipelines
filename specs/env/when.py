from typing import Optional, Callable, Any

import grpc
import requests as r
import env
import urllib.parse as url
import inspect
from client import ml_pipelines as cl

import test_api as api


def view_dataset(project, dataset):
    return _get_page("view_dataset", env.urls.view_dataset(project, dataset))


def search_datasets(query):
    return _get_page(f"search datasets for {query}", env.urls.search_datasets(query))


def view_project(project_id):
    return _get_page("view project", env.urls.view_project(project_id))


def list_datasets():
    return _get_page("list datasets", env.urls.list_datasets())


def list_projects():
    return _get_page("list projects", env.urls.list_projects())

def client(l: Callable[[cl.Client], Any], text:Optional[str]=None):
    def _(c: cl.Client):
        try:
            return l(c)
        except Exception as e:
            return e

    if not text:
        text = '('+inspect.getsource(l).replace("when.client(lambda c:", "").strip('\n ,')

    return env.When(web_action=None, client_action=_, text=text)


def _get_page(text, uri):
    def _(c: r.Session, base: str):
        full = url.urljoin(base, uri)
        return c.get(full)

    return env.When(web_action=_, text=text, client_action=None)
