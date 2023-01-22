FROM golang:alpine
WORKDIR /app

# Copy .env file
# COPY .env .

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Expose port 8080 to the host
EXPOSE 8080

# Run the binary
CMD ["./main"]