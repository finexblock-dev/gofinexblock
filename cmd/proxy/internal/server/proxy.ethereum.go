package server

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/ethereum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EthereumProxy struct {
	transaction ethereum.TransactionClient
	wallet      ethereum.HDWalletClient

	ethereum.UnimplementedEthereumProxyServer
}

func NewEthereumProxy(conn *grpc.ClientConn) *EthereumProxy {
	return &EthereumProxy{
		transaction:                      ethereum.NewTransactionClient(conn),
		wallet:                           ethereum.NewHDWalletClient(conn),
		UnimplementedEthereumProxyServer: ethereum.UnimplementedEthereumProxyServer{},
	}
}

func (e *EthereumProxy) GetReceipt(_ context.Context, input *ethereum.GetReceiptInput) (*ethereum.GetReceiptOutput, error) {
	switch {
	case input.TxHash == "":
		return nil, status.Errorf(codes.InvalidArgument, "tx hash is empty")
	case len(input.TxHash) != 66:
		return nil, status.Errorf(codes.InvalidArgument, "tx hash is invalid")
	}
	receipt, err := e.transaction.GetReceipt(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get receipt: [%v]", err)
	}
	return receipt, nil
}

func (e *EthereumProxy) SendTransaction(_ context.Context, input *ethereum.SendRawTransactionInput) (*ethereum.SendRawTransactionOutput, error) {
	output, err := e.transaction.SendRawTransaction(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to send transaction: [%v]", err)
	}
	return output, nil
}

func (e *EthereumProxy) GetBlockNumber(_ context.Context, input *ethereum.GetBlockNumberInput) (*ethereum.GetBlockNumberOutput, error) {
	blockNumber, err := e.transaction.GetBlockNumber(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get block number: [%v]", err)
	}
	return blockNumber, nil
}

func (e *EthereumProxy) GetBlocks(_ context.Context, input *ethereum.GetBlocksInput) (*ethereum.GetBlocksOutput, error) {
	output, err := e.transaction.GetBlocks(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get blocks: %v", err)
	}
	return output, nil
}

func (e *EthereumProxy) CreateWallet(_ context.Context, input *ethereum.CreateWalletInput) (*ethereum.CreateWalletOutput, error) {
	output, err := e.wallet.CreateWallet(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to create wallet: %v", err)
	}
	return output, nil
}

func (e *EthereumProxy) GetBalance(_ context.Context, input *ethereum.GetBalanceInput) (*ethereum.GetBalanceOutput, error) {
	output, err := e.wallet.GetBalance(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get balance: [%v]", err)
	}
	return output, nil
}

func (e *EthereumProxy) mustEmbedUnimplementedEthereumProxyServer() {
	//TODO implement me
	panic("implement me")
}
