package usecase

import (
	"errors"
	"reflect"
	"stockbit4/code/movie"
	mock_movie "stockbit4/code/movie/mock"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	u  movie.Usecase
)

func Test_usecase_GetMovieData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbPqConn := pqConnections{
		movieDB: pqConnection{
			master: db,
			slave:  db,
		},
	}

	type args struct {
		searchData movie.SearchData
	}
	tests := []struct {
		name       string
		args       args
		wantMovies []movie.Movie
		expect     func(mockRest *mock_movie.MockRestRepository, mockQuery *mock_movie.MockRepository, args args)
		wantErr    bool
	}{
		{
			name: "Success get movie by exact name",
			args: args{
				searchData: movie.SearchData{
					Keyword: "Koala Kumal",
					Page:    0,
				},
			},
			wantMovies: []movie.Movie{
				{
					Title:  "Koala Kumal",
					Year:   "2014",
					ImdbID: "tt5799566",
				},
			},
			expect: func(mockRest *mock_movie.MockRestRepository, mockQuery *mock_movie.MockRepository, args args) {
				mockQuery.EXPECT().LogSearchData(args.searchData).Return(nil)

				mockQuery.EXPECT().GetMovieByExactName(args.searchData.Keyword).Return(movie.Movie{Title: "Koala Kumal", Year: "2014", ImdbID: "tt5799566"}, nil)
			},
			wantErr: false,
		},
		{
			name: "Error Call omdb api",
			args: args{
				searchData: movie.SearchData{
					Keyword: "Koala Kumal",
					Page:    1,
				},
			},
			wantMovies: []movie.Movie{},
			expect: func(mockRest *mock_movie.MockRestRepository, mockQuery *mock_movie.MockRepository, args args) {
				mockQuery.EXPECT().LogSearchData(args.searchData).Return(nil)
				mockRest.EXPECT().GetMovieData(args.searchData).Return([]movie.Movie{}, errors.New("API Timeout"))
			},
			wantErr: true,
		},
		{
			name: "Exist and Success",
			args: args{
				searchData: movie.SearchData{
					Keyword: "Koala Kumal",
					Page:    1,
				},
			},
			wantMovies: []movie.Movie{{Title: "Koala Kumal", ImdbID: "tt5799566"}},
			expect: func(mockRest *mock_movie.MockRestRepository, mockQuery *mock_movie.MockRepository, args args) {
				mockQuery.EXPECT().LogSearchData(args.searchData).Return(nil)
				mockRest.EXPECT().GetMovieData(args.searchData).Return([]movie.Movie{{Title: "Koala Kumal", ImdbID: "tt5799566"}}, nil)
				mockQuery.EXPECT().IsAlreadyExist("tt5799566").Return(true, nil)

			},
			wantErr: false,
		},
		{
			name: "Not Exist on DB and Success",
			args: args{
				searchData: movie.SearchData{
					Keyword: "Koala Kumal",
					Page:    1,
				},
			},
			wantMovies: []movie.Movie{{Title: "Koala Kumal", ImdbID: "tt5799566"}},
			expect: func(mockRest *mock_movie.MockRestRepository, mockQuery *mock_movie.MockRepository, args args) {
				mockQuery.EXPECT().LogSearchData(args.searchData).Return(nil)
				mockRest.EXPECT().GetMovieData(args.searchData).Return([]movie.Movie{{Title: "Koala Kumal", ImdbID: "tt5799566"}}, nil)
				mockQuery.EXPECT().IsAlreadyExist("tt5799566").Return(false, nil)
				mockQuery.EXPECT().InsertMovieData(movie.Movie{Title: "Koala Kumal", ImdbID: "tt5799566"}).Return(nil)

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		mockQueryRepo := mock_movie.NewMockRepository(ctrl)
		mockRestRepo := mock_movie.NewMockRestRepository(ctrl)
		tt.expect(mockRestRepo, mockQueryRepo, tt.args)
		u := usecase{
			movieRepository:     mockQueryRepo,
			movieRestRepository: mockRestRepo,
			pqConnections:       dbPqConn,
		}
		t.Run(tt.name, func(t *testing.T) {
			gotMovies, err := u.GetMovieData(tt.args.searchData)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetMovieData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotMovies, tt.wantMovies) {
				t.Errorf("usecase.GetMovieData() = %v, want %v", gotMovies, tt.wantMovies)
			}
		})
	}
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	sqlDB, _, _ := sqlmock.New()

	db = sqlx.NewDb(sqlDB, "postgres")

	mockQueryRepo := mock_movie.NewMockRepository(ctrl)
	mockRestRepo := mock_movie.NewMockRestRepository(ctrl)
	u = New(mockRestRepo, mockQueryRepo, db, db)
	type args struct {
		moviesRestRepository movie.RestRepository
		moviesRepository     movie.Repository
		movieMaster          *sqlx.DB
		movieSlave           *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want movie.Usecase
	}{
		{
			name: "Success new usecase",
			args: args{
				moviesRestRepository: mockRestRepo,
				moviesRepository:     mockQueryRepo,
				movieMaster:          db,
				movieSlave:           db,
			},
			want: u,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.moviesRestRepository, tt.args.moviesRepository, tt.args.movieMaster, tt.args.movieSlave); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
