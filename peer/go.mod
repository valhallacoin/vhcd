module github.com/valhallacoin/vhcd/peer

require (
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/davecgh/go-spew v1.1.0
	github.com/valhallacoin/vhcd/blockchain v1.0.1
	github.com/valhallacoin/vhcd/chaincfg v1.2.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/txscript v1.0.1
	github.com/valhallacoin/vhcd/wire v1.2.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/valhallacoin/vhcd/blockchain => ../blockchain
	github.com/valhallacoin/vhcd/blockchain/stake => ../blockchain/stake
	github.com/valhallacoin/vhcd/chaincfg => ../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/database => ../database
	github.com/valhallacoin/vhcd/vhcec => ../vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../vhcec/secp256k1
	github.com/valhallacoin/vhcd/vhcutil => ../vhcutil
	github.com/valhallacoin/vhcd/gcs => ../gcs
	github.com/valhallacoin/vhcd/txscript => ../txscript
	github.com/valhallacoin/vhcd/wire => ../wire
)
