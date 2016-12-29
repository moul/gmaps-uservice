package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/moul/gmaps-uservice/gen/endpoints"
	"github.com/moul/gmaps-uservice/gen/pb"
	"github.com/moul/gmaps-uservice/gen/transports/grpc"
	"github.com/moul/gmaps-uservice/gen/transports/http"
	"github.com/moul/gmaps-uservice/service"
)

func main() {
	mux := http.NewServeMux()
	ctx := context.Background()
	errc := make(chan error)
	s := grpc.NewServer()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	{
		svc := gmapssvc.New()
		endpoints := gmaps_endpoints.MakeEndpoints(svc)
		srv := gmaps_grpctransport.MakeGRPCServer(ctx, endpoints)
		gmapspb.RegisterGmapsServiceServer(s, srv)
		gmaps_httptransport.RegisterHandlers(ctx, svc, mux, endpoints)
	}

	// start servers
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.NewContext(logger).With("transport", "HTTP")
		logger.Log("addr", ":8000")
		errc <- http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.NewContext(logger).With("transport", "gRPC")
		ln, err := net.Listen("tcp", ":9000")
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", ":9000")
		errc <- s.Serve(ln)
	}()

	logger.Log("exit", <-errc)
}
