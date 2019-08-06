.PHONY: pb


pb:
	venv/bin/python -m grpc_tools.protoc -I ../proto --python_out=api --grpc_python_out=api api.proto events.proto --mypy_out=api --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
	sed -i "" "s/^import events_pb2/from . import events_pb2/g" api/api_pb2.py
	sed -i "" "s/^import api_pb2/from . import api_pb2/g" api/api_pb2_grpc.py