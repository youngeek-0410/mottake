FROM golang:1.17.0

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go get -u github.com/cosmtrek/air

CMD [ "air", "-c", ".air.toml" ]

