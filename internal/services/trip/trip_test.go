package trip

import (
	"log/slog"
	"os"
	"slices"
	"testing"
)

func Test_Itinerary(t *testing.T) {
	input := [][]string{
		{"LAX", "DXB"},
		{"JFK", "LAX"},
		{"SFO", "SJC"},
		{"DXB", "SFO"},
	}
	expected := []string{"JFK", "LAX", "DXB", "SFO", "SJC"}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	tripService := New(logger)

	output, err := tripService.Intinerary(input)
	if err != nil {
		t.Errorf("Intinerary returned an error: %s", err)
	}

	result := slices.Equal(output, expected)
	if !result {
		t.Errorf("the output from Itinerary was not what was expected. Got: %s. Expected: %s", output, expected)
	}
}
