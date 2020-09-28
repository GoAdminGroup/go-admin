GOCMD = go
GOBUILD = $(GOCMD) build
BINARY_NAME = adm
LAST_VERSION = v1.2.16
VERSION = v1.2.17
CLI = adm

TEST_CONFIG_PATH=./../../common/config.json
TEST_CONFIG_PQ_PATH=./../../common/config_pg.json
TEST_CONFIG_SQLITE_PATH=./../../common/config_sqlite.json
TEST_CONFIG_MS_PATH=./../../common/config_ms.json
TEST_FRAMEWORK_DIR=./tests/frameworks

all: test

## tests

test: cp-mod black-box-test web-test restore-mod

## tests: black box tests

black-box-test: mysql-test pg-test sqlite-test ms-test

mysql-test: $(TEST_FRAMEWORK_DIR)/*
	go get github.com/ugorji/go/codec@none
	for file in $^ ; do \
	make import-mysql ; \
	gotest -v ./$${file}/... -args $(TEST_CONFIG_PATH) ; \
	done

sqlite-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-sqlite ; \
	gotest -v ./$${file}/... -args $(TEST_CONFIG_SQLITE_PATH) ; \
	done

pg-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-postgresql ; \
	gotest -v ./$${file}/... -args $(TEST_CONFIG_PQ_PATH) ; \
	done

ms-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-mssql ; \
	gotest -v ./$${file}/... -args $(TEST_CONFIG_MS_PATH) ; \
	done

## tests: user acceptance tests

web-test: import-mysql
	gotest -v ./tests/web/...
	rm -rf ./tests/web/User*

web-test-debug: import-mysql
	gotest -v ./tests/web/... -args true

## tests: unit tests

unit-test:
	gotest -v ./adm/...
	gotest -v ./context/...
	gotest -v ./modules/...
	gotest -v ./plugins/admin/controller/...
	gotest -v ./plugins/admin/modules/parameter/...
	gotest -v ./plugins/admin/modules/table/...
	gotest -v ./plugins/admin/modules/...

## tests: helpers

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

import-mssql:
	docker exec mssql /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P Aa123456 -Q "RESTORE DATABASE [goadmin] FROM DISK = N'/home/data/admin_ms.bak' WITH FILE = 1, NOUNLOAD, REPLACE, RECOVERY, STATS = 5"

backup-mssql:
	docker exec mssql /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P Aa123456 -Q "BACKUP DATABASE [goadmin] TO DISK = N'/home/data/admin_ms.bak' WITH NOFORMAT, NOINIT, NAME = 'goadmin-full', SKIP, NOREWIND, NOUNLOAD, STATS = 10"

fix-gf:
	go get -u -v github.com/gogf/gf@v1.12.1
	sudo chmod -R 777 $(GOPATH)/pkg/mod/github.com/gogf/gf@v1.12.1/net/ghttp/ghttp_server_handler.go
	sudo echo "\nfunc (s *Server) DefaultHttpHandle(w http.ResponseWriter, r *http.Request) { \n s.handleRequest(w, r) \n}\n" >> $(GOPATH)/pkg/mod/github.com/gogf/gf@v1.12.1/net/ghttp/ghttp_server_handler.go

cp-mod:
	cp go.mod go.mod.old
	cp go.sum go.sum.old

restore-mod:
	mv go.mod.old go.mod
	mv go.sum.old go.sum

## code style check

lint: fmt golint govet cilint

fmt:
	GO111MODULE=off go fmt ./...
	GO111MODULE=off goimports -l -w .

govet:
	GO111MODULE=off go vet ./...

cilint:
	GO111MODULE=off golangci-lint run

golint:
	GO111MODULE=off golint ./...

build-tmpl:
    ## form tmpl build
	adm compile tpl --src ./template/types/tmpls/ --dist ./template/types/tmpl.go --package types --var tmpls
    ## generator tmpl build
	adm compile tpl --src ./plugins/admin/modules/table/tmpl --dist ./plugins/admin/modules/table/tmpl.go --package table --var tmpls

## cli version update

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

.PHONY: all fmt golint govet cp-mod restore-mod test black-box-test mysql-test sqlite-test import-sqlite import-mysql import-postgresql pg-test fix-gf lint cilint cli