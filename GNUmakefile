MKDIR=	mkdir -p
GO=	go
PANDOC=	pandoc
CP=	cp
RM=	rm -rf

DATETIME=   "Jan 13, 2024"
VERSION=    1.0.0
LDFLAGS=    -ldflags "-X 'main.Version=$(VERSION)'"

BIN=bin/vanity
SRC=	cmd/vanity/main.go \
	vanity.go \
	go.mod \
	go.sum

all: build

build: $(BIN)

$(BIN): $(SRC)
	$(GO) build $(LDFLAGS) -o $(BIN) cmd/vanity/main.go

clean:
	$(RM) bin
