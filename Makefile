GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = gowebserver

all: test build

build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/gowebserver

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)

run:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) -v ./cmd/gowebserver
	./bin/$(BINARY_NAME)
