module github.com/valhallacoin/vhcd/blockchain/stake

require (
	github.com/valhallacoin/vhcd/chaincfg v1.2.0
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/database v1.0.1
	github.com/valhallacoin/vhcd/vhcec v0.0.0-20180801202239-0761de129164
	github.com/valhallacoin/vhcd/vhcutil v1.1.1
	github.com/valhallacoin/vhcd/txscript v1.0.1
	github.com/valhallacoin/vhcd/wire v1.2.0
	github.com/decred/slog v1.0.0
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/golang/protobuf v1.2.0 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f // indirect
	golang.org/x/sys v0.0.0-20181211161752-7da8ea5c8182 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
)

replace (
	github.com/valhallacoin/vhcd/chaincfg => ../../chaincfg
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../../chaincfg/chainhash
	github.com/valhallacoin/vhcd/database => ../../database
	github.com/valhallacoin/vhcd/vhcec => ../../vhcec
	github.com/valhallacoin/vhcd/vhcec/edwards => ../../vhcec/edwards
	github.com/valhallacoin/vhcd/vhcec/secp256k1 => ../../vhcec/secp256k1
	github.com/valhallacoin/vhcd/vhcutil => ../../vhcutil
	github.com/valhallacoin/vhcd/txscript => ../../txscript
	github.com/valhallacoin/vhcd/wire => ../../wire
)
