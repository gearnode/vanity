MKDIR=	mkdir -p
GO=	go
PANDOC=	pandoc
CP=	cp
RM=	rm -rf

DATETIME=   "Jan 13, 2024"
VERSION=    1.0.0
LDFLAGS=    -ldflags "-X 'main.Version=$(VERSION)'"

BIN=	bin/vanity
SRC=	cmd/vanity/main.go \
	vanity.go \
	go.mod \
	go.sum

all: build man

build: $(BIN)

.PHONY: man
man:
	@$(MKDIR) man
	$(PANDOC) \
		--standalone \
		--to man \
		-M footer=$(VERSION) \
		-M date=$(DATETIME) \
		doc/vanity.1.md \
		-o man/vanity.1
	$(PANDOC) \
		--standalone \
		--to man \
		-M footer=$(VERSION) \
		-M date=$(DATETIME) \
		doc/vanity.conf.5.md \
		-o man/vanity.conf.5

$(BIN): $(SRC)
	$(GO) build $(LDFLAGS) -o $(BIN) cmd/vanity/main.go

clean:
	$(RM) bin
	$(RM) man
