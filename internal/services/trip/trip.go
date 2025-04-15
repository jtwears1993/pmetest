package trip

import (
	"fmt"
	"log/slog"
)

type TripService struct {
	Logger *slog.Logger
}

func New(logger *slog.Logger) *TripService {
	return &TripService{
		Logger: logger,
	}
}

func (ts *TripService) Intinerary(tickets [][]string) ([]string, error) {
	ts.Logger.Info("received request", "ticket_count", len(tickets))

	if len(tickets) == 0 {
		ts.Logger.Warn("no tickets provided")
		return nil, ErrNoTickets
	}

	graph, inDegree, _, err := ts.buildGraph(tickets)
	if err != nil {
		ts.Logger.Error("failed to build graph", "error", err)
		return nil, err
	}
	ts.Logger.Debug("graph built", "graph", graph, "inDegree", inDegree)

	start, err := ts.findStartingPoint(graph, inDegree)
	if err != nil {
		ts.Logger.Error("failed to find starting point", "error", err)
		return nil, err
	}
	ts.Logger.Info("starting point found", "start", start)

	var itinerary []string
	current := start
	for {
		itinerary = append(itinerary, current)
		next, exists := graph[current]
		if !exists {
			break
		}
		current = next
	}

	if len(itinerary) != len(tickets)+1 {
		ts.Logger.Warn("incomplete itinerary", "expected_len", len(tickets)+1, "actual_len", len(itinerary))
		return nil, ErrIncompleteItinerary
	}

	ts.Logger.Info("itinerary successfully reconstructed", "itinerary", itinerary)
	return itinerary, nil
}

func (ts *TripService) buildGraph(tickets [][]string) (map[string]string, map[string]int, map[string]int, error) {
	graph := make(map[string]string)
	inDegree := make(map[string]int)
	outDegree := make(map[string]int)

	for _, ticket := range tickets {
		if len(ticket) != 2 {
			ts.Logger.Warn("invalid ticket format", "ticket", ticket)
			return nil, nil, nil, ErrInvalidTicketFormat
		}
		from, to := ticket[0], ticket[1]
		if _, exists := graph[from]; exists {
			ts.Logger.Warn("duplicate departure detected", "from", from)
			return nil, nil, nil, fmt.Errorf("%w: %s", ErrDuplicateDeparture, from)
		}
		graph[from] = to
		outDegree[from]++
		inDegree[to]++
	}

	return graph, inDegree, outDegree, nil
}

func (ts *TripService) findStartingPoint(graph map[string]string, inDegree map[string]int) (string, error) {
	var start string
	for from := range graph {
		if inDegree[from] == 0 {
			if start != "" {
				ts.Logger.Warn("multiple starting points found", "first", start, "duplicate", from)
				return "", ErrMultipleStarts
			}
			start = from
		}
	}
	if start == "" {
		ts.Logger.Warn("no valid starting point found")
		return "", ErrNoValidStart
	}
	return start, nil
}
