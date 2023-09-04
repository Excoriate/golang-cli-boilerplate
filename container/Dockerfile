FROM golang:1.21.0-bullseye AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod download \
    && go build -o /bin/app

FROM gcr.io/distroless/base-debian11@sha256:09e4f19add0238d44e45aa3ddab075251d3adac47914eb5437a0de588dcf1626

WORKDIR /

COPY --from=build /bin/app /bin/app

ENTRYPOINT ["/bin/app"]
