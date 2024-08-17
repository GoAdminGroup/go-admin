GOCMD = go
GOBUILD = $(GOCMD) build

TEST_CONFIG_PATH=./../../common/config.json
TEST_CONFIG_PQ_PATH=./../../common/config_pg.json
TEST_CONFIG_SQLITE_PATH=./../../common/config_sqlite.json
TEST_CONFIG_MS_PATH=./../../common/config_ms.json
TEST_FRAMEWORK_DIR=./tests/frameworks

## database configs
MYSQL_HOST = db_mysql
MYSQL_PORT = 3306
MYSQL_USER = root
MYSQL_PWD = root

POSTGRESSQL_HOST = db_pgsql
POSTGRESSQL_PORT = 5432
POSTGRESSQL_USER = postgres
POSTGRESSQL_PWD = root

TEST_DB = go-admin-test

all: test

## tests

test: cp-mod black-box-test web-test restore-mod

## tests: black box tests

black-box-test: mysql-test pg-test sqlite-test ms-test

mysql-test: $(TEST_FRAMEWORK_DIR)/*
	go get github.com/ugorji/go/codec@none
	for file in $^ ; do \
	make import-mysql ; \
	go test -mod=mod -gcflags=all=-l -v ./$${file}/... -args $(TEST_CONFIG_PATH) ; \
	done

sqlite-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-sqlite ; \
	go test -mod=mod -gcflags=all=-l ./$${file}/... -args $(TEST_CONFIG_SQLITE_PATH) ; \
	done

pg-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-postgresql ; \
	go test -mod=mod -gcflags=all=-l ./$${file}/... -args $(TEST_CONFIG_PQ_PATH) ; \
	done

ms-test: $(TEST_FRAMEWORK_DIR)/*
	for file in $^ ; do \
	make import-mssql ; \
	go test -mod=mod -gcflags=all=-l ./$${file}/... -args $(TEST_CONFIG_MS_PATH) ; \
	done

## tests: user acceptance tests

web-test: import-mysql
	go test -mod=mod ./tests/web/...
	rm -rf ./tests/web/User*

web-test-debug: import-mysql
	go test -mod=mod ./tests/web/... -args true

## tests: unit tests

unit-test:
	go test -mod=mod ./adm/...
	go test -mod=mod ./context/...
	go test -mod=mod ./modules/...
	go test -mod=mod ./plugins/admin/controller/...
	go test -mod=mod ./plugins/admin/modules/parameter/...
	go test -mod=mod ./plugins/admin/modules/table/...
	go test -mod=mod ./plugins/admin/modules/...

## tests: helpers

import-sqlite:
	rm -rf ./tests/common/admin.db
	cp ./tests/data/admin.db ./tests/common/admin.db

import-mysql:
	mysql -h$(MYSQL_HOST) -P${MYSQL_PORT} -u${MYSQL_USER} -p${MYSQL_PWD} -e "create database if not exists \`${TEST_DB}\`"
	mysql -h$(MYSQL_HOST) -P${MYSQL_PORT} -u${MYSQL_USER} -p${MYSQL_PWD} ${TEST_DB} < ./tests/data/admin.sql

import-postgresql:
	PGPASSWORD=${POSTGRESSQL_PWD} dropdb -h ${POSTGRESSQL_HOST} -p ${POSTGRESSQL_PORT} -U ${POSTGRESSQL_USER} ${TEST_DB}
	PGPASSWORD=${POSTGRESSQL_PWD} createdb -h ${POSTGRESSQL_HOST} -p ${POSTGRESSQL_PORT} -U ${POSTGRESSQL_USER} ${TEST_DB}
	PGPASSWORD=${POSTGRESSQL_PWD} psql -h ${POSTGRESSQL_HOST} -p ${POSTGRESSQL_PORT} -d ${TEST_DB} -U ${POSTGRESSQL_USER} -f ./tests/data/admin_pg.sql

import-mssql:
	/opt/mssql-tools/bin/sqlcmd -S db_mssql -U SA -P Aa123456 -Q "RESTORE DATABASE [goadmin] FROM DISK = N'/home/data/admin_ms.bak' WITH FILE = 1, NOUNLOAD, REPLACE, RECOVERY, STATS = 5"

backup-mssql:
	docker exec mssql /opt/mssql-tools/bin/sqlcmd -S localhost -U SA -P Aa123456 -Q "BACKUP DATABASE [goadmin] TO DISK = N'/home/data/admin_ms.bak' WITH NOFORMAT, NOINIT, NAME = 'goadmin-full', SKIP, NOREWIND, NOUNLOAD, STATS = 10"

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

.PHONY: all fmt golint govet cp-mod restore-mod test black-box-test mysql-test sqlite-test import-sqlite import-mysql import-postgresql pg-test lint cilint cli
