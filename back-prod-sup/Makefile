.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/InsertProduct/bootstrap function/InsertProduct/main.go
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/ListProducts/bootstrap function/ListProducts/main.go
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/CreateSupplier/bootstrap cmd/lambda/CreateSupplier/main.go
	env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o bin/ListSuppliers/bootstrap cmd/lambda/ListSuppliers/main.go

zip:
	zip -j bin/InsertProduct.zip bin/InsertProduct/bootstrap
	zip -j bin/ListProducts.zip bin/ListProducts/bootstrap
	zip -j bin/CreateSupplier.zip bin/CreateSupplier/bootstrap

clean:
	rm -rf ./bin ./vendor go.sum

deploy: build zip
	sls deploy --verbose
