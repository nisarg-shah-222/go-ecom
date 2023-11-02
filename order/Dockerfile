FROM golang:1.21.0
WORKDIR /order

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o order .