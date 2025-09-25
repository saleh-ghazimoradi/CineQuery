package handlers

import (
	"errors"
	"github.com/saleh-ghazimoradi/CineQuery/internal/dto"
	"github.com/saleh-ghazimoradi/CineQuery/internal/helper"
	"github.com/saleh-ghazimoradi/CineQuery/internal/repository"
	"github.com/saleh-ghazimoradi/CineQuery/internal/service"
	"github.com/saleh-ghazimoradi/CineQuery/internal/validator"
	"net/http"
)

type MovieHandler struct {
	movieService service.MovieService
	customErr    *helper.CustomErr
	validator    *validator.Validator
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var payload *dto.Movie
	if err := helper.ReadJSON(w, r, &payload); err != nil {
		m.customErr.BadRequestResponse(w, r, err)
		return
	}

	if dto.ValidateMovie(m.validator, payload); !m.validator.Valid() {
		m.customErr.FailedValidationResponse(w, r, m.validator.Errors)
		return
	}

	movie, err := m.movieService.CreateMovie(r.Context(), payload)
	if err != nil {

	}

	if err := helper.WriteJSON(w, http.StatusCreated, helper.Envelope{"movie": movie}, nil); err != nil {
		m.customErr.ServerErrorResponse(w, r, err)
	}
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIdParams(r)
	if err != nil {
		m.customErr.NotFoundResponse(w, r)
		return
	}

	movie, err := m.movieService.GetMovieById(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			m.customErr.NotFoundResponse(w, r)
		default:
			m.customErr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movie": movie}, nil); err != nil {
		m.customErr.ServerErrorResponse(w, r, err)
	}
}

func (m *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {}

func (m *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIdParams(r)
	if err != nil {
		m.customErr.NotFoundResponse(w, r)
		return
	}

	var payload dto.UpdateMovie
	if err := helper.ReadJSON(w, r, &payload); err != nil {
		m.customErr.BadRequestResponse(w, r, err)
		return
	}

	updatedMovie, err := m.movieService.UpdateMovie(r.Context(), id, &payload)

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movie": updatedMovie}, nil); err != nil {
		m.customErr.ServerErrorResponse(w, r, err)
	}
}

func (m *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := helper.ReadIdParams(r)
	if err != nil {
		m.customErr.NotFoundResponse(w, r)
		return
	}

	if err := m.movieService.DeleteMovie(r.Context(), id); err != nil {
		switch {
		case errors.Is(err, repository.ErrRecordNotFound):
			m.customErr.NotFoundResponse(w, r)
		default:
			m.customErr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := helper.WriteJSON(w, http.StatusOK, helper.Envelope{"movie": "movie successfully deleted"}, nil); err != nil {
		m.customErr.ServerErrorResponse(w, r, err)
	}
}

func NewMovieHandler(customErr *helper.CustomErr, movieService service.MovieService, validator *validator.Validator) *MovieHandler {
	return &MovieHandler{
		customErr:    customErr,
		movieService: movieService,
		validator:    validator,
	}
}
