package server

import (
	"context"
	"errors"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	"google.golang.org/grpc"
)

type BitcoinProxy struct {
	wallet      bitcoin.HDWalletClient
	transaction bitcoin.TransactionClient

	bitcoin.UnimplementedBitcoinProxyServer
}

func NewBitcoinProxy(conn *grpc.ClientConn) *BitcoinProxy {

	return &BitcoinProxy{
		wallet:                          bitcoin.NewHDWalletClient(conn),
		transaction:                     bitcoin.NewTransactionClient(conn),
		UnimplementedBitcoinProxyServer: bitcoin.UnimplementedBitcoinProxyServer{},
	}
}

func (b *BitcoinProxy) GetRawTransaction(ctx context.Context, request *bitcoin.GetRawTransactionInput) (*bitcoin.GetRawTransactionOutput, error) {
	result, err := b.transaction.GetRawTransaction(ctx, request)
	if err != nil {
		return nil, errors.Join(errors.New("proxy server got error"), err)
	}
	return result, nil
}

func (b *BitcoinProxy) ListUnspent(ctx context.Context, request *bitcoin.ListUnspentInput) (*bitcoin.ListUnspentOutput, error) {
	result, err := b.wallet.ListUnspent(ctx, request)
	if err != nil {
		return nil, errors.Join(errors.New("proxy server got error"), err)
	}
	return result, nil
}

func (b *BitcoinProxy) SendToAddress(ctx context.Context, request *bitcoin.SendToAddressInput) (*bitcoin.SendToAddressOutput, error) {
	result, err := b.wallet.SendToAddress(ctx, request)
	if err != nil {
		return nil, errors.Join(errors.New("proxy server got error"), err)
	}
	return result, nil
}

func (b *BitcoinProxy) GetNewAddress(ctx context.Context, request *bitcoin.GetNewAddressInput) (*bitcoin.GetNewAddressOutput, error) {
	result, err := b.wallet.GetNewAddress(ctx, request)
	if err != nil {
		return nil, errors.Join(errors.New("proxy server got error"), err)
	}
	return result, nil
}

func (b *BitcoinProxy) mustEmbedUnimplementedBitcoinProxyServer() {
	//TODO implement me
	panic("implement me")
}