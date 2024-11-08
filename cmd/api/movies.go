package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.samuellfa.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.serverErrorResponse(w, r, err)

	}
}
