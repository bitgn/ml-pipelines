import setuptools

import os

with open("README.md", "r") as fh:
    long_description = fh.read()

version = os.getenv("VERSION")

if not version:
    version = "dev"


setuptools.setup(
    name="ml_pipelines",
    version=version,
    author="Rinat Abdullin",
    author_email="rinat@abdullin.com",
    description="Python client for ML Pipelines",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/bitgn/ml-pipelines",
    packages=setuptools.find_packages(),
    install_requires=[
        'grpcio>=1.12.0',
        'protobuf'
    ],
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: BSD License",
        "Operating System :: OS Independent",
        "Development Status :: 3 - Alpha",
        "Topic :: Scientific/Engineering :: Information Analysis",
        "Topic :: Education"
    ],
)