package usecase

import (
	"stockbit4/code/movie"
	"sync"

	"github.com/jmoiron/sqlx"
)

type (
	usecase struct {
		movieRepository     movie.Repository
		movieRestRepository movie.RestRepository

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

func New(moviesRestRepository movie.RestRepository, moviesRepository movie.Repository, movieMaster *sqlx.DB, movieSlave *sqlx.DB) *usecase {
	return &usecase{
		movieRepository:     moviesRepository,
		movieRestRepository: moviesRestRepository,

		pqConnections: pqConnections{
			movieDB: pqConnection{
				master: movieMaster,
				slave:  movieSlave,
			},
		},
	}
}

func (u *usecase) GetMovieData(searchData movie.SearchData) (movies []movie.Movie, err error) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		u.movieRepository.LogSearchData(searchData)
		wg.Done()
	}()
	if searchData.Page < 1 {
		movieDetail, err := u.movieRepository.GetMovieByExactName(searchData.Keyword)
		if err == nil && movieDetail.Title != "" {
			return []movie.Movie{movieDetail}, nil
		}
	}

	movies, err = u.movieRestRepository.GetMovieData(searchData)
	if err != nil {
		return movies, err
	}
	for _, movieData := range movies {
		exist, _ := u.movieRepository.IsAlreadyExist(movieData.ImdbID)
		if !exist {
			wg.Add(1)
			go func() {
				u.movieRepository.InsertMovieData(movieData)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return movies, nil
}
