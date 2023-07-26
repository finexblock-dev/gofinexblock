package wallet

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/wallet/structs"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletService struct {
	manager          trade.Manager
	walletRepository Repository
	userRepository   user.Repository
}

func (w *walletService) FindAllUserAssets(id uint) (result []*structs.Asset, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		var wallets []*entity.Wallet
		var coins []*entity.Coin
		var _user *entity.User
		var coinIDs []uint
		var balance decimal.Decimal

		_user, err = w.userRepository.FindUserByID(tx, id)
		if err != nil {
			return err
		}

		wallets, err = w.walletRepository.ScanWalletByUserID(tx, id)
		if err != nil {
			return err
		}

		for _, w := range wallets {
			coinIDs = append(coinIDs, w.CoinID)
		}

		coins, err = w.walletRepository.FindManyCoinByID(tx, coinIDs)
		if err != nil {
			return err
		}

		for _, c := range coins {
			balance, err = w.manager.GetBalance(_user.UUID, c.Name)
			if err != nil {
				return err
			}

			result = append(result, &structs.Asset{
				CoinID:  c.ID,
				Balance: balance,
			})
		}

		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}
	return result, nil
}

func (w *walletService) FindManyCoinByID(ids []uint) (result []*entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindManyCoinByID(tx, ids)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}
	return result, nil
}

func (w *walletService) FindManyCoinByName(names []string) (result []*entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindManyCoinByName(tx, names)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}
	return result, nil
}

