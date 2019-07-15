module github.com/valhallacoin/vhcd/gcs

require (
	github.com/dchest/blake256 v1.0.0
	github.com/dchest/siphash v1.2.0
	github.com/valhallacoin/vhcd/blockchain/stake v1.0.1
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/txscript v1.0.1
	github.com/valhallacoin/vhcd/wire v1.2.0
)

replace (
	github.com/valhallacoin/vhcd/blockchain/stake => ../blockchain/stake
	github.com/valhallacoin/vhcd/chaincfg => ../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/database => ../database
	github.com/valhallacoin/vhcd/vhcec => ../vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/vhcutil => ../vhcutil
	github.com/valhallacoin/vhcd/txscript => ../txscript
	github.com/valhallacoin/vhcd/wire => ../wire
)
