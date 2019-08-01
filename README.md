# ml-pipelines

Applications for managing machine learning pipelines and human workflows around them, published under BSD-2 license.

At the moment of writing this repository includes only the MLP Catalog - a Django application for exploring projects and datasets stored within the metadata library.


## Explore Datasets

Find relevant data by searching across all datasets.

<a href="./doc/explore-datasets.png">
    <img src="./doc/explore-datasets-thumb.png" width="320" alt="explore datasets">
</a>


<a href="./doc/explore-datasets-2.png">
    <img src="./doc/explore-datasets-2-thumb.png" width="320" alt="explore datasets">
</a>

## View Projects

Organize elements of ML Pipelines into projects.


<a href="./doc/view-projects.png">
    <img src="./doc/view-projects-thumb.png" width="320" alt="view projects">
</a>

<a href="./doc/view-project.png">
    <img src="./doc/view-project-thumb.png" width="320" alt="view project">
</a>


## View Datasets

View dataset properties and relations.

<a href="./doc/view-dataset.png">
    <img src="./doc/view-dataset-thumb.png" width="320" alt="view datasets">
</a>


<a href="./doc/view-dataset-2.png">
    <img src="./doc/view-dataset-2-thumb.png" width="320" alt="view datasets">
</a>

## Specs

Application functionality is being covered with [event-driven specs](https://abdullin.com/sku-vault/event-driven-verification/). This captures business logic and UX flows in non-fragile way.

## Getting started

Application is build and tested with Python 3.7.

Prerequisites:

- Python 3.7 with dev libraries: `apt install python3.7 python3.7-dev`
- [graphviz](https://www.graphviz.org): `apt install graphviz`
- [virtualenv](https://virtualenv.pypa.io/en/latest/)

To get started, go to the `mlp` folder and:

1) set up a _virtualenv_ in `mlp` folder and activate it;
3) `pip install -r requirements.txt` - install all the dependencies;
4) `python manage.py specs` - to run tests;
5) `python manage.py demo && python manage.py runserver` to fill up DB with demo data and launch the web UI (available at localhost:8000)



