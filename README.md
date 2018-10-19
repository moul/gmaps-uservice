# gmaps

[![GuardRails badge](https://badges.production.guardrails.io/moul/gmaps-uservice.svg)](https://www.guardrails.io)

:gift: googlemaps micro-service (gRPC + http)

## Usage

server

```console
$ gmaps
2016/12/29 22:38:35 new HTTP endpoint: "/Directions" (service=Gmaps)
2016/12/29 22:38:35 new HTTP endpoint: "/Geocode" (service=Gmaps)
ts=2016-12-29T21:38:35Z caller=main.go:62 transport=HTTP addr=:8000
ts=2016-12-29T21:38:35Z caller=main.go:73 transport=gRPC addr=:9000
::1 - - [29/Dec/2016:22:39:18 +0100] "POST /Geocode HTTP/1.1" 200 3858
::1 - - [29/Dec/2016:22:39:31 +0100] "POST /Directions HTTP/1.1" 200 95725
```

client calling `Geocode`

```console
$ curl -s localhost:8000/Geocode -XPOST -d'{"LatLng":{"lat":39.28091,"lng":-76.61747}}' | jq '.results[].address_components[].long_name'
"701"
"South Sharp Street"
"Sharp Leadenhall"
"Baltimore"
"Maryland"
"United States"
...
```

client calling `Diections`

```console
$ curl -s localhost:8000/Directions -XPOST -d'{"origin":"49.4312981,1.0914374,15z","destination":"48.8966873,2.3161868,17z","departure_time":"now"}' | jq '.[][0].legs[0].steps[].start_location.lat'
46.5060997
46.1404284
44.2241507
43.6185041
42.9677023
...
```
