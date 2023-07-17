package wallet

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletRepository struct {
	db *gorm.DB
}

func (w *walletRepository) FindWithdrawalRequestByID(tx *gorm.DB, id uint) (*entity.WithdrawalRequest, error) {
	var _withdrawalRequest *entity.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).First(&_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletRepository) ScanWithdrawalRequestByStatus(tx *gorm.DB, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error) {
	var _withdrawalRequest []*entity.WithdrawalRequest
	var table *entity.WithdrawalRequest
	var err error

	if err = tx.Table(table.TableName()).Where("status = ?", status).Find(&_withdrawalRequest).Error; err != nil {
		return nil, err
	}

	return _withdrawalRequest, nil
}

func (w *walletRepository) ScanWithdrawalRequestByCond(tx *gorm.DB, coinID uint, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error) {
	var withdrawalRequests []*entity.WithdrawalRequest
	var table *entity.WithdrawalRequest

	err := tx.Table(table.TableName()).
		Select("withdrawal_request.*").
		Joins("join coin_transfer on coin_transfer.id = withdrawal_request.coin_transfer_id").
		Joins("join wallet on wallet.id = coin_transfer.wallet_id").
		Where("wallet.coin_id = ?", coinID).
		Where("withdrawal_request.status = ?", status).
		Find(&withdrawalRequests).Error

	if err != nil {
		return nil, err
	}

	return withdrawalRequests, nil
}

func (w *walletRepository) UpdateWithdrawalRequest(tx *gorm.DB, id uint, state entity.WithdrawalStatus) (*entity.WithdrawalRequest, error) {
	var _withdrawalRequest *entity.WithdrawalRequest
	var err error

	if err = tx.Table(_withdrawalRequest.TableName()).Where("id = ?", id).Update("status", state).Error; err != nil {
		return nil, err
	}

	return w.FindWithdrawalRequestByID(tx, id)
}

func (w *walletRepository) FindWalletByID(tx *gorm.DB, id uint) (*entity.Wallet, error) {
	var _wallet *entity.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("id = ?", id).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) FindWalletByAddress(tx *gorm.DB, addr string) (*entity.Wallet, error) {
	var _wallet *entity.Wallet
	var err error

	if err = tx.Table(_wallet.TableName()).Where("address = ?", addr).First(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) FindAllWallet(tx *gorm.DB, coinID uint) ([]*entity.Wallet, error) {
	var _wallet []*entity.Wallet
	var _table *entity.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) ScanWalletByCoinID(tx *gorm.DB, coinID uint) ([]*entity.Wallet, error) {
	var _wallet []*entity.Wallet
	var _table *entity.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) GetContractAddressByCoinID(tx *gorm.DB, coinID uint) ([]*entity.SmartContract, error) {
	var _smartContract []*entity.SmartContract
	var _table *entity.SmartContract
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_id = ?", coinID).Find(&_smartContract).Error; err != nil {
		return nil, err
	}

	return _smartContract, nil
}

func (w *walletRepository) ScanWalletByUserID(tx *gorm.DB, userID uint) ([]*entity.Wallet, error) {
	var _wallet []*entity.Wallet
	var _table *entity.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("user_id = ?", userID).Find(&_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) InsertWallet(tx *gorm.DB, _wallet *entity.Wallet) (*entity.Wallet, error) {
	var err error

	if err = tx.Table(_wallet.TableName()).Create(_wallet).Error; err != nil {
		return nil, err
	}

	return _wallet, nil
}

func (w *walletRepository) UpdateWallet(tx *gorm.DB, id uint, address string) (*entity.Wallet, error) {
	var _table *entity.Wallet
	var err error

	if err = tx.Table(_table.TableName()).Where("id = ?", id).Update("address", address).Error; err != nil {
		return nil, err
	}

	return w.FindWalletByID(tx, id)
}

func (w *walletRepository) FindSmartContractByCoinID(tx *gorm.DB, coinID uint) (*entity.SmartContract, error) {
	var smartContract *entity.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("coin_id = ?", coinID).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}

func (w *walletRepository) FindSmartContractByID(tx *gorm.DB, id uint) (*entity.SmartContract, error) {
	var smartContract *entity.SmartContract
	var err error

	if err = tx.Table(smartContract.TableName()).Where("id = ?", id).First(&smartContract).Error; err != nil {
		return nil, err
	}

	return smartContract, nil
}

func (w *walletRepository) InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error) {
	var coinTransfer *entity.CoinTransfer
	var err error

	coinTransfer = &entity.CoinTransfer{
		WalletID:     walletID,
		Amount:       amount,
		TransferType: transferType,
	}

	if err = tx.Table(coinTransfer.TableName()).Create(coinTransfer).Error; err != nil {
		return nil, err
	}

	return coinTransfer, nil
}

