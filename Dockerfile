FROM golang:1.17

WORKDIR /usr/src/go-rest-api

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make 

CMD ["./apiserver.exe"]