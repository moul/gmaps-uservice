DOCKER_IMAGE ?=	moul/gmaps

.PHONY: build
build: gmaps

gmaps: gen/pb/gmaps.pb.go cmd/gmaps/main.go service/service.go
	go build -o gmaps ./cmd/gmaps

gen/pb/gmaps.pb.go:	pb/gmaps.proto
	@mkdir -p gen/pb
	cd pb; protoc --gotemplate_out=destination_dir=../gen,template_dir=../vendor/github.com/moul/protoc-gen-gotemplate/examples/go-kit/templates/{{.File.Package}}/gen:../gen ./gmaps.proto
	gofmt -w gen
	cd pb; protoc --gogo_out=plugins=grpc:../gen/pb ./gmaps.proto

.PHONY: stats
stats:
	wc -l service/service.go cmd/*/*.go pb/*.proto
	wc -l $(shell find gen -name "*.go")

.PHONY: test
test:
	go test -v $(shell go list ./... | grep -v /vendor/)

.PHONY: install
install:
	go install ./cmd/gmaps

.PHONY: docker.build
docker.build:
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker.run
docker.run:
	docker run -p 8000:8000 -p 9000:9000 $(DOCKER_IMAGE)

.PHONY: docker.test
docker.test: docker.build
	docker run $(DOCKER_IMAGE) make test

.PHONY: clean
clean:
	rm -rf gen

.PHONY: re
re: clean build
