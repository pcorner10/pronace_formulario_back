# Build stage
FROM golang:1.21-rc-alpine3.17 AS build-env
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build-env /app/main .

# Expose works as a documentation for the port
EXPOSE 8080 
CMD ["/app/main"]