FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /go/src/github.com/yildizozan/gandalf

# Copy the code into the container
COPY . .

RUN go get -d -v ./...

# Build the application
RUN go build -o /go/bin/gandalf cmd/fasthttp/fasthttp.go

RUN ls -lah /go/bin/
RUN ls -lah /go/src/github.com/yildizozan/gandalf/entrypoint.sh

# Build a small image
#FROM gcr.io/distroless/base
FROM alpine

COPY --from=builder /go/bin/gandalf /
COPY --from=builder /go/src/github.com/yildizozan/gandalf/entrypoint.sh /

#RUN ["chmod", "+x", "/usr/src/app/entrypoint.sh"]
RUN ls -lah /

# Command to run
#ENTRYPOINT ["/entrypoint.sh"]
CMD ["/gandalf"]