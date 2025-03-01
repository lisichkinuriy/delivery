FROM golang:1.23.2 AS build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/app/main.go


# TODO. test stage


FROM alpine:latest AS release

WORKDIR /

COPY --from=build /app /app

EXPOSE 8081

# TODO. USER nonroot:nonroot

ENTRYPOINT ["/app"]