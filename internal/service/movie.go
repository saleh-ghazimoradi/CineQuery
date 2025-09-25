package service

import (
	"context"
	"github.com/saleh-ghazimoradi/CineQuery/internal/domain"
	"github.com/saleh-ghazimoradi/CineQuery/internal/dto"
	"github.com/saleh-ghazimoradi/CineQuery/internal/repository"
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
}

func (m *movieService) CreateMovie(ctx context.Context, input *dto.Movie) (*domain.Movie, error) {
	movie := &domain.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	if err := m.movieRepository.CreateMovie(ctx, movie); err != nil {
		return nil, err
	}
	return movie, nil
}

func (m *movieService) GetMovieById(ctx context.Context, id int64) (*domain.Movie, error) {
	return m.movieRepository.GetMovieById(ctx, id)
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

func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{
		movieRepository: movieRepository,
	}
}
