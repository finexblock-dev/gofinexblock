package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/finexblock-dev/gofinexblock/finexblock/ethereum/hdwallet"
)

func HDWallet(mnemonic string) (*hdwallet.Wallet, error) {
	return hdwallet.NewFromMnemonic(mnemonic)
}

func DerivedAccount(wallet *hdwallet.Wallet, path accounts.DerivationPath) (*accounts.Account, error) {
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get account : %v", err)
	}
	return &account, nil
}

func DerivationPath(userID, basePath string) accounts.DerivationPath {
	return hdwallet.MustParseDerivationPath(fmt.Sprintf("%v'/%v", basePath, userID))
}

func PrivateKeyHex(wallet *hdwallet.Wallet, userID string) (string, error) {
	path := DerivationPath(userID, hdwallet.DefaultBaseDerivationPath.String())
	account, err := DerivedAccount(wallet, path)
	if err != nil {
		return "", err
	}
	return wallet.PrivateKeyHex(*account)
}