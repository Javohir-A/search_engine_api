package storage

import (
	"context"
	"database/sql"
	"search_engine/internal/model"
	"search_engine/internal/search"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type MovieStorageImpl struct {
	db           *sql.DB
	sqlBuilder   sq.StatementBuilderType
	indexManager *search.IndexManager
}

func NewMovieStorage(db *sql.DB, builder sq.StatementBuilderType, indexManager *search.IndexManager) *MovieStorageImpl {
	return &MovieStorageImpl{
		db:           db,
		sqlBuilder:   builder,
		indexManager: indexManager,
	}
}

func (m *MovieStorageImpl) CreateMovie(ctx context.Context, movie *model.Movie) (*model.MovieResponse, error) {
	newID := uuid.NewString()

	builder := m.sqlBuilder.Insert("movies").
		Columns("id", "title", "director", "release_year", "genre", "plot", "actors").
		Values(newID, movie.Title, movie.Director, movie.ReleaseYear, movie.Genre, movie.Plot, movie.Actors).
		Suffix("RETURNING created_at, updated_at")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	row := m.db.QueryRowContext(ctx, query, args...)

	response := &model.MovieResponse{
		Id:          newID,
		Title:       movie.Title,
		Director:    movie.Director,
		ReleaseYear: movie.ReleaseYear,
		Genre:       movie.Genre,
		Plot:        movie.Plot,
		Actors:      movie.Actors,
	}

	err = row.Scan(&response.CreatedAt, &response.UpdatedAt)
	if err != nil {
		return nil, err
	}

	if err := m.indexManager.IndexMovie(newID, movie); err != nil {
		return nil, err
	}

	return response, nil
}
