PREFIX = /usr/local
BINDIR = $(PREFIX)/bin
MANDIR = $(PREFIX)/share/man

MKDIR = mkdir -p
GO = go
PANDOC = pandoc
INSTALL = install
RM = rm -f

DATETIME=   "Apr 17, 2024"
VERSION=    1.0.1
LDFLAGS = -ldflags "-X 'main.Version=$(VERSION)'"

BIN = bin/vanity

.PHONY: all build man install uninstall clean

all: build man

build:
	$(GO) build $(LDFLAGS) -o $(BIN) cmd/vanity/main.go

man:
	@$(MKDIR) man
	$(PANDOC) --standalone --to man -M footer=$(VERSION) -M date=$(DATETIME) doc/vanity.1.md -o man/vanity.1
	$(PANDOC) --standalone --to man -M footer=$(VERSION) -M date=$(DATETIME) doc/vanity.conf.5.md -o man/vanity.conf.5

install: build man
	$(INSTALL) -m 755 $(BIN) $(BINDIR)/vanity
	$(INSTALL) -m 644 man/vanity.1 $(MANDIR)/man1/vanity.1
	$(INSTALL) -m 644 man/vanity.conf.5 $(MANDIR)/man5/vanity.conf.5

uninstall:
	$(RM) $(BINDIR)/privatebin
	$(RM) $(MANDIR)/man1/vanity.1
	$(RM) $(MANDIR)/man5/vanity.conf.5

clean:
	$(RM) -r bin man
