# Dockerfile
# Author: Pavle Batuta (pavle.batuta@protonmail.com)

########################
# STAGE 1: RUN
########################

# Use alpine to reduce image size
FROM alpine:latest

# Download missing runtime dependencies.
RUN apk --no-cache add ca-certificates

# Create the service working directory.
RUN mkdir /service
WORKDIR /service

# Copy the service binary from the builder image.
COPY bin/consignment .

# Run the service.
CMD ["./consignment"]