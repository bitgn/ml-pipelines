

def view_dataset(ds_id):
    return f"/datasets/{ds_id}/"


def search_datasets(query):
    return f'/explore?query={query}'


def view_project(project_id):
    return f"/projects/{project_id}/"


def list_datasets():
    return "/explore"


def list_projects():
    return "/"