package wallet

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/wallet"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletService struct {
	repo Repository
}

func (w *walletService) FindBlockchainByName(name string) (result *wallet.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockchainByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindBlockchainByID(id uint) (result *wallet.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockchainByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindCoinByID(id uint) (result *wallet.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinByName(name string) (result *wallet.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByCoinID(coinID uint) (result *wallet.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockNumberByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByID(id uint) (result *wallet.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockNumberByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *wallet.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateBlockNumber(tx, coinID, fromBlockNumber, toBlockNumber)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByID(id uint) (result *wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindWalletByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByAddress(addr string) (result *wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindWalletByAddress(tx, addr)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindAllWallet(coinID uint) (result []*wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindAllWallet(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByCoinID(coinID uint) (result []*wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWalletByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByUserID(userID uint) (result []*wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWalletByUserID(tx, userID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertWallet(wallet *wallet.Wallet) (result *wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertWallet(tx, wallet)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWallet(id uint, address string) (result *wallet.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateWallet(tx, id, address)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransfer(walletID uint, amount decimal.Decimal, transferType wallet.TransferType) (result *wallet.CoinTransfer, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertCoinTransfer(tx, walletID, amount, transferType)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByID(id uint) (result *wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinTransactionByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByTxHash(txHash string) (result *wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinTransactionByTxHash(tx, txHash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByTransferID(transferID uint) (result []*wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanCoinTransactionByTransferID(tx, transferID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByCond(transferID uint, status wallet.TransactionStatus) (result []*wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanCoinTransactionByCond(tx, transferID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransaction(transferID uint, txHash string, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertCoinTransaction(tx, transferID, txHash, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionHash(id uint, hash string) (result *wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateCoinTransactionHash(tx, id, hash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionStatus(id uint, txStatus wallet.TransactionStatus) (result *wallet.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateCoinTransactionStatus(tx, id, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByStatus(status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWithdrawalRequestByStatus(tx, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByCond(coinID uint, status wallet.WithdrawalStatus) (result []*wallet.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWithdrawalRequestByCond(tx, coinID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWithdrawalRequest(id uint, state wallet.WithdrawalStatus) (result *wallet.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateWithdrawalRequest(tx, id, state)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func newWalletService(repo Repository) *walletService {
	return &walletService{repo: repo}
}

func (w *walletService) Ctx() context.Context {
	return context.Background()
}

func (w *walletService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (w *walletService) Tx(level sql.IsolationLevel) *gorm.DB {
	return w.repo.Tx(level)
}

func (w *walletService) Conn() *gorm.DB {
	return w.repo.Conn()
}
