package repository

import (
	"database/sql"
	"time"

	"stockbit4/code/movie"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

type (
	repositoryStruct struct {
		pqConnections pqConnections
	}
	pqConnections struct {
		movieDB pqConnection
	}

	pqConnection struct {
		master *sqlx.DB
		slave  *sqlx.DB
	}
)

func NewRepo(movieMaster *sqlx.DB, movieSlave *sqlx.DB) movie.Repository {
	return &repositoryStruct{
		pqConnections: pqConnections{
			movieDB: pqConnection{
				master: movieMaster,
				slave:  movieSlave,
			},
		},
	}
}

func (p *repositoryStruct) LogSearchData(searchData movie.SearchData) (err error) {
	tableName := "movie_search_log"
	query := sq.Insert(tableName).
		Columns("keyword", "page", "time").
		Values(searchData.Keyword, searchData.Page, time.Now()).
		RunWith(p.pqConnections.movieDB.master).
		PlaceholderFormat(sq.Dollar)
	_, err = query.Exec()
	return
}
func (p *repositoryStruct) InsertMovieData(movieDetail movie.Movie) (err error) {

	query := sq.Insert("movie").
		Columns("title", "year", "imdb_id", "type", "poster").
		Values(movieDetail.Title, movieDetail.Year, movieDetail.ImdbID, movieDetail.Type, movieDetail.Poster).
		RunWith(p.pqConnections.movieDB.master).
		PlaceholderFormat(sq.Dollar)
	_, err = query.Exec()

	return
}

func (p *repositoryStruct) IsAlreadyExist(ImdbID string) (exist bool, err error) {

	placeholder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	psql := placeholder.Select("title", "year", "imdb_id", "type", "poster").From("movie")
	postgresSQL := psql
	postgresSQL = psql.Where(sq.Eq{"imdb_id": ImdbID})

	rows, err := postgresSQL.RunWith(p.pqConnections.movieDB.slave).Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return exist, nil
		}

		return exist, err
	}

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			exist = true
			break
		}
	}
	return exist, nil
}

func (p *repositoryStruct) GetMovieByExactName(MovieName string) (movieData movie.Movie, err error) {
	placeholder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	psql := placeholder.Select("title", "year", "imdb_id", "type", "poster").From("movie")
	postgresSQL := psql
	postgresSQL = psql.Where(sq.Eq{"title": MovieName})

	rows, err := postgresSQL.RunWith(p.pqConnections.movieDB.slave).Query()
	if err != nil {
		if err == sql.ErrNoRows {
			return movieData, nil
		}
		return movieData, err
	}
	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&movieData.Title, &movieData.Year, &movieData.ImdbID, &movieData.Type, &movieData.Poster)
			if err != nil {
				return movieData, err
			}
		}
	}
	return movieData, nil
}
