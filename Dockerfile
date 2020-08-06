FROM golang:1.14.6

WORKDIR /src/hello-world

ENV TZ=America/Los_Angeles

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app.
RUN go build -o /app -v ./cmd/campwiz2

ENTRYPOINT ["/app"]
