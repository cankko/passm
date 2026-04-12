# Use the official Go image to create a binary.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.25-bookworm AS builder

# Create and change to the app directory.
WORKDIR /app

RUN apt-get update && apt-get install -y git

RUN git clone https://bitbucket.org/cankko/passm.git .

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
RUN go mod download

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o passm ./cmd/main.go

# remove all repository files except data, web folder and build binary file
RUN find . -maxdepth 1 ! -name 'data' ! -name 'web' ! -name 'passm' ! -name '.' -exec rm -rf {} +

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/passm .

# Optional: Einen nicht-privilegierten User nutzen (Sicherheit!)
RUN useradd -m appuser
USER appuser

# Run the web service on container startup.
CMD ["/app/passm"]
