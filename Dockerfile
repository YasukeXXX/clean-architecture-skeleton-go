FROM golang:1.13.1-alpine AS build

RUN mkdir /api
WORKDIR /api
RUN apk update && apk upgrade && apk add --no-cache git

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN go build -o ./api

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN apk --no-cache add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

COPY --from=build /api/api /bin/api

EXPOSE 8000
CMD /bin/api

