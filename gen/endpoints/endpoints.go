package gmaps_endpoints

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/moul/gmaps-uservice/gen/pb"
	context "golang.org/x/net/context"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	DirectionsEndpoint endpoint.Endpoint
}

func (e *Endpoints) Directions(ctx context.Context, in *pb.DirectionsRequest) (*pb.DirectionsResponse, error) {
	out, err := e.DirectionsEndpoint(ctx, in)
	if err != nil {
		return &pb.DirectionsResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.DirectionsResponse), err
}

func MakeDirectionsEndpoint(svc pb.GmapsServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.DirectionsRequest)
		rep, err := svc.Directions(ctx, req)
		if err != nil {
			return &pb.DirectionsResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeEndpoints(svc pb.GmapsServiceServer) Endpoints {
	return Endpoints{

		DirectionsEndpoint: MakeDirectionsEndpoint(svc),
	}
}
