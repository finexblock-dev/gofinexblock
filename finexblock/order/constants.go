package order

type Type string

type Status string

type Reason string

const (
	BID Type = "BID"
	ASK Type = "ASK"
)

const (
	Cancelled     Status = "CANCELLED"
	Placed        Status = "PLACED"
	Fulfilled     Status = "FULFILLED"
	PartialFilled Status = "PARTIAL_FILLED"
)

const (
	Cancel Reason = "CANCEL"
	Fill   Reason = "FILL"
	Place  Reason = "PLACE"
)