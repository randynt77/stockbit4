package handler

import (
	"context"
	"stockbit4/code/movie"

	pb "stockbit4/pkg/grpc/model"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	MovieGRPC struct {
		movieUsecase movie.Usecase
		pb.UnimplementedStockbit4Server
	}
)

func NewGRPC(movieUsecase movie.Usecase) *MovieGRPC {
	return &MovieGRPC{
		movieUsecase: movieUsecase,
	}
}

func (s *MovieGRPC) GetMovieData(ctx context.Context, input *pb.GetMovieDataRequest) (response *pb.GetMovieDataResponse, err error) {

	inputMovieData := movie.SearchData{
		Keyword: input.Keyword,
		Page:    int(input.Page),
	}
	movies, err := s.movieUsecase.GetMovieData(inputMovieData)
	if err != nil {
		return response, status.Errorf(codes.Internal, "Failed Upsert [%v]", err)
	}

	var movieDataArr []*pb.MovieData
	for _, val := range movies {
		movieData := pb.MovieData{}
		movieData.ImdbID = val.ImdbID
		movieData.Title = val.Title
		movieData.Poster = val.Poster
		movieData.Year = val.Year
		movieData.Type = val.Type
		movieDataArr = append(movieDataArr, &movieData)

	}
	response = &pb.GetMovieDataResponse{
		MoviesData: movieDataArr,
	}

	return response, nil
}
