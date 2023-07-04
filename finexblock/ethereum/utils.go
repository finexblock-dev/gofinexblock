package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// NewTransaction returns a new transaction.
func NewTransaction(
	chainId, gasTipCap, gasFeeCap, value *big.Int,
	nonce, gasLimit uint64,
	to common.Address,
) *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &to,
		Value:     value,
	})
}

// GasLimit returns the gas limit.
func GasLimit() uint64 {
	return uint64(42000)
}