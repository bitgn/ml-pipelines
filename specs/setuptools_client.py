import setuptools

import os

with open("README.md", "r") as fh:
    long_description = fh.read()

version = os.getenv("VERSION", "0.0.1")


setuptools.setup(
    name="mlp-client",
    version=version,
    author="Rinat Abdullin",
    author_email="rinat@abdullin.com",
    description="Python client for ML Pipelines",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/bitgn/ml-pipelines",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: BSD 2-Clause",
        "Operating System :: OS Independent",
    ],
)