module github.com/valhallacoin/vhcd/connmgr

require (
	github.com/valhallacoin/vhcd/chaincfg v1.1.1
	github.com/valhallacoin/vhcd/wire v1.2.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/valhallacoin/vhcd/chaincfg => ../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/wire => ../wire
)
