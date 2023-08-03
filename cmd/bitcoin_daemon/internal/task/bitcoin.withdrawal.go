package bitcoin

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/constant"
	"github.com/finexblock-dev/gofinexblock/pkg/daemon"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/files"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	"github.com/finexblock-dev/gofinexblock/pkg/instance"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type withdrawalDaemon struct {
	// bitcoin server connection
	bitcoin bitcoin.BitcoinProxyClient
	// status of daemon
	status daemon.State
	// interval duration
	interval time.Duration
	// wallet Repository implementation
	walletRepo wallet.Repository
	// user service implementation
	userService user.Service
	// trade Manager implementation
	tradeManager trade.Manager
	// finexblock Repository implementation
	instanceRepo instance.Repository
	// instance
	instance *entity.FinexblockServer
	// file logger service implementation
	logger files.Writer
}

func (w *withdrawalDaemon) initialize() {
	var _instance *entity.FinexblockServer
	var err error

	_instance, err = w.instanceRepo.FindServerByName(w.instanceRepo.Conn(), constant.BitcoinWithdrawalDaemon)
	if err != nil {
		panic(err)
	}

	w.instance = _instance
	w.status = daemon.Running
}

func (w *withdrawalDaemon) Run() {
	w.initialize()
	w.status = daemon.Running
}

func (w *withdrawalDaemon) Stop() {
	w.status = daemon.Stopped
}

func (w *withdrawalDaemon) State() daemon.State {
	return w.status
}

func (w *withdrawalDaemon) SetState(state daemon.State) {
	w.status = state
}

func (w *withdrawalDaemon) Sleep() {
	time.Sleep(w.interval)
}

func (w *withdrawalDaemon) InsertErrLog(err error) error {
	var _err error
	_, _err = w.instanceRepo.InsertErrorLog(w.instanceRepo.Conn(), &entity.FinexblockErrorLog{
		ServerID:    w.instance.ID,
		Process:     "BitcoinWithdrawalDaemon",
		Priority:    entity.HIGH,
		Description: "Error on Bitcoin Withdrawal Daemon",
		Err:         err.Error(),
		Metadata:    map[string]interface{}{"timestamp": time.Now(), "utc": time.Now().UTC()},
	})
	return _err
}

func (w *withdrawalDaemon) Log(v ...any) {
	w.logger.Println(v)
}

func (w *withdrawalDaemon) ScanWithdrawalRequests(tx *gorm.DB, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error) {
	return w.walletRepo.ScanWithdrawalRequestByCond(tx, constant.BitcoinCoinID, status)
}

func (w *withdrawalDaemon) UpdateWithdrawalRequest(tx *gorm.DB, id uint, status entity.WithdrawalStatus) error {
	var err error
	_, err = w.walletRepo.UpdateWithdrawalRequest(tx, id, status)
	return err
}

func (w *withdrawalDaemon) Withdrawal(ctx context.Context, toAddress string, amount decimal.Decimal) (*bitcoin.SendToAddressOutput, error) {
	return w.bitcoin.SendToAddress(ctx, &bitcoin.SendToAddressInput{ToAddress: toAddress, Amount: amount.InexactFloat64()})
}

func (w *withdrawalDaemon) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error) {
	return w.walletRepo.InsertCoinTransaction(tx, transferID, txHash, txStatus)
}

func (w *withdrawalDaemon) InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error) {
	return w.walletRepo.InsertCoinTransfer(tx, walletID, amount, transferType)
}

func (w *withdrawalDaemon) UpdateCoinTransactionHash(tx *gorm.DB, id uint, txHash string) (*entity.CoinTransaction, error) {
	return w.walletRepo.UpdateCoinTransactionHash(tx, id, txHash)
}

func (w *withdrawalDaemon) UpdateCoinTransactionStatus(tx *gorm.DB, id uint, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error) {
	return w.walletRepo.UpdateCoinTransactionStatus(tx, id, txStatus)
}

