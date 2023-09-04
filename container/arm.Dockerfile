# Build stage
FROM golang:1.21.0-bullseye AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go mod download \
    && go build -o /bin/app

# Final stage
FROM debian:bullseye

WORKDIR /

COPY --from=build /bin/app /bin/app

CMD ["/bin/app"]
