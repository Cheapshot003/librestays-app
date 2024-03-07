# Use the official Go image as a parent image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the application
RUN go build -o main ./cmd

# Command to run the executable
CMD ["./main"]
