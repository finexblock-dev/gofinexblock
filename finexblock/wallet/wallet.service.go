package wallet

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletService struct {
	repo Repository
}

func (w *walletService) FindBlockchainByName(name string) (result *entity.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockchainByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindBlockchainByID(id uint) (result *entity.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockchainByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindCoinByID(id uint) (result *entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinByName(name string) (result *entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByCoinID(coinID uint) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockNumberByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByID(id uint) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindBlockNumberByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateBlockNumber(tx, coinID, fromBlockNumber, toBlockNumber)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByID(id uint) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindWalletByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByAddress(addr string) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindWalletByAddress(tx, addr)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindAllWallet(coinID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindAllWallet(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByCoinID(coinID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWalletByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByUserID(userID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWalletByUserID(tx, userID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) GetContractAddressByCoinID(coinID uint) (result []*entity.SmartContract, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.GetContractAddressByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertWallet(wallet *entity.Wallet) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertWallet(tx, wallet)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWallet(id uint, address string) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateWallet(tx, id, address)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransfer(walletID uint, amount decimal.Decimal, transferType entity.TransferType) (result *entity.CoinTransfer, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertCoinTransfer(tx, walletID, amount, transferType)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByID(id uint) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinTransactionByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByTxHash(txHash string) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.FindCoinTransactionByTxHash(tx, txHash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByTransferID(transferID uint) (result []*entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanCoinTransactionByTransferID(tx, transferID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByCond(transferID uint, status entity.TransactionStatus) (result []*entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanCoinTransactionByCond(tx, transferID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransaction(transferID uint, txHash string, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.InsertCoinTransaction(tx, transferID, txHash, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionHash(id uint, hash string) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateCoinTransactionHash(tx, id, hash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionStatus(id uint, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.UpdateCoinTransactionStatus(tx, id, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByStatus(status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWithdrawalRequestByStatus(tx, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByCond(coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.repo.ScanWithdrawalRequestByCond(tx, coinID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWithdrawalRequest(id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error) {
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
