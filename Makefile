all: run

asserts:
	go-bindata -o ./template/adminlte/resource/assets.go ./template/adminlte/resource/assets/...