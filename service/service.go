package gmapssvc

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/moul/gmaps-uservice/gen/pb"
)

type Service struct{}

func New() gmapspb.GmapsServiceServer {
	return &Service{}
}

func (s *Service) Directions(ctx context.Context, in *gmapspb.DirectionsRequest) (*gmapspb.DirectionsResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
