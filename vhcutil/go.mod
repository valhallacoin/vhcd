module github.com/valhallacoin/vhcd/vhcutil

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/decred/base58 v1.0.0
	github.com/valhallacoin/vhcd/chaincfg v1.2.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/vhcec v0.0.0-20180721005212-59fe2b293f69
	github.com/valhallacoin/vhcd/vhcec/edwards v0.0.0-20181208004914-a0816cf4301f
	github.com/valhallacoin/vhcd/vhcec/secp256k1 v1.0.1
	github.com/valhallacoin/vhcd/wire v1.2.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
)

replace (
	github.com/valhallacoin/vhcd/chaincfg => ../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/vhcec => ../vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/wire => ../wire
)
