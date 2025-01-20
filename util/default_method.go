package util

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

var DefaultSignerFn func(prikey string, chainID *big.Int) func(address common.Address, tx *types.Transaction) (*types.Transaction, error)
var DefaultTransactOpts func(client *ethclient.Client, prikey string, chanID, gasprice *big.Int, gasLimit uint64) (*bind.TransactOpts, error)

func init() {
	DefaultSignerFn = GetDefaultSignerFn
	DefaultTransactOpts = GetDefaultTransactOpts
}

func GetDefaultSignerFn(prikey string, chainID *big.Int) func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		privateKey, err := crypto.HexToECDSA(prikey)
		if err != nil {
			return nil, err
		}
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			return nil, fmt.Errorf("failed to sign transaction: %v", err)
		}
		return signedTx, nil
	}
}

func GetDefaultTransactOpts(client *ethclient.Client, prikey string, chanID, gasprice *big.Int, gasLimit uint64) (*bind.TransactOpts, error) {
	if len(prikey) > 1 && prikey[:2] == "0x" {
		prikey = prikey[2:]
	}
	privateKey, err := crypto.HexToECDSA(prikey)
	if err != nil {
		return nil, err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return nil, err
	}
	signerFn := DefaultSignerFn(prikey, chanID)
	opts := bind.TransactOpts{
		Nonce:    big.NewInt(int64(nonce)),
		Signer:   signerFn,
		GasPrice: gasprice,
		GasLimit: gasLimit,
	}
	return &opts, nil
}

func WaitForTransactionReceipt(client *ethclient.Client, txHash common.Hash, maxRetries int, interval time.Duration) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error

	for i := 0; i < maxRetries; i++ {
		receipt, err = client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		time.Sleep(interval)
	}
	return nil, fmt.Errorf("failed to fetch transaction receipt after %d retries: %v", maxRetries, err)
}
