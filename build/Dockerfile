FROM golang:1.15

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy source files to current Working directory
COPY . .

# Build the go app
RUN go build -o main ./web

# Expose port 8080 to outside world
EXPOSE 8080

# Default command to run the exectutable
CMD [ "./main" ]