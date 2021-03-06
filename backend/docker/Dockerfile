FROM golang:1.14.3-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build 
CMD ["./Goflexify"]

