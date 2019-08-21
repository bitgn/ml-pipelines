.PHONY: pb, client-pb, test-pb, dist, clean

VERSION ?= dev


client-pb:
	# client needs the public api that doesn't have external dependencies
	venv/bin/python -m grpc_tools.protoc -I ../proto --python_out=client/ml_pipelines --grpc_python_out=client/ml_pipelines mlp_api.proto --mypy_out=client/ml_pipelines --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
	sed -i "" "s/^import api_pb2/from . import api_pb2/g" client/ml_pipelines/api_pb2_grpc.py

test-pb:
	venv/bin/python -m grpc_tools.protoc -I ../proto --python_out=test_api --grpc_python_out=test_api test_api.proto mlp_api.proto events.proto --mypy_out=test_api --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
	sed -i "" "s/^import events_pb2/from . import events_pb2/g" test_api/test_api_pb2.py
	sed -i "" "s/^import mlp_api_pb2/from . import mlp_api_pb2/g" test_api/mlp_api_pb2_grpc.py
	sed -i "" "s/^import test_api_pb2/from . import test_api_pb2/g" test_api/test_api_pb2_grpc.py

pb: client-pb test-pb

dist:
	$(MAKE) VERSION=$(VERSION) -C client dist
