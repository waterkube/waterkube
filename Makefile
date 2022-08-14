VERSION := $(if $(RELEASE_VERSION),$(RELEASE_VERSION),"master")

all: pre_clean ui darwin darwin_arm64 linux linux_arm64 windows post_clean
docker: unzip_linux unzip_linux_arm64

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

darwin_arm64:
	GOOS=darwin GOARCH=arm64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_darwin_arm64.zip .env waterkube
	rm -f dist/waterkube

linux:
	GOOS=linux GOARCH=amd64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_linux_amd64.zip .env waterkube
	rm -f dist/waterkube

unzip_linux:
	cd dist && unzip waterkube_$(VERSION)_linux_amd64.zip -d waterkube_$(VERSION)_linux_amd64

linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_linux_arm64.zip .env waterkube
	rm -f dist/waterkube

unzip_linux_arm64:
	cd dist && unzip waterkube_$(VERSION)_linux_arm64.zip -d waterkube_$(VERSION)_linux_arm64

windows:
	GOOS=windows GOARCH=amd64 go build -o dist/waterkube.exe .
	cd dist && zip waterkube_$(VERSION)_windows_amd64.zip .env waterkube.exe
	rm -f dist/waterkube.exe

post_clean:
	rm -rf dist/.env
