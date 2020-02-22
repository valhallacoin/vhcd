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

# hack to rename vhcd when installed with secpkg either to vhcd-mainnet
# or vhcd-testnet, based on secpkg build directory
vhcd=$(shell basename $(shell dirname $(CURDIR)))
ifneq ($(vhcd),vhcd-mainnet)
ifneq ($(vhcd),vhcd-testnet)
  vhcd=vhcd
endif
endif

all:
	env GO111MODULE=on go build -v -mod vendor -o $(vhcd)
	env GO111MODULE=on go build -v -mod vendor ./cmd/...

install:
	mkdir -p $(bindir)
	cp -f $(vhcd) $(bindir)/$(vhcd)
	env GO111MODULE=on GOBIN=$(bindir) go install -mod vendor -v ./cmd/...
ifneq ($(vhcd),vhcd)
	cp -f scripts/run-$(vhcd) $(bindir)/run-$(vhcd)
endif

uninstall:
	rm -f $(bindir)/$(vhcd) $(bindir)/addblock $(bindir)/addr2pkscript $(bindir)/findcheckpoint $(bindir)/gencerts $(bindir)/gennonce $(bindir)/printunixtime $(bindir)/promptsecret $(bindir)/vhcchain $(bindir)/vhcctl
ifneq ($(vhcd),vhcd)
	rm -f $(bindir)/run-$(vhcd)
endif

clean:
	rm -f $(vhcd)

test:
	env GO111MODULE=on ./run_tests.sh

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
