GOCMD = go
GOBUILD = $(GOCMD) build
BINARY_NAME = adm
LAST_VERSION = v1.2.4
VERSION = v1.2.5
CLI = adm

TEST_CONFIG_PATH=./../../common/config.json
TEST_CONFIG_PQ_PATH=./../../common/config_pg.json
TEST_CONFIG_SQLITE_PATH=./../../common/config_sqlite.json
WEB_TEST_CONFIG_PATH=./config.json

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
	go get github.com/ugorji/go/codec@none
	make mysql-test
	make pg-test
	make sqlite-test
	make web-test

mysql-test:
	make import-mysql
	gotest -v ./tests/frameworks/gin/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/beego/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/buffalo/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/chi/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/echo/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/gorilla/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/gf/... -args $(TEST_CONFIG_PATH)
	make import-mysql
	gotest -v ./tests/frameworks/fasthttp/... -args $(TEST_CONFIG_PATH)

sqlite-test:
	make import-sqlite
	gotest -v ./tests/frameworks/gin/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/beego/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/buffalo/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/chi/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/echo/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/gorilla/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/gf/... -args $(TEST_CONFIG_SQLITE_PATH)
	make import-sqlite
	gotest -v ./tests/frameworks/fasthttp/... -args $(TEST_CONFIG_SQLITE_PATH)

import-sqlite:
	rm -rf ./tests/common/admin.db
	cp ./tests/data/admin.db ./tests/common/admin.db

import-mysql:
	mysql -uroot -proot -e "create database if not exists \`go-admin-test\`"
	mysql -uroot -proot go-admin-test < ./tests/data/admin.sql

import-postgresql:
	dropdb -U postgres go-admin-test
	createdb -U postgres go-admin-test
	psql -d go-admin-test -U postgres -f ./tests/data/admin_pg.sql

pg-test:
	make import-postgresql
	gotest -v ./tests/frameworks/gin/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/beego/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/buffalo/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/chi/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/echo/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/gorilla/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/gf/... -args $(TEST_CONFIG_PQ_PATH)
	make import-postgresql
	gotest -v ./tests/frameworks/fasthttp/... -args $(TEST_CONFIG_PQ_PATH)

web-test:
	make import-mysql
	gotest -v ./tests/web/... -args $(WEB_TEST_CONFIG_PATH)
	rm -rf ./tests/web/User*

unit-test:
	gotest -v ./adm/...
	gotest -v ./context/...
	gotest -v ./modules/auth/...
	gotest -v ./modules/collection/...
	gotest -v ./modules/config/...
	gotest -v ./modules/db/...
	gotest -v ./modules/language/...
	gotest -v ./modules/logger/...
	gotest -v ./modules/menu/...
	gotest -v ./modules/utils/...
	gotest -v ./plugins/admin/controller/...
	gotest -v ./plugins/admin/modules/parameter/...
	gotest -v ./plugins/admin/modules/table/...
	gotest -v ./plugins/admin/modules/...

fix-gf:
	go get -u -v github.com/gogf/gf@v1.9.10
	sudo echo "\nfunc (s *Server) DefaultHttpHandle(w http.ResponseWriter, r *http.Request) { \n s.handleRequest(w, r) \n}\n" >> $(GOPATH)/pkg/mod/github.com/gogf/gf@v1.9.10/net/ghttp/ghttp_server_handler.go

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
	rm -rf ./adm/build/linux/armel/adm_linux_armel_$(LAST_VERSION).zip
	rm -rf ./adm/build/linux/x86_64/adm_linux_x86_64_$(LAST_VERSION).zip
	rm -rf ./adm/build/windows/x86_64/adm_windows_x86_64_$(LAST_VERSION).zip
	rm -rf ./adm/build/windows/i386/adm_windows_i386_$(LAST_VERSION).zip
	rm -rf ./adm/build/mac/adm_darwin_x86_64_$(LAST_VERSION).zip
	zip -qj ./adm/build/linux/armel/adm_linux_armel_$(VERSION).zip ./adm/build/linux/armel/adm
	zip -qj ./adm/build/linux/x86_64/adm_linux_x86_64_$(VERSION).zip ./adm/build/linux/x86_64/adm
	zip -qj ./adm/build/windows/x86_64/adm_windows_x86_64_$(VERSION).zip ./adm/build/windows/x86_64/adm.exe
	zip -qj ./adm/build/windows/i386/adm_windows_i386_$(VERSION).zip ./adm/build/windows/i386/adm.exe
	zip -qj ./adm/build/mac/adm_darwin_x86_64_$(VERSION).zip ./adm/build/mac/adm
	rm -rf ./adm/build/zip/*
	cp ./adm/build/linux/armel/adm_linux_armel_$(VERSION).zip ./adm/build/zip/
	cp ./adm/build/linux/x86_64/adm_linux_x86_64_$(VERSION).zip ./adm/build/zip/
	cp ./adm/build/windows/x86_64/adm_windows_x86_64_$(VERSION).zip ./adm/build/zip/
	cp ./adm/build/windows/i386/adm_windows_i386_$(VERSION).zip ./adm/build/zip/
	cp ./adm/build/mac/adm_darwin_x86_64_$(VERSION).zip ./adm/build/zip/

.PHONY: all tmpl fmt golint govet deps test mysql-test sqlite-test import-sqlite import-mysql import-postgresql pg-test fix-gf lint cli