module github.com/valhallacoin/vhcd

require (
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412
	github.com/btcsuite/go-socks v0.0.0-20170105172521-4720035b7bfd
	github.com/btcsuite/goleveldb v1.0.0
	github.com/btcsuite/snappy-go v1.0.0
	github.com/btcsuite/winsvc v1.0.0
	github.com/davecgh/go-spew v1.1.0
	github.com/dchest/blake256 v1.0.0
	github.com/dchest/siphash v1.2.1
	github.com/decred/base58 v1.0.0
	github.com/valhallacoin/vhcd/addrmgr v1.0.2
	github.com/valhallacoin/vhcd/blockchain v1.1.1
	github.com/valhallacoin/vhcd/blockchain/stake v1.1.0
	github.com/valhallacoin/vhcd/certgen v1.0.2
	github.com/valhallacoin/vhcd/chaincfg v1.3.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/connmgr v1.0.2
	github.com/valhallacoin/vhcd/database v1.0.3
	github.com/valhallacoin/vhcd/vhcec v0.0.0-20180801202239-0761de129164
	github.com/valhallacoin/vhcd/vhcec/secp256k1 v1.0.1
	github.com/valhallacoin/vhcd/vhcjson v1.1.0
	github.com/valhallacoin/vhcd/vhcutil v1.2.0
	github.com/valhallacoin/vhcd/fees v1.0.0
	github.com/valhallacoin/vhcd/gcs v1.0.2
	github.com/valhallacoin/vhcd/hdkeychain v1.1.1
	github.com/valhallacoin/vhcd/mempool v1.1.0
	github.com/valhallacoin/vhcd/mining v1.1.0
	github.com/valhallacoin/vhcd/peer v1.1.0
	github.com/valhallacoin/vhcd/rpcclient v1.1.0
	github.com/valhallacoin/vhcd/txscript v1.0.2
	github.com/valhallacoin/vhcd/wire v1.2.0
	github.com/decred/slog v1.0.0
	github.com/gorilla/websocket v1.2.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20180718160520-a2144134853f
	golang.org/x/sys v0.0.0-20181211161752-7da8ea5c8182
)

replace (
	github.com/valhallacoin/vhcd/addrmgr => ./addrmgr
	github.com/valhallacoin/vhcd/blockchain => ./blockchain
	github.com/valhallacoin/vhcd/blockchain/stake => ./blockchain/stake
	github.com/valhallacoin/vhcd/certgen => ./certgen
	github.com/valhallacoin/vhcd/chaincfg => ./chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ./chaincfg/chainhash
	github.com/valhallacoin/vhcd/connmgr => ./connmgr
	github.com/valhallacoin/vhcd/database => ./database
	github.com/valhallacoin/vhcd/vhcec => ./vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ./vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ./vhcec/secp256k1
	github.com/valhallacoin/vhcd/vhcjson => ./vhcjson
	github.com/valhallacoin/vhcd/vhcutil => ./vhcutil
	github.com/valhallacoin/vhcd/fees => ./fees
	github.com/valhallacoin/vhcd/gcs => ./gcs
	github.com/valhallacoin/vhcd/hdkeychain => ./hdkeychain
	github.com/valhallacoin/vhcd/limits => ./limits
	github.com/valhallacoin/vhcd/mempool => ./mempool
	github.com/valhallacoin/vhcd/mining => ./mining
	github.com/valhallacoin/vhcd/peer => ./peer
	github.com/valhallacoin/vhcd/rpcclient => ./rpcclient
	github.com/valhallacoin/vhcd/txscript => ./txscript
	github.com/valhallacoin/vhcd/wire => ./wire
)
