.PHONY: pb, client-pb, test-pb, dist, clean




client-pb:
	# client needs the public api that doesn't have external dependencies
	venv/bin/python -m grpc_tools.protoc -I ../proto --python_out=mlp_client --grpc_python_out=mlp_client api.proto --mypy_out=mlp_client --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
	sed -i "" "s/^import api_pb2/from . import api_pb2/g" mlp_client/api_pb2_grpc.py

test-pb:
	venv/bin/python -m grpc_tools.protoc -I ../proto --python_out=test_api --grpc_python_out=test_api api.proto events.proto --mypy_out=test_api --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
	sed -i "" "s/^import events_pb2/from . import events_pb2/g" test_api/api_pb2.py
	sed -i "" "s/^import api_pb2/from . import api_pb2/g" test_api/api_pb2_grpc.py


pb: client-pb test-pb


clean:
	rm -rf build || True
	rm -rf mlp_client.egg-info || True

dist: clean
	VERSION=$(VERSION) python3 setup.py sdist