FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Mike Glass <mglass2@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Build Args
ARG LOG_DIR=/app/logs
ARG OUTPUT_DIR=/app/output
ARG CONFIG_DIR=/app/configs

# Create Log Directory
RUN mkdir -p ${LOG_DIR} && \
    mkdir -p ${OUTPUT_DIR} && \
    mkdir -p ${CONFIG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

# Copy application code
COPY . .

# Install go modules for dependency
RUN go get -u github.com/spf13/cobra/cobra

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./

# Declare volumes to mount
VOLUME ["${LOG_DIR}","${OUTPUT_DIR}","${CONFIG_DIR}"]

CMD ["./main"]