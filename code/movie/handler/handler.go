package handler

import (
	"encoding/json"
	"net/http"
	"stockbit4/code/movie"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	restHandler struct {
		movieUsecase movie.Usecase
	}
)

func RegisterRoute(router *httprouter.Router, stockUC movie.Usecase) {
	r := &restHandler{
		movieUsecase: stockUC,
	}
	router.GET("/movie/search/", r.GetMovieData)

}

func (rh *restHandler) GetMovieData(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	req := r.URL.Query()
	MovieTitle := req.Get("movie_title")
	page, _ := strconv.Atoi(req.Get("page")) // if page is zero, the it means we do exact search
	SearchInput := movie.SearchData{
		Keyword: MovieTitle,
		Page:    page,
	}
	MovieLIst, err := rh.movieUsecase.GetMovieData(SearchInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(MovieLIst)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
	return
}
