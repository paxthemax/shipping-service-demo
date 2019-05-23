# Dockerfile
# Author: Pavle Batuta (pavle.batuta@protonmail.com)

########################
# STAGE 1: RUN
########################

# Use alpine linux to reduce image size.
FROM alpine:latest

# Download missing runtime dependencies.
RUN apk --no-cache add ca-certificates

RUN mkdir /service
WORKDIR /service

COPY data/data.json .
COPY bin/cli .

# Run the service.
ENTRYPOINT ["./cli", "data.json"]
CMD ["./cli"]
