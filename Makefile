VERSION := $(if $(RELEASE_VERSION),$(RELEASE_VERSION),"master")

all: pre_clean ui darwin linux windows post_clean

pre_clean:
	rm -rf dist
	mkdir dist
	cp .env.example dist/.env

ui:
	yarn prod

darwin:
	GOOS=darwin GOARCH=amd64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_darwin_amd64.zip .env waterkube
	rm -f dist/waterkube

linux:
	GOOS=linux GOARCH=amd64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_linux_amd64.zip .env waterkube
	rm -f dist/waterkube

windows:
	GOOS=windows GOARCH=amd64 go build -o dist/waterkube.exe .
	cd dist && zip waterkube_$(VERSION)_windows_amd64.zip .env waterkube.exe
	rm -f dist/waterkube.exe

post_clean:
	rm -rf dist/.env
