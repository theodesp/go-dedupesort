GO := go
GODEP := godep
CONVEY_PORT ?= 9042
SOURCES := $(shell find . -type f -name "*.go")
PACKAGES :=	$(shell go list ./... | grep -v /vendor/ | grep -v /cmd)

all:    build

build:
	$(GO) build
	
	
.PHONY: godep-save
godep-save:
	$(GODEP) save $(PACKAGES)
	
.PHONY: convey
convey:
	$(GO) get github.com/smartystreets/goconvey
	goconvey -cover -port=$(CONVEY_PORT) -workDir="$(realpath .)" -depth=1

.PHONY: test
test:
	$(GO) get -t .
	$(GO) test -v .
	
.PHONY:	cover
cover:	profile.out

profile.out: $(SOURCES)
	rm -f $@
	$(GO) test -covermode=count -coverpkg=. -coverprofile=$@ .

.PHONY: doc
doc:
    godoc -http=:6060 -index
	
.PHONY: fmt
fmt:
	go fmt ./... 
