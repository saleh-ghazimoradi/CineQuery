package repository

import (
	"context"
	"database/sql"
	"github.com/saleh-ghazimoradi/CineQuery/internal/domain"
)

type MovieRepository interface {
	CreateMovie(ctx context.Context, movie *domain.Movie) error
	GetMovieById(ctx context.Context, id int64) (*domain.Movie, error)
	GetMovies(ctx context.Context, offset, limit int32) ([]*domain.Movie, error)
	UpdateMovie(ctx context.Context, movie *domain.Movie) error
	DeleteMovie(ctx context.Context, id int64) error
	WithTx(tx *sql.Tx) MovieRepository
}

type movieRepository struct {
	dbWrite *sql.DB
	dbRead  *sql.DB
	tx      *sql.Tx
}

func (m *movieRepository) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	return nil
}

func (m *movieRepository) GetMovieById(ctx context.Context, id int64) (*domain.Movie, error) {
	return nil, nil
}

func (m *movieRepository) GetMovies(ctx context.Context, offset, limit int32) ([]*domain.Movie, error) {
	return nil, nil
}

func (m *movieRepository) UpdateMovie(ctx context.Context, movie *domain.Movie) error {
	return nil
}

func (m *movieRepository) DeleteMovie(ctx context.Context, id int64) error {
	return nil
}

func (m *movieRepository) WithTx(tx *sql.Tx) MovieRepository {
	return &movieRepository{
		dbWrite: m.dbWrite,
		dbRead:  m.dbRead,
		tx:      tx,
	}
}

func NewMovieRepository(dbWrite, dbRead *sql.DB) MovieRepository {
	return &movieRepository{
		dbWrite: dbWrite,
		dbRead:  dbRead,
	}
}
