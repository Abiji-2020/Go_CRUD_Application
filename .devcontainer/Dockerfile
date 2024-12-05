# Use the official Go image as the base
FROM golang:latest

# Install Redis
RUN apt-get update && apt-get install -y redis-server && apt-get clean

# Expose Redis and application ports
EXPOSE 6379 8080

# Set the working directory
WORKDIR /workspace

# Start Redis in the background and keep the container running
CMD service redis-server start && tail -f /dev/null
