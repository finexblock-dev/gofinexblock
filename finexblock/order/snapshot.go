package order

import (
	"encoding/json"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"gorm.io/gorm"
)

func (o *orderService) Snapshot(tx *gorm.DB, symbolID uint, bids []*grpc_order.Order, asks []*grpc_order.Order) error {

	askOrderList, err := json.Marshal(asks)
	if err != nil {
		return fmt.Errorf("failed to marshal order list: [%v]", err)
	}

	bidOrderList, err := json.Marshal(bids)
	if err != nil {
		return fmt.Errorf("failed to marshal order list: [%v]", err)
	}

	var _snapShotOrderBook = &order.SnapshotOrderBook{
		OrderSymbolID: symbolID,
		BidOrderList:  string(bidOrderList),
		AskOrderList:  string(askOrderList),
	}
	if err := o.db.Table(_snapShotOrderBook.TableName()).Create(_snapShotOrderBook).Error; err != nil {
		return fmt.Errorf("failed to create snapshot order book: [%v]", err)
	}

	return nil
}