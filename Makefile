.PHONY: protobuf

PYLINT = mlp/venv/bin/pylint
GET_VERSION=./lib/get-version
INCREMENT_VERSION=./lib/inc-version

PROJECT=catalog
CURRENT_VER ?= $(shell $(GET_VERSION) $(PROJECT)-)
NEXT_VER ?= $(shell $(INCREMENT_VERSION) $(CURRENT_VER))



pylint:
	@PYLINT mlp --ignore=venv,notebooks,proto --rcfile=mlp/.pylintrc


specs:
	@cd mlp && source venv/bin/activate && ./manage.py specs

demo:
	@cd mlp && source venv/bin/activate && ./manage.py demo && ./manage.py runserver

dev:
	lib/deploy-dev

screenshots:
	mlp/venv/bin/python mlp/lib/screenshot.py doc

ready-to-test:
	echo $(CURRENT_VER)
	git tag -a -m "Release to 'Ready to test'"  $(PROJECT)-$(NEXT_VER)
	git push --tags
	lib/deploy-qa $(PROJECT)-$(NEXT_VER)

protobuf:
	protoc -I proto proto/events.proto --go_out=server/events
	protoc -I proto proto/db.proto --go_out=server/db
	protoc -I proto/ proto/api.proto --go_out=plugins=grpc:server/api

	cd mlp && protoc --python_out=. --mypy_out=.  proto/dto.proto proto/events.proto  --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
    protoc -I proto/ proto/catalog_service.proto --go_out=plugins=grpc:ml-go/service
