.PHONY: build zip clean deploy

# Directories and files
BIN_DIR := bin
SRC_DIR := cmd/lambda
LAMBDA_FUNCTIONS := CreateSupplier ListSuppliers CreateFilm ListFilms

# Go build settings
GO_SETTINGS := GOARCH=arm64 GOOS=linux
LDFLAGS := -ldflags="-s -w"

build:
	export GO111MODULE=on
	@for function in $(LAMBDA_FUNCTIONS); do \
		env $(GO_SETTINGS) go build $(LDFLAGS) -o $(BIN_DIR)/$$function/bootstrap $(SRC_DIR)/$$function/main.go; \
	done

zip:
	@for function in $(LAMBDA_FUNCTIONS); do \
		zip -j $(BIN_DIR)/$$function.zip $(BIN_DIR)/$$function/bootstrap; \
	done

deploy: build zip
	sls deploy --verbose
