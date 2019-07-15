// Copyright (c) 2019 The Valhalla Coin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// addr2pkscript converts a Valhalla address to a public key script.
package main

import (
	"fmt"
	"os"

	"github.com/valhallacoin/vhcd/txscript"
	"github.com/valhallacoin/vhcd/vhcutil"
)

func addr2PKScript(addrStr string) ([]byte, error) {
	addr, err := vhcutil.DecodeAddress(addrStr)
	if err != nil {
		return nil, err
	}
	pkScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		return nil, err
	}
	return pkScript, nil
}

func fatal(err error) {
	fmt.Fprintf(os.Stderr, "%s: error: %s\n", os.Args[0], err)
	os.Exit(1)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s address\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	pkScript, err := addr2PKScript(os.Args[1])
	if err != nil {
		fatal(err)
	}
	fmt.Printf("%x\n", pkScript)
}
