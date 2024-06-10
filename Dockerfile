FROM alpine:latest

WORKDIR /app

COPY . ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]