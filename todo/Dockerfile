FROM golang:1.18 AS build
WORKDIR /build
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./todo .

FROM alpine:latest AS prod
WORKDIR /app
COPY --from=build /build/todo /app/todo
EXPOSE 8080
CMD ./todo
