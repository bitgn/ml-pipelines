.PHONY: clean, dist


clean:
	rm -rf dist || True
	rm -rf build || True
	rm -rf mlp_client.egg-info || True


dist: clean
	@VERSION=$(VERSION) python3 setup.py sdist bdist_wheel