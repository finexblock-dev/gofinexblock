package wallet

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/cache"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet/structs"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type walletService struct {
	manager          trade.Manager
	walletRepository Repository
	userRepository   user.Repository
}

func (w *walletService) ScanWithdrawalRequestByCondWithLimitOffset(coinID uint, status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWithdrawalRequestByCondWithLimitOffset(tx, coinID, status, limit, offset)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (w *walletService) ScanWithdrawalRequestByStatusWithLimitOffset(status entity.WithdrawalStatus, limit, offset int) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWithdrawalRequestByStatusWithLimitOffset(tx, status, limit, offset)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (w *walletService) ScanWithdrawalRequestByUser(userID, coinID uint, limit, offset int) (result []*entity.WithdrawalRequest, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanWithdrawalRequestByUser(tx, userID, coinID, limit, offset)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (w *walletService) ScanCoinTransferByUser(userID uint, limit, offset int) (result []*entity.CoinTransfer, err error) {
	if err = w.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = w.walletRepository.ScanCoinTransferByUserID(tx, userID, limit, offset)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
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

		// 유저 캐시 용도 map
		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		// 코인 캐시 용도 map
		var coinCache = cache.NewDefaultKeyValueStore[entity.Coin](len(event))
		// 지갑 캐시 용도 map
		var walletCache = cache.NewDefaultKeyValueStore[entity.Wallet](len(event))

		// 일단 이벤트 순회
		for _, e := range event {
			// user uuid 담기 => 유저 찾는 용도 (in query)
			userUUIDs = append(userUUIDs, e.GetUserUUID())
			// currency 담기 => 코인 찾는 용도 (in query)
			currencies = append(currencies, e.GetCurrency().String())
		}

		// SELECT * FROM user u where u.uuid in (userUUIDs)
		users, err = w.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return fmt.Errorf("failed to find users: %w", err)
		}

		// 유저 캐싱
		for _, u := range users {
			_ = userCache.Set(u.UUID, u)
		}

		// SELECT * FROM coin c where c.name in (currencies)
		coins, err = w.walletRepository.FindManyCoinByName(tx, currencies)
		if err != nil {
			return fmt.Errorf("failed to find coins: %w", err)

		}

		// 코인 캐싱
		for _, c := range coins {
			_ = coinCache.Set(c.Name, c)
		}

		// SELECT FROM wallet w
		queryBuilder := tx.Table(_wallet.TableName())

		// 다시 이벤트 순회
		for i, v := range event {
			// checkpoint: 캐시에서 가져오지 못했다면 무조건 continue
			// 코인 캐시에서 코인 가져오기
			_coin, err = coinCache.Get(v.GetCurrency().String())
			if err != nil {
				continue
			}

			// 유저 캐시에서 유저 가져오기
			_user, err = userCache.Get(v.GetUserUUID())
			if err != nil {
				continue
			}

			// 첫 번째인 경우 where 절에 AND로 붙이고, 그 외에는 OR로 붙이기
			if i == 0 {
				queryBuilder = queryBuilder.Where("user_id = ? AND coin_id = ?", _user.ID, _coin.ID)
			} else {
				queryBuilder = queryBuilder.Or("user_id = ? AND coin_id = ?", _user.ID, _coin.ID)
			}
		}

		// 쿼리 실행
		if err = queryBuilder.Find(&wallets).Error; err != nil {
			return fmt.Errorf("failed to find wallets: %w", err)
		}

		// 지갑 캐싱
		for _, w := range wallets {
			_ = walletCache.Set(fmt.Sprintf("%d-%d", w.UserID, w.CoinID), w)
		}

		// 이벤트 다시 한 번 순회
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
				TransferType: entity.TransferType(e.GetReason().String()),
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

func newWalletService(db *gorm.DB, cluster *redis.ClusterClient) *walletService {
	return &walletService{walletRepository: NewRepository(db), userRepository: user.NewRepository(db), manager: trade.New(cluster)}
}