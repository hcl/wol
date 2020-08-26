GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=wol
# export GOOS=linux
# export GOARCH=arm

all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
install:
	$(GOINSTALL)