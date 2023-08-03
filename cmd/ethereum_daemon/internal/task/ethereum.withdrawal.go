package ethereum

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/constant"
	"github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/files"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/ethereum"
	"github.com/finexblock-dev/gofinexblock/finexblock/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/wallet"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"time"
)

type withdrawalDaemon struct {
	// ethereum proxy server connection
	ethereum ethereum.EthereumProxyClient
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
	// instance Repository implementation
	instanceRepo instance.Repository
	// instance
	instance *entity.FinexblockServer
	// logger service implementation
	logger files.Writer
	// central wallet address
	centralWalletAddress string
}

func (w *withdrawalDaemon) initialize() {
	var _instance *entity.FinexblockServer
	var err error

	_instance, err = w.instanceRepo.FindServerByName(w.instanceRepo.Conn(), constant.EthereumWithdrawalDaemon)
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

func (w *withdrawalDaemon) Sleep() {
	time.Sleep(w.interval)
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

func (w *withdrawalDaemon) Log(v ...any) {
	log.Println(v)
	w.logger.Println(v)
}

func (w *withdrawalDaemon) InsertErrLog(err error) error {
	var _err error

	_, _err = w.instanceRepo.InsertErrorLog(w.instanceRepo.Conn(), &entity.FinexblockErrorLog{
		ID:          0,
		ServerID:    w.instance.ID,
		Process:     constant.EthereumWithdrawalDaemon,
		Priority:    entity.HIGH,
		Description: "Error on Ethereum withdrawal daemon",
		Err:         err.Error(),
		Metadata:    map[string]interface{}{"timestamp": time.Now(), "utc": time.Now().UTC()},
	})
	return _err
}

func (w *withdrawalDaemon) ScanWithdrawalRequests(tx *gorm.DB, status entity.WithdrawalStatus) ([]*entity.WithdrawalRequest, error) {
	return w.walletRepo.ScanWithdrawalRequestByCond(tx, constant.EthereumCoinID, status)
}

func (w *withdrawalDaemon) UpdateWithdrawalRequest(tx *gorm.DB, id uint, status entity.WithdrawalStatus) error {
	var err error
	_, err = w.walletRepo.UpdateWithdrawalRequest(tx, id, status)
	return err
}

func (w *withdrawalDaemon) Transfer(ctx context.Context, from, to string, amount decimal.Decimal) (*ethereum.SendRawTransactionOutput, error) {
	return w.ethereum.SendRawTransaction(ctx, &ethereum.SendRawTransactionInput{
		From:   from,
		To:     to,
		Amount: amount.String(),
	})
}

func (w *withdrawalDaemon) GetReceipt(ctx context.Context, txHash string) (*ethereum.GetReceiptOutput, error) {
	return w.ethereum.GetReceipt(ctx, &ethereum.GetReceiptInput{TxHash: txHash})
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
		var output *ethereum.GetReceiptOutput
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

					output, err = w.GetReceipt(context.Background(), _transaction.TxHash)
					if err != nil {
						w.Log("coin_transaction.id = ", _transaction.ID, "error on get receipt:", err)
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on get receipt: %w", _transaction.ID, err)))
						return fmt.Errorf("coin_transaction.id = %d error on get receipt: %w", _transaction.ID, err)
					}

					if output.TxHash != _transaction.TxHash {
						w.Log("coin_transaction.id = ", _transaction.ID, "transaction hash mismatch:", output.TxHash, "!=", _transaction.TxHash)
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d transaction hash mismatch: %s != %s", _transaction.ID, output.TxHash, _transaction.TxHash)))
						return fmt.Errorf("coin_transaction.id = %d transaction hash mismatch: %s != %s", _transaction.ID, output.TxHash, _transaction.TxHash)
					}

					// If there are only failed transaction, we will update the withdrawal request to failed
					// If there are success transaction, we will update the withdrawal request to completed
					switch output.Status {
					// case 1: success transaction
					case 1:
						if err = w.UpdateWithdrawalRequest(tx2, request.ID, entity.COMPLETED); err != nil {
							w.Log("coin_transaction.id = ", _transaction.ID, "error on update withdrawal request:", err)
							w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update withdrawal request: %w", _transaction.ID, err)))
							return fmt.Errorf("coin_transaction.id = %d error on update withdrawal request: %w", _transaction.ID, err)
						}

						if _, err = w.UpdateCoinTransactionStatus(tx2, _transaction.ID, entity.DONE); err != nil {
							w.Log("coin_transaction.id = ", _transaction.ID, "error on update coin transaction status:", err)
							w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %w", _transaction.ID, err)))
							return fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %w", _transaction.ID, err)
						}
					// case 0: failed transaction
					case 0:
						if _, err = w.UpdateCoinTransactionStatus(tx2, _transaction.ID, entity.REVERT); err != nil {
							w.Log("coin_transaction.id = ", _transaction.ID, "error on update coin transaction status:", err)
							w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %w", _transaction.ID, err)))
							return fmt.Errorf("coin_transaction.id = %d error on update coin transaction status: %w", _transaction.ID, err)
						}
					// default is unknown status
					default:
						w.Log("coin_transaction.id = ", _transaction.ID, "unknown status:", output.Status)
						w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d unknown status: %d", _transaction.ID, output.Status)))
					}
					return nil
				}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
					w.Log("coin_transaction.id = ", _transaction.ID, "error on transaction:", err)
					w.Log("insert error log:", w.InsertErrLog(fmt.Errorf("coin_transaction.id = %d error on transaction: %w", _transaction.ID, err)))
				}
			}
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
}

func (w *withdrawalDaemon) ProcessApprovedWithdrawalRequest() error {
	return w.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {
		var requests []*entity.WithdrawalRequest
		var output *ethereum.SendRawTransactionOutput
		var coinTransaction *entity.CoinTransaction
		var err error

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

				output, err = w.Transfer(context.Background(), w.centralWalletAddress, request.ToAddress, request.Amount)
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

func newWithdrawalDaemon(ethereum ethereum.EthereumProxyClient, centralWalletAddress string, db *gorm.DB, client *redis.ClusterClient, interval time.Duration) *withdrawalDaemon {
	return &withdrawalDaemon{
		ethereum:             ethereum,
		status:               daemon.Running,
		interval:             interval,
		walletRepo:           wallet.NewRepository(db),
		userService:          user.NewService(db, client),
		tradeManager:         trade.New(client),
		instanceRepo:         instance.NewRepository(db),
		logger:               files.NewWriter("ethereum-withdrawal-daemon", "ethereum-withdrawal-daemon"),
		centralWalletAddress: centralWalletAddress,
	}
}
