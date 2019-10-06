GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=admincli
LASTVERSION=v0.3.0
VERSION=v0.3.1

all:
	GO111MODULE=on $(GOBUILD) -o ./build/mac/$(BINARY_NAME)
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/linux/x86_64/$(BINARY_NAME)
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o ./build/linux/armel/$(BINARY_NAME)
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./build/windows/x86_64/$(BINARY_NAME).exe
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o ./build/windows/i386/$(BINARY_NAME).exe
	rm -rf build/linux/armel/admincli_linux_armel_$(LASTVERSION).zip
	rm -rf build/linux/x86_64/admincli_linux_x86_64_$(LASTVERSION).zip
	rm -rf build/windows/x86_64/admincli_windows_x86_64_$(LASTVERSION).zip
	rm -rf build/windows/i386/admincli_windows_i386_$(LASTVERSION).zip
	rm -rf build/mac/admincli_darwin_x86_64_$(LASTVERSION).zip
	zip -q build/linux/armel/admincli_linux_armel_$(VERSION).zip build/linux/armel/admincli
	zip -q build/linux/x86_64/admincli_linux_x86_64_$(VERSION).zip build/linux/x86_64/admincli
	zip -q build/windows/x86_64/admincli_windows_x86_64_$(VERSION).zip build/windows/x86_64/admincli.exe
	zip -q build/windows/i386/admincli_windows_i386_$(VERSION).zip build/windows/i386/admincli.exe
	zip -q build/mac/admincli_darwin_x86_64_$(VERSION).zip build/mac/admincli