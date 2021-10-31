VERSION := $(if $(RELEASE_VERSION),$(RELEASE_VERSION),"master")

all: pre_clean darwin linux windows

pre_clean:
	rm -rf dist

darwin:
	GOOS=darwin GOARCH=amd64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_darwin_amd64.zip waterkube
	rm -f dist/waterkube

linux:
	GOOS=linux GOARCH=amd64 go build -o dist/waterkube .
	cd dist && zip waterkube_$(VERSION)_linux_amd64.zip waterkube
	rm -f dist/waterkube

windows:
	GOOS=windows GOARCH=amd64 go build -o dist/waterkube.exe .
	cd dist && zip waterkube_$(VERSION)_windows_amd64.zip waterkube.exe
	rm -f dist/waterkube.exe
