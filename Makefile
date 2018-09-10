all: run

assets:
	go-bindata -o ./template/adminlte/resource/assets.go ./template/adminlte/resource/assets/...

tmpl:
	admincli compile