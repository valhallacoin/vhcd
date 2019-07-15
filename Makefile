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

all:
	env GO111MODULE=on go build -v . ./cmd/...

install:
	env GO111MODULE=on GOBIN=$(bindir) go install -v . ./cmd/...

uninstall:
	rm -f $(bindir)/vhcd $(bindir)/addblock $(bindir)/findcheckpoint $(bindir)/gencerts $(bindir)/promptsecret $(bindir)/vhctl

clean:
	rm -f vhcd

test:
	env GO111MODULE=on ./run_tests.sh

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
