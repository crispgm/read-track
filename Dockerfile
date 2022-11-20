FROM flyio/litefs:0.2 AS litefs

FROM golang:1.19 AS builder
WORKDIR /src/read-track
COPY . .
RUN go build -ldflags "-s -w -extldflags '-static'" -tags osusergo,netgo -o /usr/local/bin/read-track .

FROM alpine

COPY --from=builder /usr/local/bin/read-track /usr/local/bin/read-track
COPY --from=litefs /usr/local/bin/litefs /usr/local/bin/litefs

ADD etc/litefs.yml /etc/litefs.yml

RUN apk add bash curl fuse sqlite tzdata

RUN mkdir -p /data /mnt/data

ENTRYPOINT "litefs"
