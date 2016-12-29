FROM    golang:1.7.3
COPY    . /go/src/github.com/moul/gmaps
WORKDIR /go/src/github.com/moul/gmaps
CMD     ["gmaps"]
EXPOSE  8000 9000
RUN     make install