func (w *walletRepository) ScanCoinTransactionByTransferID(tx *gorm.DB, transferID uint) ([]*entity.CoinTransaction, error) {
	var coinTransaction []*entity.CoinTransaction
	var _table *entity.CoinTransaction
	var err error

	if err = tx.Table(_table.TableName()).Where("coin_transfer_id = ?", transferID).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error) {
	var err error
	var coinTransaction *entity.CoinTransaction

	coinTransaction = &entity.CoinTransaction{
		CoinTransferID: transferID,
		TxHash:         txHash,
		Status:         txStatus,
	}

	if err = tx.Table(coinTransaction.TableName()).Create(coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) FindCoinTransactionByTxHash(tx *gorm.DB, txHash string) (*entity.CoinTransaction, error) {
	var coinTransaction *entity.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("tx_hash = ?", txHash).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) FindCoinTransactionByID(tx *gorm.DB, id uint) (*entity.CoinTransaction, error) {
	var coinTransaction *entity.CoinTransaction
	var err error

	if err = tx.Table(coinTransaction.TableName()).Where("id = ?", id).First(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) ScanCoinTransactionByCond(tx *gorm.DB, transferID uint, status entity.TransactionStatus) ([]*entity.CoinTransaction, error) {
	var coinTransaction []*entity.CoinTransaction
	var table *entity.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("coin_transfer_id = ?", transferID).Where("status = ?", status).Find(&coinTransaction).Error; err != nil {
		return nil, err
	}

	return coinTransaction, nil
}

func (w *walletRepository) UpdateCoinTransactionHash(tx *gorm.DB, id uint, hash string) (*entity.CoinTransaction, error) {
	var table *entity.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("tx_hash", hash).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByTxHash(tx, hash)
}

func (w *walletRepository) UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error) {
	var table *entity.CoinTransaction
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("status", txStatus).Error; err != nil {
		return nil, err
	}

	return w.FindCoinTransactionByID(tx, id)
}

func (w *walletRepository) FindCoinByID(tx *gorm.DB, id uint) (*entity.Coin, error) {
	var coin *entity.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("id = ?", id).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (w *walletRepository) FindCoinByName(tx *gorm.DB, name string) (*entity.Coin, error) {
	var coin *entity.Coin
	var err error

	if err = tx.Table(coin.TableName()).Where("name = ?", name).First(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (w *walletRepository) FindBlockchainByName(tx *gorm.DB, name string) (*entity.Blockchain, error) {
	var blockchain *entity.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("name = ?", name).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (w *walletRepository) FindBlockchainByID(tx *gorm.DB, id uint) (*entity.Blockchain, error) {
	var blockchain *entity.Blockchain
	var err error

	if err = tx.Table(blockchain.TableName()).Where("id = ?", id).First(&blockchain).Error; err != nil {
		return nil, err
	}

	return blockchain, nil
}

func (w *walletRepository) FindBlockNumberByCoinID(tx *gorm.DB, coinID uint) (*entity.BlockNumber, error) {
	var blockNumber *entity.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("coin_id = ?", coinID).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletRepository) FindBlockNumberByID(tx *gorm.DB, id uint) (*entity.BlockNumber, error) {
	var blockNumber *entity.BlockNumber
	var err error

	if err = tx.Table(blockNumber.TableName()).Where("id = ?", id).First(&blockNumber).Error; err != nil {
		return nil, err
	}

	return blockNumber, nil
}

func (w *walletRepository) UpdateBlockNumber(tx *gorm.DB, coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (*entity.BlockNumber, error) {
	var err error
	var table *entity.BlockNumber

	if err = tx.Table(table.TableName()).Where("coin_id = ?", coinID).Update("from_block", fromBlockNumber).Update("to_block", toBlockNumber).Error; err != nil {
		return nil, err
	}

	return w.FindBlockNumberByCoinID(tx, coinID)
}

func (w *walletRepository) Conn() *gorm.DB {
	return w.db
}

func (w *walletRepository) Ctx() context.Context {
	return context.Background()
}

func (w *walletRepository) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (w *walletRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return w.db.Begin(&sql.TxOptions{Isolation: level})
}

func newWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db: db}
}
