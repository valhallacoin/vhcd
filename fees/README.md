fees
=======


[![Build Status](http://img.shields.io/travis/valhallacoin/vhcd.svg)](https://travis-ci.org/valhallacoin/vhcd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/valhallacoin/vhcd/fees)

Package fees provides valhallacoin-specific methods for tracking and estimating fee
rates for new transactions to be mined into the network. Fee rate estimation has
two main goals:

- Ensuring transactions are mined within a target _confirmation range_
  (expressed in blocks);
- Attempting to minimize fees while maintaining be above restriction.

This package was started in order to resolve issue valhallacoin/vhcd#1412 and related.
See that issue for discussion of the selected approach.

This package was developed for vhcd, a full-node implementation of Valhalla which
is under active development.  Although it was primarily written for
vhcd, this package has intentionally been designed so it can be used as a
standalone package for any projects needing the functionality provided.

## Installation and Updating

```bash
$ go get -u github.com/valhallacoin/vhcd/fees
```

## License

Package vhcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
