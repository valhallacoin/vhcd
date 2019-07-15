vhcd
====

[![Build Status](https://travis-ci.org/valhallacoin/vhcd.png?branch=master)](https://travis-ci.org/valhallacoin/vhcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/valhallacoin/vhcd)
[![Go Report Card](https://goreportcard.com/badge/github.com/valhallacoin/vhcd)](https://goreportcard.com/report/github.com/valhallacoin/vhcd)

## Valhalla Overview

Valhalla is a blockchain-based cryptocurrency with a strong focus on community
input, open governance, and sustainable funding for development. It utilizes a
hybrid proof-of-work and proof-of-stake mining system to ensure that a small
group cannot dominate the flow of transactions or make changes to Valhalla without
the input of the community.  A unit of the currency is called a `valhallacoin` (VHC).

https://valhallacoin.org

## Latest Downloads

https://valhallacoin.org/downloads

## What is vhcd?

vhcd is a full node implementation of Valhalla written in Go (golang).

It acts as a fully-validating chain daemon for the Valhalla cryptocurrency.  vhcd
maintains the entire past transactional ledger of Valhalla and allows relaying of
transactions to other Valhalla nodes around the world.

This software is currently under active development.  It is extremely stable and
has been in production use since February 2016.

The software was originally forked from [btcd](https://github.com/btcsuite/btcd),
which is a bitcoin full node implementation that is still under active
development.  To gain the benefit of btcd's ongoing upgrades, including improved
peer and connection handling, database optimization, and other blockchain
related technology improvements, vhcd is continuously synced with the btcd
codebase.

## What is a full node?

The term 'full node' is short for 'fully-validating node' and refers to software
that fully validates all transactions and blocks, as opposed to trusting a 3rd
party.  In addition to validating transactions and blocks, nearly all full nodes
also participate in relaying transactions and blocks to other full nodes around
the world, thus forming the peer-to-peer network that is the backbone of the
Valhalla cryptocurrency.

The full node distinction is important, since full nodes are not the only type
of software participating in the Valhalla peer network. For instance, there are
'lightweight nodes' which rely on full nodes to serve the transactions, blocks,
and cryptographic proofs they require to function, as well as relay their
transactions to the rest of the global network.

## Why run vhcd?

As described in the previous section, the Valhalla cryptocurrency relies on having
a peer-to-peer network of nodes that fully validate all transactions and blocks
and then relay them to other full nodes.

Running a full node with vhcd contributes to the overall security of the
network, increases the available paths for transactions and blocks to relay,
and helps ensure there are an adequate number of nodes available to serve
lightweight clients, such as Simplified Payment Verification (SPV) wallets.

Without enough full nodes, the network could be unable to expediently serve
users of lightweight clients which could force them to have to rely on
centralized services that significantly reduce privacy and are vulnerable to
censorship.

In terms of individual benefits, since vhcd fully validates every block and
transaction, it provides the highest security and privacy possible when used in
conjunction with a wallet that also supports directly connecting to it in full
validation mode, such as [vhcwallet (CLI)](https://github.com/valhallacoin/vhcwallet)
and [Valhallaiton (GUI)](https://github.com/valhallacoin/valhallacoiniton).

## Minimum Recommended Specifications (vhcd only)

* 10 GB disk space (as of September 2018, increases over time)
* 1GB memory (RAM)
* ~150MB/day download, ~1.5GB/day upload
  * Plus one-time initial download of the entire block chain
* Windows 7/8.x/10 (server preferred), macOS, Linux
* High uptime

## Getting Started

So, you've decided to help the network by running a full node.  Great!  Running
vhcd is simple.  All you need to do is install vhcd on a machine that is
connected to the internet and meets the minimum recommended specifications, and
launch it.

Also, make sure your firewall is configured to allow inbound connections to port
9208.

<a name="Installation" />

## Installing and updating

### Binaries (Windows/Linux/macOS)

Binary releases are provided for common operating systems and architectures:

https://valhallacoin.org/downloads

### Build from source (all platforms)

Building or updating from source requires the following build dependencies:

- **Go 1.10 or 1.11**

  Installation instructions can be found here: https://golang.org/doc/install.
  It is recommended to add `$GOPATH/bin` to your `PATH` at this point.

- **Vgo (Go 1.10 only)**

  The `GO111MODULE` experiment is used to manage project dependencies and
  provide reproducible builds.  The module experiment is provided by the Go 1.11
  toolchain, but the Go 1.10 toolchain does not provide any module support.  To
  perform module-aware builds with Go 1.10,
  [vgo](https://godoc.org/golang.org/x/vgo) (a drop-in replacement for the go
  command) must be used instead.

- **Git**

  Installation instructions can be found at https://git-scm.com or
  https://gitforwindows.org.

To build and install from a checked-out repo, run `go install . ./cmd/...` in
the repo's root directory.  Some notes:

* Set the `GO111MODULE=on` environment variable if using Go 1.11 and building
  from within `GOPATH`.

* Replace `go` with `vgo` when using Go 1.10.

* The `vhcd` executable will be installed to `$GOPATH/bin`.  `GOPATH`
  defaults to `$HOME/go` (or `%USERPROFILE%\go` on Windows) if unset.


### Example of obtaining and building from source on Windows 10 with Go 1.11:

```PowerShell
PS> git clone https://github.com/valhallacoin/vhcd $env:USERPROFILE\src\vhcd
PS> cd $env:USERPROFILE\src\vhcd
PS> go install . .\cmd\...
PS> & "$(go env GOPATH)\bin\vhcd" -V

```

### Example of obtaining and building from source on Linux with Go 1.10:

```bash
$ git clone https://github.com/valhallacoin/vhcd ~/src/vhcd
$ cd ~/src/vhcd
$ vgo install . ./cmd/...
$ $(vgo env GOPATH)/bin/vhcd -V
```

## Docker

### Running vhcd

You can run a valhallacoin node from inside a docker container.  To build the image
yourself, use the following command:

```
docker build -t valhallacoin/vhcd .
```

Or you can create an alpine based image (requires Docker 17.05 or higher):

```
docker build -t valhallacoin/vhcd:alpine -f Dockerfile.alpine .
```

You can then run the image using:

```
docker run valhallacoin/vhcd
```

You may wish to use an external volume to customise your config and persist the
data in an external volume:

```
docker run --rm -v /home/user/vhcdata:/root/.vhcd/data valhallacoin/vhcd
```

For a minimal image, you can use the valhallacoin/vhcd:alpine tag.  This is typically
a more secure option while also being a much smaller image.

You can run vhcctl from inside the image.  For example, run an image (mounting
your data from externally) with:

```
docker run --rm -ti --name=vhcd-1 -v /home/user/.vhcd:/root/.vhcd \
  valhallacoin/vhcd:alpine
```

And then run vhcctl commands against it.  For example:

```
docker exec -ti vhcd-1 vhcctl getbestblock
```

### Running Tests

All tests and linters may be run in a docker (or podman) container using the
script `run_tests.sh` by specifying either `docker` or `podman` as the first
parameter.  This script defaults to using the current latest supported version
of Go, but it also respects the `GOVERSION` environment variable set to the
major version of Go to allow testing on a previous version of Go.  Generally,
Valhalla only supports the current and previous major versions of Go.

```
./run_tests.sh docker
```

To run the tests locally without docker on the latest supported version of Go:

```
./run_tests.sh
```

To run the tests locally without docker on Go 1.10:

```
GOVERSION=1.10 ./run_tests.sh
```

## Contact

If you have any further questions you can find us at:

https://valhallacoin.org/community

## Issue Tracker

The [integrated github issue tracker](https://github.com/valhallacoin/vhcd/issues)
is used for this project.

## Documentation

The documentation for vhcd is a work-in-progress.  It is located in the
[docs](https://github.com/valhallacoin/vhcd/tree/master/docs) folder.

## License

vhcd is licensed under the [copyfree](http://copyfree.org) ISC License.
