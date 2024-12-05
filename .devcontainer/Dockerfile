# Base image
FROM golang:latest

# Install Redis server
RUN apt-get update && apt-get install -y redis-server && apt-get install sudo

# Set up workspace
WORKDIR /workspace
COPY . .

# Start Redis and keep it running in the background
RUN mkdir -p /var/run/redis
CMD service redis-server start && tail -f /dev/null
