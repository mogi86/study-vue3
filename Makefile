.PHONY: go-api-server-build
go-api-server-build:
	cd go-api-server && go build -o ./../server main.go
