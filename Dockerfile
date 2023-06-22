FROM golang:1.20 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Run the tests in the container
FROM build AS test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY --from=build /docker-gs-ping /docker-gs-ping

EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]