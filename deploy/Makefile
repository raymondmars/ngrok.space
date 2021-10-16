
re-generate-cert:
	bash scripts/recreate-cert.sh

clean:
	rm -rf ../client/assets/ ../server/assets/

client-assets: BUILDTAGS=release
client-assets: 
	bin/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=../client/assets/assets_$(BUILDTAGS).go \
		assets/client/...

server-assets: BUILDTAGS=release
server-assets: 
	bin/go-bindata -nomemcopy -pkg=assets -tags=$(BUILDTAGS) \
		-debug=$(if $(findstring debug,$(BUILDTAGS)),true,false) \
		-o=../server/assets/assets_$(BUILDTAGS).go \
		assets/server/...

assets: client-assets server-assets

deploy:
	scp bin/ngrokd root@cat:~/up/ngrok/
