package order

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
)

func (o *orderService) Snapshot(tx *gorm.DB, symbolID uint, bids []*grpc_order.Order, asks []*grpc_order.Order) error {
	var symbol string
	switch {
	case len(asks) != 0:
		symbol = asks[0].Symbol.String()
	case len(bids) != 0:
		symbol = bids[0].Symbol.String()
	default:
		return fmt.Errorf("there is no order")
	}

	askOrderList, err := json.Marshal(asks)
	if err != nil {
		return fmt.Errorf("failed to marshal order list: [%v]", err)
	}

	bidOrderList, err := json.Marshal(bids)
	if err != nil {
		return fmt.Errorf("failed to marshal order list: [%v]", err)
	}

	var orderSymbol *order.OrderSymbol
	if err := o.db.Table(orderSymbol.TableName()).Where("name = ?", symbol).First(&orderSymbol).Error; err != nil {
		return fmt.Errorf("failed to get order symbol: [%v]", err)
	}

	var _snapShotOrderBook *order.SnapshotOrderBook
	_snapShotOrderBook.OrderSymbolID = orderSymbol.ID
	_snapShotOrderBook.AskOrderList = string(askOrderList)
	_snapShotOrderBook.BidOrderList = string(bidOrderList)
	if err := o.db.Table(_snapShotOrderBook.TableName()).Create(_snapShotOrderBook).Error; err != nil {
		return fmt.Errorf("failed to create snapshot order book: [%v]", err)
	}

	return nil
}