module github.com/valhallacoin/vhcd/chaincfg

require (
	github.com/davecgh/go-spew v1.1.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/vhcec/edwards v0.0.0-20181208004914-a0816cf4301f
	github.com/valhallacoin/vhcd/vhcec/secp256k1 v1.0.1
	github.com/valhallacoin/vhcd/wire v1.2.0
)

replace (
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ./chainhash
	github.com/valhallacoin/vhcd/vhcec/edwards => ../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/wire => ../wire
)
