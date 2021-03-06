// Copyright (c) 2020 The Valhalla Coin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package updater

import (
	"github.com/frankbraun/codechain/secpkg"
)

// UpToDate ensures that the package with the base name of the executing
// process is up-to-date, if it is installed as a secure package. If the
// package is not installed as a secure package a corresponding message is
// shown on stderr.
//
// UpToDate times out after a while if DNS cannot be queried and return nil.
func UpToDate() error {
	name, err := getPackageName()
	if err != nil {
		return err
	}
	return secpkg.UpToDate(name)
}
