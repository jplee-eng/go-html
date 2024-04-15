GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOCLEAN = $(GOCMD) clean
API_DIR = api
WEB_DIR = web
BINARY_NAME = gowebserver

all: test build

build-api:
	cd $(API_DIR) && $(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/main.go

build-docs:
	hugo --source $(WEB_DIR) --minify

test-api:
	cd $(API_DIR) && $(GOTEST) -v ./...

test-docs:
	hugo --source $(WEB_DIR) --buildDrafts --buildFuture --cleanDestinationDir

run-api:
	make build-api
	./$(API_DIR)/bin/$(BINARY_NAME)

run-docs:
	make build-docs
	hugo --source $(WEB_DIR) server --minify

clean:
	$(GOCLEAN)
	rm -f ./$(API_DIR)/bin/$(BINARY_NAME)
	rm -rf ./$(WEB_DIR)/public

build:
	make build-docs
	make build-api

test:
	make test-api
	make test-docs

run:
	make clean
	make run-api &
	make run-docs &
	caddy run

docker-build:
	docker build -t $(BINARY_NAME):latest .

docker-run:
	make clean
	make
	make docker-build
	docker rm -f $(BINARY_NAME) || true && docker run --name $(BINARY_NAME) -p 80:80 $(BINARY_NAME)
