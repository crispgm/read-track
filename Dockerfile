FROM golang:1.19 AS builder
WORKDIR /src/read-track
COPY . .
RUN go build -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo -o /usr/local/bin/read-track .

FROM alpine
RUN apk add bash curl fuse sqlite tzdata

COPY --from=builder /usr/local/bin/read-track /usr/local/bin/read-track

RUN mkdir -p /app
COPY --from=builder /src/read-track/templates /app/templates
COPY --from=builder /src/read-track/static /app/static

ENTRYPOINT ["read-track", "-working-path=/app"]
