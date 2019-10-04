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
RUN mkdir -p ${LOG_DIR}
RUN mkdir -p ${OUTPUT_DIR}
RUN mkdir -p ${CONFIG_DIR}

# Environment Variables
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

# Copy go mod and sum files
# COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy application code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/wrangoler

# Declare volumes to mount
VOLUME ["${LOG_DIR}","${OUTPUT_DIR}","${CONFIG_DIR}"]

CMD ["./main"]