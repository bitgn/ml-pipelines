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


release:
	echo "$(CURRENT_VER) - $(NEXT_VER)"


	lib/release-server-docker $(NEXT_VER)
	lib/release-client $(NEXT_VER)

	git tag -a -m "Release $(NEXT_VER)" $(PROJECT)-$(NEXT_VER)
	git push --tags
