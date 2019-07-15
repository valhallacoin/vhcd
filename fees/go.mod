module github.com/valhallacoin/vhcd/fees

require (
	github.com/btcsuite/goleveldb v1.0.0
	github.com/valhallacoin/vhcd/blockchain/stake v1.1.0
	github.com/valhallacoin/vhcd/chaincfg v1.2.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/vhcutil v1.2.0
	github.com/decred/slog v1.0.0
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	golang.org/x/sys v0.0.0-20180831094639-fa5fdf94c789 // indirect
)

replace (
	github.com/valhallacoin/vhcd/chaincfg => ../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/vhcec => ../vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/vhcutil => ../vhcutil
	github.com/valhallacoin/vhcd/wire => ../wire
)
