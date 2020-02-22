// Copyright (c) 2013-2016 The btcsuite developers
// Copyright (c) 2016-2018 The Decred developers
// Copyright (c) 2020 The Valhalla Coin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package updater

import (
	"errors"
)

// ErrNeedsUpdate is returned if vhcd needs an update.
var ErrNeedsUpdate = errors.New("updater: vhcd needs an update")
