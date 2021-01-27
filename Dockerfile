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
RUN go build -o /go/bin/gandalf

# Build a small image
#FROM gcr.io/distroless/base
FROM alpine

COPY --from=builder /go/bin/gandalf /
COPY --from=builder /go/src/github.com/yildizozan/gandalf/entrypoint.sh /

# Command to run
ENTRYPOINT [ "/gandalf" ]
