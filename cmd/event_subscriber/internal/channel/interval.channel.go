package channel

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

type IntervalChannel struct {
	svc order.Service
}

func NewIntervalChannel(orderService order.Service) *IntervalChannel {
	return &IntervalChannel{svc: orderService}
}

// Subscribe : Insert order_interval data to database.
func (i *IntervalChannel) Subscribe() {
	for {
		g, _ := errgroup.WithContext(context.Background())
		g.Go(func() error {
			i.Schedule(types.OneMinute, time.Minute)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.ThreeMinute, time.Minute*3)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.FiveMinute, time.Minute*5)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.FifteenMinute, time.Minute*15)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.ThirtyMinute, time.Minute*30)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.OneHour, time.Hour)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.TwoHour, time.Hour*2)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.FourHour, time.Hour*4)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.TwelveHour, time.Hour*12)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.OneDay, time.Hour*24)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.ThreeDay, time.Hour*72)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.OneWeek, time.Hour*168)
			return nil
		})
		g.Go(func() error {
			i.Schedule(types.OneMonth, time.Hour*744)
			return nil
		})

		if err := g.Wait(); err != nil {
			i.Send(err)
			log.Panicf("Panic occurred : %v", err)
		}
	}
}

func (i *IntervalChannel) Send(err error) {
	log.Println(err)
}

func (i *IntervalChannel) Schedule(name types.Duration, duration time.Duration) {
	var _interval *entity.OrderInterval
	var sleepTime time.Duration
	var err error
	for {

		now := time.Now()

		// Get recent order_interval, if not exist, panic.
		_interval, err = i.svc.FindRecentIntervalByDuration(name)
		if err != nil {
			log.Println(err)
			continue
		}

		// If now is before end time of recent order_interval, sleep until end time.
		if now.Before(_interval.EndTime) {
			// If now is before end time of recent order_interval, sleep until end time.
			sleepTime = _interval.EndTime.Sub(now)
			time.Sleep(sleepTime)
		}

		if err := i.svc.InsertOrderInterval(name, duration); err != nil {
			log.Printf("error occurred : %v", err)
			i.Send(err)
		}
	}
}