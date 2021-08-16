package grpc

import (
	"fmt"
	"log"
	"net"
	"runtime/debug"

	movieGRPC "stockbit4/code/movie/handler"

	pb "stockbit4/pkg/grpc/model"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/valyala/fasthttp/reuseport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var gServer *Service

type Service struct {
	options  *Options
	server   *grpc.Server
	listener net.Listener
}

type Options struct {
	ListenAddress string
	MovieGRPC     *movieGRPC.MovieGRPC
}

func NewListener(hport string) (net.Listener, error) {
	var l net.Listener
	var err error
	l, err = reuseport.Listen("tcp4", hport)
	if err != nil {
		return nil, err
	}

	return l, nil

}

func Init(o *Options) error {
	lis, err := NewListener(o.ListenAddress)
	if err != nil {
		return err
	}
	opts := []recovery.Option{
		recovery.WithRecoveryHandler(panicRecoveryOption),
	}

	s := grpc.NewServer(middleware.WithUnaryServerChain(recovery.UnaryServerInterceptor(opts...)))
	pb.RegisterStockbit4Server(s, o.MovieGRPC)
	reflection.Register(s)
	gServer = &Service{
		options:  o,
		server:   s,
		listener: lis,
	}
	return nil
}

var panicRecoveryOption = func(p interface{}) (err error) {
	trace := debug.Stack()
	log.Println(fmt.Sprintf("PANIC : GRPC Server : %v", p), string(trace))
	err = fmt.Errorf("Fail to complete request : %v", p)
	return
}

// Start will start listening
func Start() {
	if gServer == nil {
		log.Panic("Server is not initialized.")
	}

	go func() {
		log.Printf("Starting gRPC server on %s ...", gServer.options.ListenAddress)
		err := gServer.Run()
		if err != nil {
			log.Println("Running rpc error : ", err)
		}
	}()
}

func Stop() {
	log.Println("Stopping gRPC")
	gServer.server.GracefulStop()
}

func (s *Service) Run() error {
	// return grace.Serve()
	return s.server.Serve(s.listener)
}

func (s *Service) GetServer() *grpc.Server {
	return s.server
}
