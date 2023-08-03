package bitcoin

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/constant"
	"github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/files"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/bitcoin"
	"github.com/finexblock-dev/gofinexblock/finexblock/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/wallet"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type depositDaemon struct {
	// bitcoin server connection
	bitcoin bitcoin.BitcoinProxyClient
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
	// list of utxo to gather central wallet
	list []*bitcoin.UnspentOutput
}

func (d *depositDaemon) ScanWallets(tx *gorm.DB) ([]*entity.Wallet, error) {
	return d.walletRepo.ScanWalletByCoinID(tx, constant.BitcoinCoinID)
}

func (d *depositDaemon) ListUnspentUTXO(ctx context.Context, address string) (*bitcoin.ListUnspentOutput, error) {
	return d.bitcoin.ListUnspent(ctx, &bitcoin.ListUnspentInput{
		MinConf: 0,
		MaxConf: 9999999,
		Address: address,
	})
}

func (d *depositDaemon) IsValidTransaction(tx *gorm.DB, txHash string) (bool, error) {
	var err error

	_, err = d.walletRepo.FindCoinTransactionByTxHash(tx, txHash)
	if err == gorm.ErrRecordNotFound {
		return true, nil
	}

	return false, err
}

func (d *depositDaemon) initialize() {
	var _instance *entity.FinexblockServer
	var err error

	_instance, err = d.instanceRepo.FindServerByName(d.instanceRepo.Conn(), constant.BitcoinDepositDaemon)
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

func (d *depositDaemon) InsertErrLog(err error) error {
	var _err error
	_, _err = d.instanceRepo.InsertErrorLog(d.instanceRepo.Conn(), &entity.FinexblockErrorLog{
		ServerID:    d.instance.ID,
		Process:     "Bitcoin Deposit Daemon",
		Priority:    "HIGH",
		Description: "Error on Bitcoin Deposit Daemon",
		Err:         err.Error(),
		Metadata:    map[string]interface{}{"timestamp": time.Now(), "utc": time.Now().UTC()},
	})
	return _err
}

func (d *depositDaemon) Log(v ...any) {
	d.logger.Println(v)
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

func (d *depositDaemon) ReleaseLock(uuid, coin string) error {
	return d.tradeManager.ReleaseLock(uuid, coin)
}

func (d *depositDaemon) Gathering(ctx context.Context) {
	// If list is empty, return
	if len(d.list) == 0 {
		return
	}

	// New list for remain utxo
	var remainList []*bitcoin.UnspentOutput

	// Send to hot wallet
	for _, utxo := range d.list {

		_, err := d.bitcoin.SendToAddress(ctx, &bitcoin.SendToAddressInput{
			ToAddress: d.centralWalletAddress,
			Amount:    utxo.Amount,
		})
		if err != nil {
			d.Log("error on send to address", err)
			d.Log(d.InsertErrLog(err))
			remainList = append(remainList, utxo)
		}
	}

	d.list = remainList
}

func (d *depositDaemon) Task() error {
	return d.walletRepo.Conn().Transaction(func(tx *gorm.DB) error {
		var unspent *bitcoin.ListUnspentOutput
		var wallets []*entity.Wallet
		var _wallet *entity.Wallet
		var _user *entity.User
		var _coinTransfer *entity.CoinTransfer
		var amount decimal.Decimal
		var ctx context.Context
		var err error

		ctx = context.TODO()
		wallets, err = d.ScanWallets(tx)
		if err != nil {
			d.Log("error on scan wallets", err)
			return d.InsertErrLog(err)
		}

		for _, _wallet = range wallets {

			// If address is empty, continue
			if _wallet.Address == "" {
				continue
			}

			// Get user's unspent UTXO list
			unspent, err = d.ListUnspentUTXO(ctx, _wallet.Address)
			if err != nil {
				d.Log("error on list unspent utxo", err)
				return d.InsertErrLog(err)
			}

			// If no unspent UTXO, continue
			if len(unspent.UnspentOutputs) == 0 {
				continue
			}

			for _, utxo := range unspent.UnspentOutputs {
				if utxo.Amount == 0 {
					continue
				}

				if ok, err := d.IsValidTransaction(tx, utxo.GetTxHash()); err != nil || !ok {
					continue
				}

				_user, err = d.userRepo.FindUserByID(tx, _wallet.UserID)
				if err != nil {
					d.Log("error on find user by id", err)
					return d.InsertErrLog(err)
				}

				if ok, err := d.AcquireLock(_user.UUID, "BTC"); err != nil || !ok {
					d.Log("error on acquire lock", err)
					return d.InsertErrLog(err)
				}

				amount = decimal.NewFromFloat(utxo.GetAmount())
				_coinTransfer, err = d.InsertCoinTransfer(tx, _wallet.ID, amount, entity.Deposit)
				if err != nil {
					d.Log("error on insert coin transfer", err)
					return d.InsertErrLog(err)
				}

				_, err = d.InsertCoinTransaction(tx, _coinTransfer.ID, utxo.GetTxHash(), entity.DONE)
				if err != nil {
					d.Log("error on insert coin transaction", err)
					return d.InsertErrLog(err)
				}

				if err = d.tradeManager.PlusBalance(_user.UUID, "BTC", amount); err != nil {
					d.Log("error on plus balance", err)
					return d.InsertErrLog(err)
				}

				if err = d.ReleaseLock(_user.UUID, "BTC"); err != nil {
					d.Log("error on release lock", err)
					return d.InsertErrLog(err)
				}

				d.list = append(d.list, utxo)
			}
		}

		// If there is unspent UTXO, gather it
		d.Gathering(ctx)

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
}

func newDepositDaemon(proxy bitcoin.BitcoinProxyClient, cluster *redis.ClusterClient, db *gorm.DB, interval time.Duration, address string) *depositDaemon {
	return &depositDaemon{
		bitcoin:              proxy,
		status:               daemon.Running,
		interval:             interval,
		walletRepo:           wallet.NewRepository(db),
		userRepo:             user.NewRepository(db),
		tradeManager:         trade.New(cluster),
		instanceRepo:         instance.NewRepository(db),
		centralWalletAddress: address,
		logger:               files.NewWriter("bitcoin-deposit-daemon", "bitcoin-deposit-daemon"),
		list:                 []*bitcoin.UnspentOutput{},
	}
}
