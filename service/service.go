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
	if err = json.Unmarshal(inJson, &req); err != nil {
		return nil, err
	}

	routes, geocodedWaypoint, err := s.gm.Directions(ctx, &req)
	if err != nil {
		return nil, err
	}

	var rep gmapspb.DirectionsResponse
	outJson, err := json.Marshal(routes)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(outJson, &rep.Routes); err != nil {
		return nil, err
	}
	outJson, err = json.Marshal(geocodedWaypoint)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(outJson, &rep.GeocodedWaypoint); err != nil {
		return nil, err
	}

	return &rep, nil
}
