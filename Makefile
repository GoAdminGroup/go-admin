TEST_CONFIG_PATH=./../common/config.json
TEST_CONFIG_PQ_PATH=./../common/config_pg.json
TEST_CONFIG_SQLITE_PATH=./../common/config_sqlite.json
ASSETS_PATH=./template/adminlte/resource/assets

all: run

assets:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	rm -rf $(ASSETS_PATH)/dist
	mkdir $(ASSETS_PATH)/dist
	mkdir $(ASSETS_PATH)/dist/js
	mkdir $(ASSETS_PATH)/dist/css
	cp $(ASSETS_PATH)/src/js/*.js $(ASSETS_PATH)/dist/js/
	cp $(ASSETS_PATH)/src/css/*.png $(ASSETS_PATH)/dist/css/
	cp -R $(ASSETS_PATH)/src/css/fonts $(ASSETS_PATH)/dist/css/
	cp -R $(ASSETS_PATH)/src/img $(ASSETS_PATH)/dist/
	cp -R $(ASSETS_PATH)/src/fonts $(ASSETS_PATH)/dist/
	make combine
	admincli compile asset
	make tmpl
	make fmt

combine:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	make combine-js
	make combine-css

combine-js:
	admincli combine js
	admincli combine js --path=$(ASSETS_PATH)/src/js/combine2/ --out=$(ASSETS_PATH)/dist/js/all_2.min.js
	admincli combine js --path=$(ASSETS_PATH)/src/js/combine3/ --out=$(ASSETS_PATH)/dist/js/form.min.js

combine-css:
	admincli combine css

tmpl:
	admincli compile tpl

fmt:
	go fmt ./adapter/...
	go fmt ./admincli/...
	go fmt ./context/...
	go fmt ./engine/...
	go fmt ./tests/...
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
	golangci-lint run