package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"pmetest/internal/response"
	"pmetest/internal/services/trip"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.logger.Error("error when marshaling status response", slog.String("err", err.Error()))
		app.serverError(w, r, err)
	}
}

func (app *application) itineraryHandler(w http.ResponseWriter, r *http.Request) {
	var tickets [][]string
	if err := json.NewDecoder(r.Body).Decode(&tickets); err != nil {
		app.logger.Error("error unmarshaling itinerary post request body", slog.String("err", err.Error()))
		app.badRequest(w, r, err)
	}

	itinerary, err := app.tripService.Intinerary(tickets)
	if err != nil {
		if isClientError(err) {
			app.logger.Error("client error in trip.Itinerary", slog.String("err", err.Error()))
			app.badRequest(w, r, err)
		}
		app.logger.Error("error while working out trip itinerary", slog.String("err", err.Error()))
		app.serverError(w, r, err)
	}

	err = response.JSON(w,
		http.StatusOK,
		itinerary)
	if err != nil {
		app.logger.Error("error marshaling itinerary response", slog.String("err", err.Error()))
		app.serverError(w, r, err)
	}
}

func isClientError(err error) bool {
	switch {
	case errors.Is(err, trip.ErrInvalidTicketFormat),
		errors.Is(err, trip.ErrDuplicateDeparture),
		errors.Is(err, trip.ErrNoValidStart),
		errors.Is(err, trip.ErrMultipleStarts),
		errors.Is(err, trip.ErrIncompleteItinerary):
		return true
	}
	return false
}
