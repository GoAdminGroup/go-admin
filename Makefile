TEST_CONFIG_PATH=./../common/config.json
TEST_CONFIG_PQ_PATH=./../common/config_pg.json

all: run

assets:
	find ./ -name ".DS_Store" -depth -exec rm {} \;
	rm -rf ./template/adminlte/resource/assets/dist
	mkdir ./template/adminlte/resource/assets/dist
	mkdir ./template/adminlte/resource/assets/dist/js
	mkdir ./template/adminlte/resource/assets/dist/css
	cp ./template/adminlte/resource/assets/src/js/*.js ./template/adminlte/resource/assets/dist/js/
	cp ./template/adminlte/resource/assets/src/css/blue.png ./template/adminlte/resource/assets/dist/css/blue.png
	cp -R ./template/adminlte/resource/assets/src/css/fonts ./template/adminlte/resource/assets/dist/css/
	cp -R ./template/adminlte/resource/assets/src/img ./template/adminlte/resource/assets/dist/
	cp -R ./template/adminlte/resource/assets/src/fonts ./template/adminlte/resource/assets/dist/
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
	admincli combine js --path=./template/adminlte/resource/assets/src/js/combine2/ --out=./template/adminlte/resource/assets/dist/js/all_2.min.js
	admincli combine js --path=./template/adminlte/resource/assets/src/js/combine3/ --out=./template/adminlte/resource/assets/dist/js/form.min.js

combine-css:
	admincli combine css

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
	dropdb -U postgres go-admin-test
	createdb -U postgres go-admin-test
	psql -d go-admin-test -U postgres -f ./examples/datamodel/admin.pgsql

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