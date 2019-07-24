from django.test import client
from test import env


def view_dataset(ds_id):
    return _get_page("view_dataset", f"/datasets/{ds_id}/")


def search_datasets(query):
    return _get_page(f"search datasets for {query}", f'/datasets/?query={query}')


def view_project(project_id):
    return _get_page("view project", f"/projects/{project_id}/")


def list_datasets():
    return _get_page("list datasets", "/datasets/")


def list_projects():
    return _get_page("list projects", "")


def _get_page(text, viewname):
    def _(c: client.Client):
        return c.get(viewname)

    return env.When(action=_, text=text)
