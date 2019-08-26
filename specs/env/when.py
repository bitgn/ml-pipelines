from typing import Optional

import grpc
import requests as r
import env
import urllib.parse as url

from client import ml_pipelines as client

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


def client_create_project(name: str, title: Optional[str] = None):
    def _(c: client.Client):
        try:
            return c.create_project(name=name)
        except Exception as e:
            return e

    return env.When(web_action=None, text="create project", client_action=_)


def _get_page(text, uri):
    def _(c: r.Session, base: str):
        full = url.urljoin(base, uri)
        return c.get(full)

    return env.When(web_action=_, text=text, client_action=None)
