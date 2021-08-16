package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"stockbit4/code/movie"
	mock_movie "stockbit4/code/movie/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
)

func TestRegisterRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mock_movie.NewMockUsecase(ctrl)
	type args struct {
		router  *httprouter.Router
		stockUC movie.Usecase
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Success",
			args: args{
				router:  httprouter.New(),
				stockUC: mockUsecase,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterRoute(tt.args.router, tt.args.stockUC)
		})
	}
}

func Test_restHandler_GetMovieData(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name   string
		args   func(t *testing.T) args
		expect func(mockUsecase *mock_movie.MockUsecase, args args)
	}{
		{
			name: "Error call Usecase",
			expect: func(mockUsecase *mock_movie.MockUsecase, args args) {
				mockUsecase.EXPECT().GetMovieData(movie.SearchData{Keyword: "koala", Page: 1}).Return([]movie.Movie{}, errors.New("error from usecase"))
			},
			args: func(t *testing.T) args {
				req := http.Request{
					Method: "GET",
					Header: make(http.Header),
					URL: &url.URL{
						RawQuery: "movie_title=koala&page=1",
					},
				}
				params := []httprouter.Param{}
				a := args{
					w:      httptest.NewRecorder(),
					r:      &req,
					params: params,
				}

				return a
			},
		},
		{
			name: "Success",
			expect: func(mockUsecase *mock_movie.MockUsecase, args args) {
				mockUsecase.EXPECT().GetMovieData(movie.SearchData{Keyword: "koala", Page: 1}).Return([]movie.Movie{{Title: "Koala Kumal", Year: "2014", ImdbID: "Xxxx", Type: "movie"}}, nil)
			},
			args: func(t *testing.T) args {
				req := http.Request{
					Method: "GET",
					Header: make(http.Header),
					URL: &url.URL{
						RawQuery: "movie_title=koala&page=1",
					},
				}
				params := []httprouter.Param{}
				a := args{
					w:      httptest.NewRecorder(),
					r:      &req,
					params: params,
				}

				return a
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args(t)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockUsecase := mock_movie.NewMockUsecase(ctrl)
			tt.expect(mockUsecase, args)
			rh := &restHandler{
				movieUsecase: mockUsecase,
			}
			rh.GetMovieData(args.w, args.r, args.params)
		})
	}
}
