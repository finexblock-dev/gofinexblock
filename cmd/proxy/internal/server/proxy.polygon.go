package server

import (
	context "context"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PolygonProxy struct {
	transaction polygon.TransactionClient
	wallet      polygon.HDWalletClient

	polygon.UnimplementedPolygonProxyServer
}

func (e *PolygonProxy) GetReceipt(_ context.Context, input *polygon.GetReceiptInput) (*polygon.GetReceiptOutput, error) {
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

func (e *PolygonProxy) SendTransaction(_ context.Context, input *polygon.SendRawTransactionInput) (*polygon.SendRawTransactionOutput, error) {
	output, err := e.transaction.SendRawTransaction(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to send transaction: [%v]", err)
	}
	return output, nil
}

func (e *PolygonProxy) GetBlockNumber(_ context.Context, input *polygon.GetBlockNumberInput) (*polygon.GetBlockNumberOutput, error) {
	blockNumber, err := e.transaction.GetBlockNumber(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get block number: [%v]", err)
	}
	return blockNumber, nil
}

func (e *PolygonProxy) GetBlocks(_ context.Context, input *polygon.GetBlocksInput) (*polygon.GetBlocksOutput, error) {
	output, err := e.transaction.GetBlocks(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get blocks: %v", err)
	}
	return output, nil
}

func (e *PolygonProxy) CreateWallet(_ context.Context, input *polygon.CreateWalletInput) (*polygon.CreateWalletOutput, error) {
	output, err := e.wallet.CreateWallet(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to create wallet: %v", err)
	}
	return output, nil
}

func (e *PolygonProxy) GetBalance(_ context.Context, input *polygon.GetBalanceInput) (*polygon.GetBalanceOutput, error) {
	output, err := e.wallet.GetBalance(context.Background(), input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "proxy: failed to get balance: [%v]", err)
	}
	return output, nil
}

func (e *PolygonProxy) mustEmbedUnimplementedPolygonProxyServer() {
	//TODO implement me
	panic("implement me")
}

func NewPolygonProxy(conn *grpc.ClientConn) *PolygonProxy {
	return &PolygonProxy{
		transaction:                     polygon.NewTransactionClient(conn),
		wallet:                          polygon.NewHDWalletClient(conn),
		UnimplementedPolygonProxyServer: polygon.UnimplementedPolygonProxyServer{},
	}
}