module github.com/valhallacoin/vhcd/addrmgr

require (
	github.com/valhallacoin/vhcd/chaincfg/chainhash v1.0.1
	github.com/valhallacoin/vhcd/wire v1.1.0
	github.com/decred/slog v1.0.0
)

replace (
	github.com/valhallacoin/vhcd/chaincfg/chainhash => ../chaincfg/chainhash
	github.com/valhallacoin/vhcd/wire => ../wire
)
