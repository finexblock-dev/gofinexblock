package server

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_key/internal/config"
	"github.com/finexblock-dev/gofinexblock/finexblock/btcd"
	"github.com/finexblock-dev/gofinexblock/finexblock/constant"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/bitcoin"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
)

type BitcoinKey struct {
	btcClient   *rpcclient.Client
	key         string
	user        string
	pass        string
	account     string
	walletType  string
	chainParams *chaincfg.Params

	bitcoin.UnimplementedHDWalletServer
	bitcoin.UnimplementedTransactionServer
}

func (b *BitcoinKey) Listen(gRPCServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}
	log.Println("GRPC SERVER START")
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener... : %v", err)
	}
}

func (b *BitcoinKey) Register(grpcServer *grpc.Server) {
	bitcoin.RegisterHDWalletServer(grpcServer, b)
	bitcoin.RegisterTransactionServer(grpcServer, b)
}

func NewBitcoinKey(configuration *config.BitcoinKeyConfiguration) (*BitcoinKey, error) {
	var params *chaincfg.Params
	switch os.Getenv("APPMODE") {
	case "PROD":
		params = &chaincfg.MainNetParams
	default:
		params = &chaincfg.TestNet3Params
	}

	cfg := &rpcclient.ConnConfig{
		Host:         fmt.Sprintf("%v:%v", configuration.RpcHost, configuration.RpcPort),
		User:         configuration.RpcUser,
		Pass:         configuration.RpcPassword,
		DisableTLS:   true,
		HTTPPostMode: true,
		Params:       params.Name,
	}
	btcClient := btcd.CreateClient(cfg)

	return &BitcoinKey{
		btcClient:                      btcClient,
		key:                            configuration.Mnemonic,
		account:                        configuration.WalletAccount,
		walletType:                     configuration.WalletType,
		chainParams:                    params,
		UnimplementedHDWalletServer:    bitcoin.UnimplementedHDWalletServer{},
		UnimplementedTransactionServer: bitcoin.UnimplementedTransactionServer{},
	}, nil
}

func (b *BitcoinKey) GetNewAddress(_ context.Context, _ *bitcoin.GetNewAddressInput) (*bitcoin.GetNewAddressOutput, error) {
	address, err := b.btcClient.GetNewAddress(b.account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating new wallet: %v, wallet type : %v", err, b.walletType)
	}
	return &bitcoin.GetNewAddressOutput{Address: address.EncodeAddress()}, nil
}

func (b *BitcoinKey) ListUnspent(_ context.Context, request *bitcoin.ListUnspentInput) (*bitcoin.ListUnspentOutput, error) {
	var future rpcclient.FutureListUnspentResult
	var unspent []btcjson.ListUnspentResult
	var addresses []btcutil.Address
	var addr btcutil.Address
	var txIdHash *chainhash.Hash
	var err error

	switch {
	case request.GetAddress() == "":
		return nil, status.Errorf(codes.InvalidArgument, "address is required")
	case request.GetMinConf() > request.GetMaxConf():
		return nil, status.Errorf(codes.InvalidArgument, "minConf must be less than maxConf")
	case request.GetMaxConf() == 0:
		return nil, status.Errorf(codes.InvalidArgument, "maxConf must be greater than 0")
	}

	addr, err = btcutil.DecodeAddress(request.GetAddress(), b.chainParams)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error decode address: %v", err)
	}

	addresses = append(addresses, addr)
	future = b.btcClient.ListUnspentMinMaxAddressesAsync(int(request.GetMinConf()), int(request.GetMaxConf()), addresses)

	response := &bitcoin.ListUnspentOutput{
		UnspentOutputs: []*bitcoin.UnspentOutput{},
	}
	response = new(bitcoin.ListUnspentOutput)
	response.UnspentOutputs = make([]*bitcoin.UnspentOutput, 0)

	unspent, err = future.Receive()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get unspent list on wallet: %v", err)
	}

	for _, output := range unspent {

		txIdHash, err = chainhash.NewHashFromStr(output.TxID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, " error get master key while get public key hash: %v", err)
		}

		response.UnspentOutputs = append(response.UnspentOutputs, &bitcoin.UnspentOutput{
			TxHash:        txIdHash.String(),
			Address:       output.Address,
			Account:       output.Account,
			Amount:        output.Amount,
			Confirmations: output.Confirmations,
			Spendable:     output.Spendable,
		})
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, "error get unspent list on wallet: %v", err)
	}
	return response, nil
}

