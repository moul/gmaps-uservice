package gmapssvc

import (
	"encoding/json"

	"golang.org/x/net/context"
	googlemap "googlemaps.github.io/maps"

	"github.com/moul/gmaps-uservice/gen/pb"
)

type Service struct {
	gm *googlemap.Client
}

func New(gm *googlemap.Client) gmapspb.GmapsServiceServer {
	return &Service{gm: gm}
}

func (s *Service) Directions(ctx context.Context, in *gmapspb.DirectionsRequest) (*gmapspb.DirectionsResponse, error) {
	inJson, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	var req googlemap.DirectionsRequest
	err = json.Unmarshal(inJson, &req)
	if err != nil {
		return nil, err
	}

	resp, _, err := s.gm.Directions(ctx, &req)
	if err != nil {
		return nil, err
	}

	outJson, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	var rep gmapspb.DirectionsResponse
	err = json.Unmarshal(outJson, &rep)
	return &rep, err
}
