package trip

import "errors"

var (
	ErrNoTickets           = errors.New("no tickets provided")
	ErrInvalidTicketFormat = errors.New("invalid ticket format (each ticket must have exactly 2 items)")
	ErrDuplicateDeparture  = errors.New("duplicate departure from airport")
	ErrMultipleStarts      = errors.New("multiple possible starting points found")
	ErrNoValidStart        = errors.New("no valid starting point found")
	ErrIncompleteItinerary = errors.New("incomplete itinerary, some segments may be disconnected or form a cycle")
)
