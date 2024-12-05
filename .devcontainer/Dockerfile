# Use Alpine as the base image
FROM alpine:latest

# Install necessary packages: bash, curl, Redis, and Go
RUN apk update && apk add --no-cache bash curl redis go

# Expose Redis and application ports
EXPOSE 6379 8080

# Set the working directory
WORKDIR /workspace

# Start Redis and keep the container running
CMD ["sh", "-c", "redis-server --daemonize yes && tail -f /dev/null"]
