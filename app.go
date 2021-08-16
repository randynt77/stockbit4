package main

import (
	"fmt"
	movie_handler "stockbit4/code/movie/handler"
	movie_repo "stockbit4/code/movie/repo"

	movie_usecase "stockbit4/code/movie/usecase"
	"stockbit4/pkg/config"
	"stockbit4/pkg/grpc"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/paytm/grace.v1"
)

func main() {
	// Init Config
	appConfig := config.InitConfig()
	dbConnection := appConfig.GetDatabaseConns()
	router := httprouter.New()

	//Init Repository
	repoRepo := movie_repo.NewRepo(dbConnection.Movie.Master, dbConnection.Movie.Slave)
	restRepo := movie_repo.NewRest()
	//init usecase
	usecase := movie_usecase.New(restRepo, repoRepo, dbConnection.Movie.Master, dbConnection.Movie.Slave)

	//Init Handler router
	movie_handler.RegisterRoute(router, usecase)

	// init grpc handler
	movieGRPC := movie_handler.NewGRPC(usecase)

	// init and start GRPC server
	err := grpc.Init(&grpc.Options{
		ListenAddress: ":7777",
		MovieGRPC:     movieGRPC,
	})
	if err != nil {
		fmt.Printf("Error Init gRPC config %v", err)
	}
	grpc.Start()

	// start gracefull service
	err = grace.Serve(":1234", router)
	if err != nil {
		fmt.Println("ERROR graceful", err)
		return
	}
}
