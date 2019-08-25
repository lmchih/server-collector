FROM golang:1.12.7-stretch as builder

COPY go.mod /go/src/github.com/lmchih/server-collector/go.mod
COPY go.sum /go/src/github.com/lmchih/server-collector/go.sum

# Run golang at any directory, not neccessary $GOROOT, $GOPATH
ENV GO111MODULE=on
WORKDIR /go/src/github.com/lmchih/server-collector

# RUN go mod init github.com/lmchih/server-collector
RUN go mod download
COPY cmd /go/src/github.com/lmchih/server-collector/cmd
COPY internal /go/src/github.com/lmchih/server-collector/internal

# Build the Go app
RUN env GOOS=linux GOARCH=amd64 go build -o /server-collector -v -ldflags "-s" github.com/lmchih/server-collector/cmd/server-collector

# To reduce the final image size, start a new stage with alpine from scratch
FROM alpine:3.9
RUN apk --no-cache add ca-certificates libc6-compat

# Copy the pre-built binary file from the previous stage
COPY --from=builder /server-collector /usr/local/bin/server-collector

# Run the executable
CMD [ "server-collector" ] 