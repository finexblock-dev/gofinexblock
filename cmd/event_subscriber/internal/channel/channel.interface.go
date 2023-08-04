package channel

type Events interface {
}

type Channel interface {
	// Subscribe : Subscribe events and insert events to database.
	Subscribe()
	// Send : Send error message to slack.
	Send(events []Events, err error)
	// Receive : Receive event and send to its own channel.
	Receive(event Events)
}