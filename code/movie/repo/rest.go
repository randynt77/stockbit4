package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"stockbit4/code/movie"
	"time"
)

type (
	restRepo struct {
	}
)

func NewRest() movie.RestRepository {
	return &restRepo{}
}

func (r *restRepo) GetMovieData(requestParam movie.SearchData) (movieList []movie.Movie, err error) {
	var omdbResp movie.OmdbResponse
	timeout := 5000 * time.Millisecond
	transport := &http.Transport{}
	httpClient := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	defer transport.CloseIdleConnections()

	host := "https://www.omdbapi.com/"
	key := "faf7e5bb"
	keyword := url.PathEscape(requestParam.Keyword)
	url := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", host, key, keyword, requestParam.Page)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&omdbResp)
	if err != nil {
		return
	}
	return omdbResp.Search, nil
}
