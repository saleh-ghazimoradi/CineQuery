package service

import (
	"context"
	"github.com/saleh-ghazimoradi/CineQuery/internal/domain"
	"github.com/saleh-ghazimoradi/CineQuery/internal/dto"
	"github.com/saleh-ghazimoradi/CineQuery/internal/repository"
	"github.com/saleh-ghazimoradi/CineQuery/internal/validator"
)

type MovieService interface {
	CreateMovie(ctx context.Context, input *dto.Movie) (*domain.Movie, error)
	GetMovieById(ctx context.Context, id int64) (*domain.Movie, error)
	GetMovies(ctx context.Context, offset, limit int32) ([]*domain.Movie, error)
	UpdateMovie(ctx context.Context, input *dto.Movie) (*domain.Movie, error)
	DeleteMovie(ctx context.Context, id int64) error
}

type movieService struct {
	movieRepository repository.MovieRepository
	validator       *validator.Validator
}

func (m *movieService) CreateMovie(ctx context.Context, input *dto.Movie) (*domain.Movie, error) {
	return nil, nil
}

func (m *movieService) GetMovieById(ctx context.Context, id int64) (*domain.Movie, error) {
	return nil, nil
}

func (m *movieService) GetMovies(ctx context.Context, offset, limit int32) ([]*domain.Movie, error) {
	return nil, nil
}

func (m *movieService) UpdateMovie(ctx context.Context, input *dto.Movie) (*domain.Movie, error) {
	return nil, nil
}

func (m *movieService) DeleteMovie(ctx context.Context, id int64) error {
	return nil
}

func NewMovieService(movieRepository repository.MovieRepository, validator *validator.Validator) MovieService {
	return &movieService{
		movieRepository: movieRepository,
		validator:       validator,
	}
}
