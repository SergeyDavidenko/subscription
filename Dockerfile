FROM golang:alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/SergeyDavidenko/subscription
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Get all dependencies
# Build the Go app
# CGO_ENABLED=0 
RUN apk --no-cache add ca-certificates build-base librdkafka-dev pkgconf 
RUN GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go 
######## Start a new stage from scratch #######
FROM alpine:latest  
RUN apk --no-cache add ca-certificates librdkafka-dev pkgconf
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/src/github.com/SergeyDavidenko/subscription/main .
COPY config.yaml .
# Expose port 8080 to the outside world
EXPOSE 8080
# Command to run the executable
CMD ["./main"] 
