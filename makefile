up: postgres && prometheus

dependency-build:
	@cd chart && helm dependency build
app:
	@cd chart && helm install app . -f values.yaml


remove:
	@helm uninstall app

