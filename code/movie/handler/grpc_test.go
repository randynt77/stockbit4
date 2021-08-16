package handler

import (
	"context"
	"errors"
	"reflect"
	"stockbit4/code/movie"
	mock_movie "stockbit4/code/movie/mock"
	pb "stockbit4/pkg/grpc/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMovieGRPC_GetMovieData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		ctx   context.Context
		input *pb.GetMovieDataRequest
	}
	tests := []struct {
		name    string
		args    args
		expect  func(mockUsecase *mock_movie.MockUsecase, args args)
		wantErr bool
	}{
		{
			name: "Error Calling Usecase",
			args: args{
				ctx: context.Background(),
				input: &pb.GetMovieDataRequest{
					Keyword: "Koala",
					Page:    1,
				},
			},
			expect: func(mockUsecase *mock_movie.MockUsecase, args args) {
				mockUsecase.EXPECT().GetMovieData(movie.SearchData{Keyword: "Koala", Page: 1}).Return([]movie.Movie{}, errors.New("error from usecase"))
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				input: &pb.GetMovieDataRequest{
					Keyword: "Koala",
					Page:    1,
				},
			},
			expect: func(mockUsecase *mock_movie.MockUsecase, args args) {
				mockUsecase.EXPECT().GetMovieData(movie.SearchData{Keyword: "Koala", Page: 1}).Return([]movie.Movie{{Title: "Koala Kumal", Year: "2014", Type: "Movie"}}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase := mock_movie.NewMockUsecase(ctrl)
			tt.expect(mockUsecase, tt.args)
			s := MovieGRPC{
				movieUsecase: mockUsecase,
			}
			_, err := s.GetMovieData(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("MovieGRPC.GetMovieData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewGRPC(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockUsecase := mock_movie.NewMockUsecase(ctrl)
	grpcNew := NewGRPC(mockUsecase)
	type args struct {
		movieUsecase movie.Usecase
	}
	tests := []struct {
		name string
		args args
		want *MovieGRPC
	}{
		{
			name: "Success",
			args: args{
				movieUsecase: mockUsecase,
			},
			want: grpcNew,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGRPC(tt.args.movieUsecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGRPC() = %v, want %v", got, tt.want)
			}
		})
	}
}
