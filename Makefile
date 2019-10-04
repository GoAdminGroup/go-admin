TEST_CONFIG_PATH=./../common/config.json
TEST_CONFIG_PQ_PATH=./../common/config_pg.json

all: run

assets:
	go-bindata -o ./template/adminlte/resource/assets.go ./template/adminlte/resource/assets/...

tmpl:
	admincli compile tpl

fmt:
	go fmt ./adapter/...
	go fmt ./admincli/...
	go fmt ./context/...
	go fmt ./engine/...
	go fmt ./examples/...
	go fmt ./modules/...
	go fmt ./plugins/...
	go fmt ./template/...

deps:
	go get github.com/kardianos/govendor
	govendor sync

test:
	make mysql-test
	make pg-test

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

import-mysql:
	mysql -uroot -proot go-admin-test < ./examples/datamodel/admin.sql

import-postgresql:
	createdb -U postgres go-admin-test
	psql -d go-admin-test -U postgres -f ./examples/datamodel/admin.pgsql

pg-test:
	make import-postgresql
	gotest -v ./tests/gin/... -args $(TEST_CONFIG_PQ_PATH)
	dropdb -U postgres go-admin-test
	make import-postgresql
	gotest -v ./tests/beego/... -args $(TEST_CONFIG_PQ_PATH)
	dropdb -U postgres go-admin-test
	make import-postgresql
	gotest -v ./tests/buffalo/... -args $(TEST_CONFIG_PQ_PATH)
	dropdb -U postgres go-admin-test
	make import-postgresql
	gotest -v ./tests/chi/... -args $(TEST_CONFIG_PQ_PATH)
	dropdb -U postgres go-admin-test
	make import-postgresql
	gotest -v ./tests/echo/... -args $(TEST_CONFIG_PQ_PATH)
	dropdb -U postgres go-admin-test
	make import-postgresql
	gotest -v ./tests/gorilla/... -args $(TEST_CONFIG_PQ_PATH)

lint:
	golangci-lint run