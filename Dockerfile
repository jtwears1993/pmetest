FROM golang:1.24 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /pmetest -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/api

FROM gcr.io/distroless/static-debian12 as pmetest

WORKDIR /

COPY --from=build-stage /pmetest /pmetest
EXPOSE 4444

ENTRYPOINT ["/pmetest"]
