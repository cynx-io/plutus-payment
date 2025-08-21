tidy:
	go mod tidy
	go fmt ./...
	fieldalignment -fix ./...
	go vet ./...
	golangci-lint run --fix ./...

.PHONY: run
run:
	make tidy
	go run main.go

install_deps:
	# These needs sudo
	# apt install build-essential -y
    # curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
	go install github.com/google/wire/cmd/wire@latest
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite


.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -f api/proto/user/*.pb.go
	rm -f bin/plutus

.PHONY: all
all: clean proto build


PROTO_SRC_DIR := api/proto/src
PROTO_GEN_DIR := api/proto/gen
MICROSERVICES := $(notdir $(wildcard $(PROTO_SRC_DIR)/*))

# Proto generation
proto-clean:
	@echo "Cleaning generated proto files..."
	rm -rf $(PROTO_GEN_DIR)/*

proto-gen:
	@echo "Generating proto files..."
	cd . && buf generate

proto: proto-clean proto-gen

build_and_push:
	docker build -t plutus-payment-dev:latest .
	docker tag plutus-payment-dev:latest derwin334/plutus-payment-dev:latest
	docker push derwin334/plutus-payment-dev:latest
