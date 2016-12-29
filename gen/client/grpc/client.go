package gmaps_clientgrpc

import (
	jwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	endpoints "github.com/moul/gmaps-uservice/gen/endpoints"
	pb "github.com/moul/gmaps-uservice/gen/pb"
)

func New(conn *grpc.ClientConn, logger log.Logger) pb.GmapsServiceServer {

	var directionsEndpoint endpoint.Endpoint
	{
		directionsEndpoint = grpctransport.NewClient(
			conn,
			"gmaps.GmapsService",
			"Directions",
			EncodeDirectionsRequest,
			DecodeDirectionsResponse,
			pb.DirectionsResponse{},
			append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.FromGRPCContext()))...,
		).Endpoint()
	}

	return &endpoints.Endpoints{

		DirectionsEndpoint: directionsEndpoint,
	}
}

func EncodeDirectionsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DirectionsRequest)
	return req, nil
}

func DecodeDirectionsResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.DirectionsResponse)
	return response, nil
}
