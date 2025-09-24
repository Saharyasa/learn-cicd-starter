# syntax=docker/dockerfile:1

########################
# 1) Build the binary
########################
FROM golang:1.22 AS builder

WORKDIR /app

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o notely .

########################
# 2) Minimal runtime image
########################
FROM gcr.io/distroless/static-debian12 AS runner

# Copy the compiled binary
COPY --from=builder /app/notely /usr/bin/notely

# Cloud Run listens on 8080 by default
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/usr/bin/notely"]
