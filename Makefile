GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=adm
LASTVERSION=v1.0.2
VERSION=v1.0.3
CLI=adm

TEST_CONFIG_PATH=./../common/config.json
TEST_CONFIG_PQ_PATH=./../common/config_pg.json
TEST_CONFIG_SQLITE_PATH=./../common/config_sqlite.json

all: run

tmpl:
	$(CLI) compile tpl

fmt:
	go fmt ./adapter/...
	go fmt ./adm/...
	go fmt ./context/...
	go fmt ./engine/...
	go fmt ./tests/...
	go fmt ./examples/...
	go fmt ./modules/...
	go fmt ./plugins/...
	go fmt ./template/...

golint:
	golint ./adapter/...
	golint ./adm/...
	golint ./context/...
	golint ./engine/...
	golint ./tests/...
	golint ./examples/...
	golint ./modules/...
	golint ./plugins/...
	golint ./template/...

govet:
	go vet ./adapter/...
	go vet ./adm/...
	go vet ./context/...
	go vet ./engine/...
	go vet ./tests/...
	go vet ./examples/...
	go vet ./modules/...
	go vet ./plugins/...
	go vet ./template/...

deps:
	go get github.com/kardianos/govendor
	govendor sync

test:
	make mysql-test
	make pg-test
	make sqlite-test

mysql-test:
	make import-mysql
	gotest -v ./tests/gin/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/beego/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/buffalo/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/chi/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/echo/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/gorilla/... -args $(TEST_CONFIG_PATH)

sqlite-test:
	make import-sqlite
	gotest -v ./tests/gin/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/beego/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/buffalo/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/chi/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/echo/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/gorilla/... -args $(TEST_CONFIG_SQLITE_PATH)


import-sqlite:
	rm -rf ./tests/common/admin.db
	cp ./data/admin.db ./tests/common/admin.db

import-mysql:
	mysql -uroot -proot go-admin-test < ./data/admin.sql

import-postgresql:
	dropdb -U postgres go-admin-test
	createdb -U postgres go-admin-test
	psql -d go-admin-test -U postgres -f ./data/admin.pgsql

pg-test:
	make import-postgresql
	gotest -v ./tests/gin/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/beego/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/buffalo/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/chi/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/echo/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/gorilla/... -args $(TEST_CONFIG_PQ_PATH)

lint:
	make golint
	make govet
	golangci-lint run

cli:
	GO111MODULE=on $(GOBUILD) -ldflags "-w" -o ./adm/build/mac/$(BINARY_NAME) ./adm/...
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./adm/build/linux/x86_64/$(BINARY_NAME) ./adm/...
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o ./adm/build/linux/armel/$(BINARY_NAME) ./adm/...
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./adm/build/windows/x86_64/$(BINARY_NAME).exe ./adm/...
	GO111MODULE=on CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o ./adm/build/windows/i386/$(BINARY_NAME).exe ./adm/...
	rm -rf ./adm/build/linux/armel/adm_linux_armel_$(LASTVERSION).zip
	rm -rf ./adm/build/linux/x86_64/adm_linux_x86_64_$(LASTVERSION).zip
	rm -rf ./adm/build/windows/x86_64/adm_windows_x86_64_$(LASTVERSION).zip
	rm -rf ./adm/build/windows/i386/adm_windows_i386_$(LASTVERSION).zip
	rm -rf ./adm/build/mac/adm_darwin_x86_64_$(LASTVERSION).zip
	zip -qj ./adm/build/linux/armel/adm_linux_armel_$(VERSION).zip ./adm/build/linux/armel/adm
	zip -qj ./adm/build/linux/x86_64/adm_linux_x86_64_$(VERSION).zip ./adm/build/linux/x86_64/adm
	zip -qj ./adm/build/windows/x86_64/adm_windows_x86_64_$(VERSION).zip ./adm/build/windows/x86_64/adm.exe
	zip -qj ./adm/build/windows/i386/adm_windows_i386_$(VERSION).zip ./adm/build/windows/i386/adm.exe
	zip -qj ./adm/build/mac/adm_darwin_x86_64_$(VERSION).zip ./adm/build/mac/adm