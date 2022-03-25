FROM golang:1.16-alpine AS build-stage
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build server.go

FROM alpine:3.9
WORKDIR /app
COPY --from=build-stage /app/server /app
EXPOSE 8001
CMD ["/app/server"]