func (w *walletService) BalanceUpdateInBatch(event []*grpc_order.BalanceUpdate) (err error) {
	return w.Conn().Transaction(func(tx *gorm.DB) error {
		var _transfer *entity.CoinTransfer
		var coinTransfers []*entity.CoinTransfer

		var _user *entity.User
		var users []*entity.User

		var _coin *entity.Coin
		var coins []*entity.Coin

		var _wallet *entity.Wallet
		var wallets []*entity.Wallet

		var userUUIDs []string
		var currencies []string

		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var coinCache = cache.NewDefaultKeyValueStore[entity.Coin](len(event))
		var walletCache = cache.NewDefaultKeyValueStore[entity.Wallet](len(event))

		for _, e := range event {
			if _user, err = userCache.Get(e.GetUserUUID()); err != nil {
				userUUIDs = append(userUUIDs, e.GetUserUUID())
				_ = userCache.Set(e.GetUserUUID(), new(entity.User))
			}

			if _coin, err = coinCache.Get(e.GetCurrency().String()); err != nil {
				currencies = append(currencies, e.GetCurrency().String())
				_ = coinCache.Set(e.GetCurrency().String(), new(entity.Coin))
			}
		}

		users, err = w.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return fmt.Errorf("failed to find users: %w", err)
		}

		for _, u := range users {
			_ = userCache.Set(u.UUID, u)
		}

		coins, err = w.walletRepository.FindManyCoinByName(tx, currencies)
		if err != nil {
			return fmt.Errorf("failed to find coins: %w", err)

		}

		for _, c := range coins {
			_ = coinCache.Set(c.Name, c)
		}

		queryBuilder := tx.Table(_wallet.TableName())

		for i, v := range event {
			_coin, err = coinCache.Get(v.GetCurrency().String())
			if err != nil {
				continue
			}

			_user, err = userCache.Get(v.GetUserUUID())
			if err != nil {
				continue
			}

			_wallet, err = walletCache.Get(fmt.Sprintf("%d-%d", _user.ID, _coin.ID))
			if err != nil {
				if i == 0 {
					queryBuilder = queryBuilder.Where("user_id = ? AND coin_id = ?", _user.ID, _coin.ID)
				} else {
					queryBuilder = queryBuilder.Or("user_id = ? AND coin_id = ?", _user.ID, _coin.ID)
				}
				_ = walletCache.Set(fmt.Sprintf("%d-%d", _user.ID, _coin.ID), new(entity.Wallet))
			}
		}

		if err = queryBuilder.Find(&wallets).Error; err != nil {
			return fmt.Errorf("failed to find wallets: %w", err)
		}

		for _, w := range wallets {
			_ = walletCache.Set(fmt.Sprintf("%d-%d", w.UserID, w.CoinID), w)
		}

		for _, e := range event {
			_coin, err = coinCache.Get(e.GetCurrency().String())
			if err != nil {
				continue
			}

			_user, err = userCache.Get(e.GetUserUUID())
			if err != nil {
				continue
			}

			_wallet, err = walletCache.Get(fmt.Sprintf("%d-%d", _user.ID, _coin.ID))
			if err != nil {
				continue
			}

			_transfer = &entity.CoinTransfer{
				WalletID:     _wallet.ID,
				Amount:       decimal.NewFromFloat(e.GetDiff()),
				TransferType: entity.Trade,
			}

			coinTransfers = append(coinTransfers, _transfer)
		}

		return tx.Table(_transfer.TableName()).CreateInBatches(coinTransfers, 100).Error
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (w *walletService) FindBlockchainByName(name string) (result *entity.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindBlockchainByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindBlockchainByID(id uint) (result *entity.Blockchain, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindBlockchainByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, err
}

func (w *walletService) FindCoinByID(id uint) (result *entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindCoinByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinByName(name string) (result *entity.Coin, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindCoinByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByCoinID(coinID uint) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindBlockNumberByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindBlockNumberByID(id uint) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindBlockNumberByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateBlockNumber(coinID uint, fromBlockNumber, toBlockNumber decimal.Decimal) (result *entity.BlockNumber, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.UpdateBlockNumber(tx, coinID, fromBlockNumber, toBlockNumber)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByID(id uint) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindWalletByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindWalletByAddress(addr string) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindWalletByAddress(tx, addr)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindAllWallet(coinID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindAllWallet(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByCoinID(coinID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWalletByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWalletByUserID(userID uint) (result []*entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWalletByUserID(tx, userID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) GetContractAddressByCoinID(coinID uint) (result *entity.SmartContract, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.GetContractAddressByCoinID(tx, coinID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertWallet(wallet *entity.Wallet) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.InsertWallet(tx, wallet)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWallet(id uint, address string) (result *entity.Wallet, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.UpdateWallet(tx, id, address)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransfer(walletID uint, amount decimal.Decimal, transferType entity.TransferType) (result *entity.CoinTransfer, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.InsertCoinTransfer(tx, walletID, amount, transferType)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByID(id uint) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindCoinTransactionByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) FindCoinTransactionByTxHash(txHash string) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.FindCoinTransactionByTxHash(tx, txHash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByTransferID(transferID uint) (result []*entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanCoinTransactionByTransferID(tx, transferID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanCoinTransactionByCond(transferID uint, status entity.TransactionStatus) (result []*entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanCoinTransactionByCond(tx, transferID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) InsertCoinTransaction(transferID uint, txHash string, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.InsertCoinTransaction(tx, transferID, txHash, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionHash(id uint, hash string) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.UpdateCoinTransactionHash(tx, id, hash)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateCoinTransactionStatus(id uint, txStatus entity.TransactionStatus) (result *entity.CoinTransaction, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.UpdateCoinTransactionStatus(tx, id, txStatus)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByStatus(status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWithdrawalRequestByStatus(tx, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) ScanWithdrawalRequestByCond(coinID uint, status entity.WithdrawalStatus) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWithdrawalRequestByCond(tx, coinID, status)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) UpdateWithdrawalRequest(id uint, state entity.WithdrawalStatus) (result *entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.UpdateWithdrawalRequest(tx, id, state)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, err
}

func (w *walletService) Ctx() context.Context {
	return context.Background()
}

func (w *walletService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (w *walletService) Tx(level sql.IsolationLevel) *gorm.DB {
	return w.walletRepository.Tx(level)
}

func (w *walletService) Conn() *gorm.DB {
	return w.walletRepository.Conn()
}

func newWalletService(db *gorm.DB) *walletService {
	return &walletService{walletRepository: NewRepository(db), userRepository: user.NewRepository(db)}
}