FROM golang:1.18

# set working directory
WORKDIR /go/src/app

# Copye the source code
COPY . .

#EXPOSE the port
EXPOSE 8000

# Build the go app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]