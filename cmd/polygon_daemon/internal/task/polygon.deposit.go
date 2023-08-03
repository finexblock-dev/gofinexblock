package polygon

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/constant"
	"github.com/finexblock-dev/gofinexblock/pkg/daemon"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/files"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	"github.com/finexblock-dev/gofinexblock/pkg/instance"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"log"
	"time"
)

type depositDaemon struct {
	polygon polygon.PolygonProxyClient
	// status of daemon
	status daemon.State
	// interval duration
	interval time.Duration
	// wallet Repository implementation
	walletRepo wallet.Repository
	// user Repository implementation
	userRepo user.Repository
	// trade Manager implementation
	tradeManager trade.Manager
	// finexblock Repository implementation
	instanceRepo instance.Repository
	// central wallet address
	centralWalletAddress string
	// instance
	instance *entity.FinexblockServer
	// file logger service implementation
	logger files.Writer
}

func (d *depositDaemon) initialize() {
	var _instance *entity.FinexblockServer
	var err error

	_instance, err = d.instanceRepo.FindServerByName(d.instanceRepo.Conn(), constant.PolygonDepositDaemon)
	if err != nil {
		panic(err)
	}

	d.instance = _instance
	d.status = daemon.Running
}

func (d *depositDaemon) Run() {
	d.initialize()
	d.status = daemon.Running
}

func (d *depositDaemon) Sleep() {
	time.Sleep(d.interval)
}

func (d *depositDaemon) Stop() {
	d.status = daemon.Stopped
}

func (d *depositDaemon) State() daemon.State {
	return d.status
}

func (d *depositDaemon) SetState(state daemon.State) {
	d.status = state
}

func (d *depositDaemon) Log(v ...any) {
	log.Println(v)
	d.logger.Println(v)
}

func (d *depositDaemon) InsertErrLog(err error) error {
	var _err error

	_, _err = d.instanceRepo.InsertErrorLog(d.instanceRepo.Conn(), &entity.FinexblockErrorLog{
		ID:          0,
		ServerID:    d.instance.ID,
		Process:     "Polygon Deposit Daemon",
		Priority:    entity.HIGH,
		Description: "Error on Polygon Deposit Daemon",
		Err:         err.Error(),
		Metadata:    map[string]interface{}{"timestamp": time.Now(), "utc": time.Now().UTC()},
	})
	return _err
}

func (d *depositDaemon) GetCurrentBlockNumber(ctx context.Context) (uint64, error) {
	var output *polygon.GetBlockNumberOutput
	var err error

	output, err = d.polygon.GetBlockNumber(ctx, &polygon.GetBlockNumberInput{})
	if err != nil {
		return 0, err
	}

	return output.BlockNumber, nil
}

func (d *depositDaemon) GetBlock(ctx context.Context, start, end decimal.Decimal) ([]*polygon.TxData, error) {
	var output *polygon.GetBlocksOutput
	var err error

	output, err = d.polygon.GetBlocks(ctx, &polygon.GetBlocksInput{Start: uint64(start.IntPart()), End: uint64(end.IntPart())})
	if err != nil {
		d.Log("error on get block", err)
		d.Log("insert error log", d.InsertErrLog(err))
		return nil, err
	}

	return output.GetResult(), nil
}

func (d *depositDaemon) GetBalance(ctx context.Context, address string) (*polygon.GetBalanceOutput, error) {
	return d.polygon.GetBalance(ctx, &polygon.GetBalanceInput{Address: address})
}

func (d *depositDaemon) IsValidTransaction(tx *gorm.DB, txHash string) (bool, error) {
	var err error

	_, err = d.walletRepo.FindCoinTransactionByTxHash(tx, txHash)
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	return false, fmt.Errorf("transaction already exists")
}

func (d *depositDaemon) ScanWallets(tx *gorm.DB) ([]*entity.Wallet, error) {
	return d.walletRepo.ScanWalletByCoinID(tx, constant.PolygonCoinID)
}

func (d *depositDaemon) InsertCoinTransaction(tx *gorm.DB, transferID uint, txHash string, txStatus entity.TransactionStatus) (*entity.CoinTransaction, error) {
	return d.walletRepo.InsertCoinTransaction(tx, transferID, txHash, txStatus)
}

func (d *depositDaemon) InsertCoinTransfer(tx *gorm.DB, walletID uint, amount decimal.Decimal, transferType entity.TransferType) (*entity.CoinTransfer, error) {
	return d.walletRepo.InsertCoinTransfer(tx, walletID, amount, transferType)
}

