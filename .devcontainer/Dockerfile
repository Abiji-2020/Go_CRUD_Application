# Use the official Go image (Debian-based)
FROM golang:latest

# Install Redis
RUN apt-get update && apt-get install -y redis-server && apt-get clean

# Expose Redis and application ports
EXPOSE 6379 8080

# Set the working directory
WORKDIR /workspace

# Start Redis and keep the container running
CMD ["sh", "-c", "redis-server --daemonize yes && tail -f /dev/null"]
