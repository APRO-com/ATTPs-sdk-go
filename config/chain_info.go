package config

import "math/big"

type ChainInfo struct {
	Url     string
	ChainID *big.Int
}

var (
	BSC_TEST_NET = ChainInfo{
		Url:     "https://bsc-testnet-rpc.publicnode.com",
		ChainID: big.NewInt(97),
	}

	BSC_MAIN_NET = ChainInfo{
		Url:     "https://binance.llamarpc.com",
		ChainID: big.NewInt(56),
	}
)
