package trade

import "golang.org/x/sync/errgroup"

func (m *manager) marketOrderMatchingStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MarketOrderMatchingStream.String(), MarketOrderMatchingGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MarketOrderMatchingStream.String(), EventGroup.String())
	})
}

func (m *manager) marketOrderMatchingConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), MarketOrderMatchingGroup.String(), MarketOrderMatchingConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), MarketOrderMatchingGroup.String(), MarketOrderMatchingClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) matchingStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderMatchingStream.String(), OrderMatchingGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderMatchingStream.String(), EventGroup.String())
	})
}

func (m *manager) matchingConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), OrderMatchingGroup.String(), OrderMatchingConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), OrderMatchingGroup.String(), OrderMatchingClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) partialFillStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPartialFillStream.String(), OrderPartialFillGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPartialFillStream.String(), EventGroup.String())
	})
}

func (m *manager) partialFillConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), OrderPartialFillGroup.String(), OrderPartialFillConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), OrderPartialFillGroup.String(), OrderPartialFillClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) fulfillmentStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderFulfillmentStream.String(), OrderFulfillmentGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderFulfillmentStream.String(), EventGroup.String())
	})
}

func (m *manager) fulfillmentConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), EventGroup.String(), EventClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), OrderFulfillmentGroup.String(), OrderFulfillmentConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), OrderFulfillmentGroup.String(), OrderFulfillmentClaimer.String())
	})
}

func (m *manager) initializeStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderInitializeStream.String(), OrderInitializeGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderInitializeStream.String(), EventGroup.String())
	})
}

func (m *manager) initializeConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), OrderInitializeGroup.String(), OrderInitializeConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), OrderInitializeGroup.String(), OrderInitializeClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) balanceUpdateStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(BalanceUpdateStream.String(), BalanceUpdateGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(BalanceUpdateStream.String(), EventGroup.String())
	})
}

func (m *manager) balanceUpdateConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), BalanceUpdateGroup.String(), BalanceUpdateConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), BalanceUpdateGroup.String(), BalanceUpdateClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) cancellationStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderCancellationStream.String(), OrderCancellationGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderCancellationStream.String(), EventGroup.String())
	})
}

func (m *manager) cancellationConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderCancellationStream.String(), OrderCancellationGroup.String(), OrderCancellationConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderCancellationStream.String(), OrderCancellationGroup.String(), OrderCancellationClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderCancellationStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderCancellationStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) errorStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(ErrorStream.String(), ErrorGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(ErrorStream.String(), EventGroup.String())
	})
}

func (m *manager) errorConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), ErrorGroup.String(), ErrorConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), ErrorGroup.String(), ErrorClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) matchStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MatchStream.String(), MatchGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MatchStream.String(), EventGroup.String())
	})
}

func (m *manager) matchConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), MatchGroup.String(), MatchConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), MatchGroup.String(), MatchClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), EventGroup.String(), EventClaimer.String())
	})
}

func (m *manager) placementStreamInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPlacementStream.String(), OrderPlacementGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPlacementStream.String(), EventGroup.String())
	})
}

func (m *manager) placementConsumerInit(group *errgroup.Group) {
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), OrderPlacementGroup.String(), OrderPlacementConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), OrderPlacementGroup.String(), OrderPlacementClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), EventGroup.String(), EventClaimer.String())
	})
}