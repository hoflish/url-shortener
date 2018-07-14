# Initialize and name a new build stage,
# Set the 'golang:1.9.0' image for subsequent instructions
FROM golang:1.9.0 as builder

# install dep
RUN go get -u github.com/golang/dep/cmd/dep

# setup the working directory
WORKDIR /go/src/github.com/hoflish/url-shortener/api

# add Gopkg.toml and Gopkg.lock
COPY api/Gopkg.toml .
COPY api/Gopkg.lock .

# install packages
# --vendor-only is used to restrict dep from scanning source code
# and finding dependencies
RUN dep ensure --vendor-only

# copy source code
COPY api/ .

# build the source
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go router.go

# use a minimal alpine image
FROM alpine:latest

# add ca-certificates in case you need them
RUN apk --no-cache add ca-certificates

# create a working directory
RUN mkdir /app

# set working directory
WORKDIR /app

# copy the binary from builder
COPY --from=builder /go/src/github.com/hoflish/url-shortener/api/main .
COPY --from=builder /go/src/github.com/hoflish/url-shortener/api/log .
# run the binary
CMD ["./main"]