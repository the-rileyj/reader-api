FROM golang:1.11.3-alpine3.8 AS buildenv

RUN apk add --no-cache git

WORKDIR /go/src/github.com/the-rileyj/reader-api

RUN mkdir ./functionality

COPY main.go .
COPY ./functionality/functionality.go ./functionality

RUN go get -d -v ./...
RUN env CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -a -v -o reader-api

# Second stage of build, adding in files and running
# newly compiled program
FROM alpine

# Nesessary for making requests to most https sites
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

EXPOSE 80

# Create and navigate into /app so that the files we
# bring in aren't cluttered with the dirs in /
WORKDIR /app

# Copy the *.go program compiled in the first stage
COPY --from=buildenv /go/src/github.com/the-rileyj/reader-api/reader-api .

ENTRYPOINT ./reader-api