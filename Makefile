.PHONY: re-generate-cert clean-assets client-assets server-assets assets build-client build-server build-docker deploy remove-ds-store
export SHELL := env PWD=$(CURDIR) bash
ifeq ($(SHELL),)
$(error bash is required)
endif

BUILDTAGS=debug
OS=linux
CLIENT_MAIN_LOCATION=$(PWD)/cmd/client
SERVER_MAIN_LOCATION=$(PWD)/cmd/server 

CLIENT_BUILD_OUTPUT=$(PWD)/build/package/release/client
SERVER_BUILD_OUTPUT=$(PWD)/build/package/release/server 

BUILD_DEPLOY_PATH=$(PWD)/build/package/deploy

DOCKER_IMAGE='registry.raymondjiang.net/raymond/ngrok-space:latest'


re-generate-cert:
	bash scripts/recreate-cert.sh


clean-assets:
	rm -rf internal/app/client/assets/ internal/app/server/assets/

client-assets: BUILDTAGS=release
client-assets: 
	script/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=./internal/app/client/assets/assets_$(BUILDTAGS).go \
		assets/client/...

server-assets: BUILDTAGS=release
server-assets: 
	script/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=./internal/app/server/assets/assets_$(BUILDTAGS).go \
		assets/server/...

assets: client-assets server-assets


build-client:
	bash build/package/client/build.sh $(OS) $(CLIENT_MAIN_LOCATION) $(CLIENT_BUILD_OUTPUT)

build-client-linux: OS=linux
build-client-linux: build-client
	
build-client-macos: OS=macos
build-client-macos: build-client

build-client-windows: OS=windows
build-client-windows: build-client

build-server:
	bash build/package/server/build.sh $(OS) $(SERVER_MAIN_LOCATION) $(SERVER_BUILD_OUTPUT)


build-docker:
	bash scripts/docker-build.sh $(SERVER_BUILD_OUTPUT) $(BUILD_DEPLOY_PATH) $(DOCKER_IMAGE)

deploy: build-docker
	bash scripts/deploy.sh $(PWD)/deployments

remove-ds-store:
	find . -name ".DS_Store" -delete

