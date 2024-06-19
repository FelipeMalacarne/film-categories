.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/InsertProduct/bootstrap function/InsertProduct/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/ListProducts/bootstrap function/ListProducts/main.go

clean:
	rm -rf ./bin ./vendor go.sum

deploy: clean build
	sls deploy --verbose
