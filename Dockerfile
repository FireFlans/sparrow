# Stage 1: Build stage
FROM golang:1.23-alpine AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o sparrow .

# Stage 2: Final stage
FROM scratch

WORKDIR /app

# Copy the builded app
COPY --from=build /app/sparrow .

COPY config /app/config
COPY playground /app/playground

EXPOSE 8080

ENTRYPOINT ["/app/sparrow"]