func (d *depositDaemon) AcquireLock(uuid, coin string) (bool, error) {
	for {
		if ok, err := d.tradeManager.AcquireLock(uuid, coin); !ok || err != nil {
			d.logger.Println("acquire lock failed, retrying...", "uuid: ", uuid, "coin: ", coin)
			continue
		}
		return true, nil
	}
}

func (d *depositDaemon) Transfer(ctx context.Context, from, to string, amount decimal.Decimal) (*polygon.SendRawTransactionOutput, error) {
	return d.polygon.SendRawTransaction(ctx, &polygon.SendRawTransactionInput{From: from, To: to, Amount: amount.String()})
}

func (d *depositDaemon) CheckReceipt(txHash string) (*polygon.GetReceiptOutput, error) {
	return d.polygon.GetReceipt(context.Background(), &polygon.GetReceiptInput{TxHash: txHash})
}

func (d *depositDaemon) ReleaseLock(uuid, coin string) error {
	return d.tradeManager.ReleaseLock(uuid, coin)
}

func (d *depositDaemon) Gathering(ctx context.Context) {
	var err error

	d.Log("gathering start")

	if err = d.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {
		var wallets []*entity.Wallet
		var _wallet *entity.Wallet
		var _transfer *polygon.SendRawTransactionOutput
		var _balance *polygon.GetBalanceOutput
		var _value, _transferValue decimal.Decimal
		var _err error

		if wallets, _err = d.ScanWallets(tx); _err != nil {
			return fmt.Errorf("error on scan wallets: %v", err)
		}

		for _, _wallet = range wallets {
			address := _wallet.Address

			if address == "" {
				continue
			}

			// get balance
			_balance, _err = d.GetBalance(ctx, address)
			if _err != nil {
				continue
			}

			_value, _err = decimal.NewFromString(_balance.Balance)
			if _err != nil {
				continue
			}

			_transferValue = _value.Sub(constant.PolygonMinimumGatheringAmount)

			if _transferValue.IsNegative() {
				continue
			}

			_transfer, _err = d.Transfer(ctx, address, d.centralWalletAddress, _transferValue)
			if _err != nil {
				d.Log("error on transfer", _err)
				continue
			}

			d.Log("Transfer from", address, "to", d.centralWalletAddress, "amount", _transferValue.String(), "txHash", _transfer.TxHash)
		}
		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
		d.Log("error on transaction", err)
		d.Log("insert error log", d.InsertErrLog(err))
	}

}

