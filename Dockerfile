FROM golang:latest

WORKDIR /app

# Copy the source code
COPY . .


# download depedency
RUN go mod tidy
# Build the binary files
RUN go build -o main ./main.go
RUN go build -o bot ./bot/bot.go
RUN go build -o watcher ./watcher/watcher.go

# Set the command to run when the container starts
CMD ["./main", ".bot", ".watcher"]