package gmaps_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/gmaps-uservice/gen/endpoints"
	pb "github.com/moul/gmaps-uservice/gen/pb"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient

func MakeDirectionsHandler(ctx context.Context, svc pb.GmapsServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeDirectionsRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeDirectionsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.DirectionsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func MakeGeocodeHandler(ctx context.Context, svc pb.GmapsServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeGeocodeRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeGeocodeRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.GeocodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(ctx context.Context, svc pb.GmapsServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/Directions\" (service=Gmaps)")
	mux.Handle("/Directions", MakeDirectionsHandler(ctx, svc, endpoints.DirectionsEndpoint))

	log.Println("new HTTP endpoint: \"/Geocode\" (service=Gmaps)")
	mux.Handle("/Geocode", MakeGeocodeHandler(ctx, svc, endpoints.GeocodeEndpoint))

	return nil
}
