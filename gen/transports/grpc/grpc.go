package gmaps_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	endpoints "github.com/moul/gmaps-uservice/gen/endpoints"
	pb "github.com/moul/gmaps-uservice/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.GmapsServiceServer {
	var options []grpctransport.ServerOption
	_ = options
	return &grpcServer{

		directions: grpctransport.NewServer(
			ctx,
			endpoints.DirectionsEndpoint,
			decodeRequest,
			encodeDirectionsResponse,
			options...,
		),

		geocode: grpctransport.NewServer(
			ctx,
			endpoints.GeocodeEndpoint,
			decodeRequest,
			encodeGeocodeResponse,
			options...,
		),
	}
}

type grpcServer struct {
	directions grpctransport.Handler

	geocode grpctransport.Handler
}

func (s *grpcServer) Directions(ctx context.Context, req *pb.DirectionsRequest) (*pb.DirectionsResponse, error) {
	_, rep, err := s.directions.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DirectionsResponse), nil
}

func encodeDirectionsResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.DirectionsResponse)
	return resp, nil
}

func (s *grpcServer) Geocode(ctx context.Context, req *pb.GeocodeRequest) (*pb.GeocodeResponse, error) {
	_, rep, err := s.geocode.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GeocodeResponse), nil
}

func encodeGeocodeResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GeocodeResponse)
	return resp, nil
}

func decodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

type streamHandler interface {
	Do(server interface{}, req interface{}) (err error)
}

type server struct {
	e endpoints.StreamEndpoint
}

func (s server) Do(server interface{}, req interface{}) (err error) {
	if err := s.e(server, req); err != nil {
		return err
	}
	return nil
}
