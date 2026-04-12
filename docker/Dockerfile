# Use the official Go image to create a binary. This is based on Debian and sets the GOPATH to /go. https://hub.docker.com/_/golang
FROM golang:1.25-bookworm AS builder

# Create and change to the app directory.
WORKDIR /app

RUN apt-get update && apt-get install -y git-all
RUN git clone https://bitbucket.org/cankko/passm.git
RUN mv passm/* /app/
# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
RUN go mod download

# Build the binary.
RUN go build -o passmApp ./cmd/main.go

# remove all repository files except data, web folder and build binary file
# RUN find . -maxdepth 1 ! -name 'data' ! -name 'web' ! -name 'passmApp' ! -name '.' -exec rm -rf {} +

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:bookworm-slim
WORKDIR /app/
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/* .

# Run the web service on container startup.
CMD ["./passmApp"]