func (b *BitcoinKey) SendToAddress(_ context.Context, request *bitcoin.SendToAddressInput) (*bitcoin.SendToAddressOutput, error) {
	addr, err := btcutil.DecodeAddress(request.GetToAddress(), b.chainParams)
	if err != nil {
		return nil, fmt.Errorf("error decode address while send to address: %v", err)
	}

	requestAmount := decimal.NewFromFloat(request.Amount)
	value := requestAmount.Mul(constant.BitcoinDecimal)
	amount := btcutil.Amount(value.IntPart())

	txHash, err := b.btcClient.SendToAddress(addr, amount)
	if err != nil {
		return nil, fmt.Errorf("error send to address: %v", err)
	}
	return &bitcoin.SendToAddressOutput{TxHash: txHash.String()}, nil
}

func (b *BitcoinKey) GetRawTransaction(_ context.Context, request *bitcoin.GetRawTransactionInput) (*bitcoin.GetRawTransactionOutput, error) {
	txHash, err := chainhash.NewHashFromStr(request.GetTxId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction ID: %v", err)
	}
	txRawResult, err := b.btcClient.GetRawTransactionVerbose(txHash)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error while getting raw transaction: %v", err)
	}
	vinList := make([]*bitcoin.Vin, len(txRawResult.Vin))
	for _, vin := range txRawResult.Vin {
		vinList = append(vinList, &bitcoin.Vin{
			Coinbase:     vin.Coinbase,
			TxId:         vin.Txid,
			Vout:         vin.Vout,
			ScriptSigAms: vin.ScriptSig.Asm,
			ScriptSigHex: vin.ScriptSig.Hex,
			Sequence:     vin.Sequence,
			Witness:      vin.Witness,
		})
	}
	voutList := make([]*bitcoin.Vout, len(txRawResult.Vout))
	for _, vout := range txRawResult.Vout {
		voutList = append(voutList, &bitcoin.Vout{
			Value:               vout.Value,
			N:                   vout.N,
			ScriptPubKeyAms:     vout.ScriptPubKey.Asm,
			ScriptPubKeyHex:     vout.ScriptPubKey.Hex,
			ScriptPubKeyReqSigs: vout.ScriptPubKey.ReqSigs,
			ScriptPubKeyType:    vout.ScriptPubKey.Type,
			Address:             vout.ScriptPubKey.Addresses,
		})
	}
	return &bitcoin.GetRawTransactionOutput{
		Hex:           txRawResult.Hex,
		TxId:          txRawResult.Txid,
		Hash:          txRawResult.Hash,
		Size:          txRawResult.Size,
		Version:       txRawResult.Version,
		Vsize:         txRawResult.Vsize,
		Weight:        txRawResult.Weight,
		LockTime:      txRawResult.LockTime,
		Vin:           vinList,
		Vout:          voutList,
		BlockHash:     txRawResult.BlockHash,
		Confirmations: txRawResult.Confirmations,
		Time:          txRawResult.Time,
		BlockTime:     txRawResult.Blocktime,
	}, nil
}

func (b *BitcoinKey) mustEmbedUnimplementedBlockchainServer() {
	//TODO implement me
	panic("implement me")
}

func (b *BitcoinKey) mustEmbedUnimplementedHDWalletServer() {
	//TODO implement me
	panic("implement me")
}

func (b *BitcoinKey) mustEmbedUnimplementedTransactionServer() {
	//TODO implement me
	panic("implement me")
}
