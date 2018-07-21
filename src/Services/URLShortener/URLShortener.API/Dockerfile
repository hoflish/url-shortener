# Initialize and name a new build stage,
# Set the 'golang:1.9.0' image for subsequent instructions
FROM golang:1.9.0 as builder

RUN apt-get update && apt-get install -y xz-utils \
    && rm -rf /var/lib/apt/lists/*

# install UPX
ADD https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz /usr/local

RUN xz -d -c /usr/local/upx-3.94-amd64_linux.tar.xz | \
    tar -xOf - upx-3.94-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx

# install dep
RUN go get github.com/golang/dep/cmd/dep

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
# strip and compress the binary
RUN strip --strip-unneeded main
RUN upx main

# use scratch (base for a docker image)
FROM scratch

# add ca-certificates in case you need them
#RUN apk --no-cache add ca-certificates

# set working directory
WORKDIR /app

# copy the binary from builder
COPY --from=builder /go/src/github.com/hoflish/url-shortener/api/main .
COPY --from=builder /go/src/github.com/hoflish/url-shortener/api/log .
# run the binary
CMD ["./main"]