func (w *withdrawalDaemon) ProcessPendingWithdrawalRequest() error {
	return w.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {
		var requests []*entity.WithdrawalRequest
		var output *bitcoin.GetRawTransactionOutput
		var transactions []*entity.CoinTransaction
		var _transaction *entity.CoinTransaction
		var err error

		requests, err = w.ScanWithdrawalRequests(tx, entity.PENDING)
		if err != nil {
			return fmt.Errorf("error on scan withdrawal requests: %w", err)
		}

		for _, request := range requests {
			transactions, err = w.walletRepo.ScanCoinTransactionByCond(tx, request.CoinTransferID, entity.INIT)
			if err != nil {
				return fmt.Errorf("error on scan coin transaction: %w", err)
			}

			// no transaction found
			if len(transactions) == 0 {
				// update withdrawal request status to failed
				if err = w.UpdateWithdrawalRequest(tx, request.ID, entity.FAILED); err != nil {
					w.Log("error on update withdrawal request:", err)
					w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("error on update withdrawal request: %w", err)))
					return fmt.Errorf("error on update withdrawal request: %w", err)
				}

				return nil
			}

			for _, _transaction = range transactions {

				if err = tx.Transaction(func(tx2 *gorm.DB) error {
					// invalid transaction hash
					if _transaction.TxHash == "" {
						w.Log("coin_transaction.id = ", _transaction.ID, "has invalid transaction hash", _transaction.TxHash)
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d has invalid transaction hash %s", _transaction.ID, _transaction.TxHash)))
						return fmt.Errorf("coin_transaction.id = %d has invalid transaction hash %s", _transaction.ID, _transaction.TxHash)
					}

					output, err = w.bitcoin.GetRawTransaction(context.Background(), &bitcoin.GetRawTransactionInput{TxId: _transaction.TxHash})
					if err != nil {
						w.Log("coin_transaction.id = ", _transaction.ID, "error on get raw transaction:", err.Error())
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on get raw transaction: %s", _transaction.ID, err.Error())))
						return fmt.Errorf("coin_transaction.id = %d error on get raw transaction: %s", _transaction.ID, err.Error())
					}

					// Less than minimum confirmations, continue
					if output.GetConfirmations() < constant.BitcoinMinConfirmations {
						if _, err = w.UpdateCoinTransactionStatus(tx2, _transaction.ID, entity.REVERT); err != nil {
							w.Log("coin_transaction.id = ", _transaction.ID, "error on update coin transaction status:", err.Error())
							w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %s", _transaction.ID, err.Error())))
							return fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %s", _transaction.ID, err.Error())
						}
						return nil
					}

					if output.TxId != _transaction.TxHash {
						w.Log("coin_transaction.id = ", _transaction.ID, "transaction hash mismatch:", output.TxId, "!=", _transaction.TxHash)
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d transaction hash mismatch: %s != %s", _transaction.ID, output.TxId, _transaction.TxHash)))
						return fmt.Errorf("coin_transaction.id = %d transaction hash mismatch: %s != %s", _transaction.ID, output.TxId, _transaction.TxHash)
					}

					// Else, update withdrawal request status to completed
					if err = w.UpdateWithdrawalRequest(tx2, request.ID, entity.COMPLETED); err != nil {
						w.Log("coin_transaction.id = ", _transaction.ID, "error on update withdrawal request:", err.Error())
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update withdrawal request: %s", _transaction.ID, err.Error())))
						return fmt.Errorf("coin_transaction.id = %d error on update withdrawal request: %s", _transaction.ID, err.Error())
					}

					if _, err = w.UpdateCoinTransactionStatus(tx2, _transaction.ID, entity.DONE); err != nil {
						w.Log("coin_transaction.id = ", _transaction.ID, "error on update coin transaction:", err.Error())
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update coin transaction: %s", _transaction.ID, err.Error())))
						return fmt.Errorf("coin_transaction.id = %d error on update coin transaction: %s", _transaction.ID, err.Error())
					}

					return nil
				}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
					w.Log("coin_transaction.id = ", _transaction.ID, "error on transaction:", err.Error())
					w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on transaction: %s", _transaction.ID, err.Error())))
				}
			}
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
}

func (w *withdrawalDaemon) ProcessApprovedWithdrawalRequest() error {
	var requests []*entity.WithdrawalRequest
	var output *bitcoin.SendToAddressOutput
	var coinTransaction *entity.CoinTransaction
	var err error

	return w.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {

		requests, err = w.ScanWithdrawalRequests(tx, entity.APPROVED)
		if err != nil {
			return fmt.Errorf("error on scan withdrawal requests: %w", err)
		}

		for _, request := range requests {

			if err = tx.Transaction(func(tx2 *gorm.DB) error {
				coinTransaction, err = w.InsertCoinTransaction(tx2, request.CoinTransferID, "", entity.INIT)
				if err != nil {
					return fmt.Errorf("error on insert coin transaction: %w", err)
				}

				if err = w.UpdateWithdrawalRequest(tx2, request.ID, entity.PENDING); err != nil {
					return fmt.Errorf("error on update withdrawal request: %w", err)
				}

				output, err = w.Withdrawal(context.Background(), request.ToAddress, request.Amount)
				if err != nil {
					return fmt.Errorf("error on withdrawal: %w", err)
				}

				if _, err = w.UpdateCoinTransactionHash(tx2, coinTransaction.ID, output.TxHash); err != nil {
					return fmt.Errorf("error on update coin transaction: %w", err)
				}

				return nil
			}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
				w.Log("error on transaction:", err, "withdrawal_request.id = ", request.ID)
				w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("error on transaction: %w", err)))
			}
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
}

func (w *withdrawalDaemon) Task() error {
	var err error
	if err = w.ProcessApprovedWithdrawalRequest(); err != nil {
		return err
	}

	if err = w.ProcessPendingWithdrawalRequest(); err != nil {
		return err
	}

	return nil
}

func newWithdrawalDaemon(proxy bitcoin.BitcoinProxyClient, db *gorm.DB, client *redis.ClusterClient, interval time.Duration) *withdrawalDaemon {
	return &withdrawalDaemon{
		bitcoin:      proxy,
		status:       daemon.Running,
		interval:     interval,
		walletRepo:   wallet.NewRepository(db),
		userService:  user.NewService(db, client),
		tradeManager: trade.New(client),
		instanceRepo: instance.NewRepository(db),
		logger:       files.NewWriter("bitcoin-withdrawal-daemon", "bitcoin-withdrawal-daemon"),
	}
}