func (d *depositDaemon) Task() error {
	var blockNumber *entity.BlockNumber
	var from, to decimal.Decimal
	var currentBlockNumber uint64
	var err error

	return d.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {
		// Find last from-to block number
		blockNumber, err = d.walletRepo.FindBlockNumberByCoinID(tx, constant.PolygonCoinID)
		if err != nil {
			d.Log("error while find from-to blocknumber: [%v]", err)
			d.Log("insert error log", d.InsertErrLog(err))
			return fmt.Errorf("error while find from-to blocknumber: [%v]", err)
		}

		from, to = blockNumber.FromBlock, blockNumber.ToBlock

		currentBlockNumber, err = d.GetCurrentBlockNumber(context.Background())
		if err != nil {
			d.Log("error while get current block number: [%v]", err)
			d.Log("insert error log", d.InsertErrLog(err))
			return fmt.Errorf("error while get current block number: [%v]", err)
		}

		// If to block number is greater than current block number, set (to block number) to (current block number)
		if currentBlockNumber < uint64(to.IntPart()) {
			to = decimal.NewFromFloat(float64(currentBlockNumber))
		}

		// If from block number is greater than to block number, set (from block number) to (to block number)
		if from.GreaterThan(to) {
			from = to
		}

		// Update block number (do not care about this task is successfully end or not)
		if _, err = d.walletRepo.UpdateBlockNumber(tx, constant.PolygonCoinID, to.Add(decimal.NewFromFloat(1)), to.Add(decimal.NewFromFloat(10))); err != nil {
			d.Log("error while update block number: [%v]", err)
			d.Log("insert error log", d.InsertErrLog(err))
			return fmt.Errorf("error while update block number: [%v]", err)
		}

		return d.walletRepo.Conn().Transaction(func(tx2 *gorm.DB) error {
			var wallets []*entity.Wallet
			var blocks []*polygon.TxData

			var walletMap = make(map[string]*entity.Wallet)
			var addressMap = make(map[string]bool)

			var ctx = context.TODO()
			// Get block data
			blocks, err = d.GetBlock(ctx, from, to)
			if err != nil {
				d.Log(fmt.Errorf("error while get blocks: [%v]", err))
				d.Log("insert error log", d.InsertErrLog(err))
				return fmt.Errorf("error while get blocks: [%v]", err)
			}

			wallets, err = d.ScanWallets(tx2)
			if err != nil {
				d.Log(fmt.Errorf("error while scan wallets: [%v]", err))
				d.Log("insert error log", d.InsertErrLog(err))
				return fmt.Errorf("error while scan wallets: [%v]", err)
			}

			for _, w := range wallets {
				addressMap[w.Address] = true
				walletMap[w.Address] = w
			}

			for _, txData := range blocks {
				if err = tx2.Transaction(func(tx3 *gorm.DB) error {
					var receipt *polygon.GetReceiptOutput
					var amount decimal.Decimal

					var _wallet *entity.Wallet
					var _transfer *entity.CoinTransfer
					var _user *entity.User

					// Check if the to address is in our wallet
					if _, ok := addressMap[txData.ToAddress]; !ok {
						return nil
					}

					// Check if the transaction is valid
					if ok, err := d.IsValidTransaction(tx3, txData.TxHash); !ok || err != nil {
						return fmt.Errorf("error while check valid transaction: [%v]", err)
					}

					// Check receipt
					receipt, err = d.CheckReceipt(txData.TxHash)
					if err != nil {
						return fmt.Errorf("error while check receipt: [%v]", err)
					}

					// Check if the transaction is success
					if receipt.Status != 1 {
						return fmt.Errorf("failed transaction")
					}

					// Logging transaction
					d.Log("Tx Data.TxHash", txData.TxHash)
					d.Log("Tx Data.Amount", txData.Amount)
					d.Log("Tx Data.ToAddress", txData.ToAddress)

					_wallet = walletMap[txData.ToAddress]

					// Find User
					_user, err = d.userRepo.FindUserByID(tx3, walletMap[txData.ToAddress].UserID)
					if err != nil {
						return fmt.Errorf("error while find user: [%v]", err)
					}

					// Acquire lock
					if ok, err := d.AcquireLock(_user.UUID, constant.MATIC); !ok || err != nil {
						return fmt.Errorf("error while acquire lock: [%v]", err)
					}
					defer func(d *depositDaemon, uuid, coin string) {
						err = d.ReleaseLock(uuid, coin)
						if err != nil {
							d.Log(fmt.Errorf("error while release lock: [%v]", err))
							d.Log("insert error log", d.InsertErrLog(err))
						}
					}(d, _user.UUID, constant.MATIC)

					amount, err = decimal.NewFromString(txData.Amount)
					if err != nil {
						return fmt.Errorf("error while convert amount to decimal: [%v]", err)
					}

					// Plus balance
					if err = d.tradeManager.PlusBalance(_user.UUID, constant.MATIC, amount); err != nil {
						return fmt.Errorf("error while plus balance: [%v]", err)
					}

					// Insert transaction
					_transfer, err = d.InsertCoinTransfer(tx3, _wallet.ID, amount, entity.Deposit)
					if err != nil {
						return fmt.Errorf("error while insert coin transfer: [%v]", err)
					}

					_, err = d.InsertCoinTransaction(tx3, _transfer.ID, txData.TxHash, entity.DONE)
					if err != nil {
						return fmt.Errorf("error while insert coin transaction: [%v]", err)
					}

					return nil
				}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
					d.Log(fmt.Errorf("error while transaction: [%v]", err))
					d.Log("insert error log", d.InsertErrLog(err))
				}

			}
			return nil
		}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
}

func newDepositDaemon(
	proxy polygon.PolygonProxyClient,
	centralWalletAddress string,
	db *gorm.DB,
	client *redis.ClusterClient,
	interval time.Duration,
) *depositDaemon {
	return &depositDaemon{
		polygon:              proxy,
		status:               daemon.Running,
		interval:             interval,
		walletRepo:           wallet.NewRepository(db),
		userRepo:             user.NewRepository(db),
		tradeManager:         trade.New(client),
		instanceRepo:         instance.NewRepository(db),
		centralWalletAddress: centralWalletAddress,
		logger:               files.NewWriter("polygon-deposit-daemon", "polygon-deposit-daemon"),
	}
}