package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	ERC20 "github.com/finexblock-dev/gofinexblock/finexblock/gen/contracts"
)

func (g *gethClient) NewERC20(address string) (*ERC20.ERC20, error) {
	var instance *ERC20.ERC20
	var err error

	instance, err = ERC20.NewERC20(common.HexToAddress(address), g.conn)
	if err != nil {
		return nil, err
	}

	return instance, nil
}