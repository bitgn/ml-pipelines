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

ready-to-test:
	echo $(CURRENT_VER)
	git tag -a -m "Release to 'Ready to test'"  $(PROJECT)-$(NEXT_VER)
	git push --tags
	lib/deploy-qa $(PROJECT)-$(NEXT_VER)

protobuf:
	cd mlp && protoc --python_out=. --mypy_out=.  proto/dto.proto proto/events.proto  --plugin=protoc-gen-mypy=venv/bin/protoc-gen-mypy
