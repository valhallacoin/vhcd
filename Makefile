ifneq ($(GOPATH),)
  prefix ?= $(GOPATH)
endif
prefix ?= /usr/local
exec_prefix ?= $(prefix)
ifneq ($(GOBIN),)
  bindir ?= $(GOBIN)
endif
bindir ?= $(exec_prefix)/bin

.PHONY: all install uninstall clean test update-vendor

vhcd?=$(shell basename $(CURDIR))

all:
	env GO111MODULE=on go build -v -mod vendor -o $(vhcd)
	env GO111MODULE=on go build -v -mod vendor ./cmd/...

install:
	mkdir -p $(bindir)
	cp -f $(vhcd) $(bindir)/$(vhcd)
	env GO111MODULE=on GOBIN=$(bindir) go install -mod vendor -v ./cmd/...

uninstall:
	rm -f $(bindir)/$(vhcd) $(bindir)/addblock $(bindir)/addr2pkscript $(bindir)/findcheckpoint $(bindir)/gencerts $(bindir)/gennonce $(bindir)/printunixtime $(bindir)/promptsecret $(bindir)/vhcchain $(bindir)/vhcctl

clean:
	rm -f $(vhcd)

test:
	env GO111MODULE=on ./run_tests.sh

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
