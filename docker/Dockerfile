FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Mike Glass <mglass2@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Build Args
ARG APP_DIR=/app

# Create Log Directory
RUN mkdir -p ${APP_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${APP_DIR}/logs/app.log

# Copy application code
COPY ./app .

# Install go modules for dependency
RUN go get -u github.com/spf13/cobra/cobra

# Set vars to help compile the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Declare volumes to mount
VOLUME ["${APP_DIR}"]

# init with compile
ENTRYPOINT ["/app/scripts/init"]