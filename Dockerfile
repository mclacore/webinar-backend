# Download base image
FROM golang:1.21 as builder

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY fizzbuzz.go .

# Compile the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o fizzbuzz fizzbuzz.go

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./fizzbuzz"]