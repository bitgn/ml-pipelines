

def view_dataset(project_name, dataset_name):
    return f"/projects/{project_name}/datasets/{dataset_name}"


def search_datasets(query):
    return f'/explore?query={query}'


def view_project(project_name):
    return f"/projects/{project_name}"


def list_datasets():
    return "/explore"


def list_projects():
    return "/"