FROM golang:alpine AS builder

# Install git
RUN apk update && apk add --no-cache git

# Copy project
WORKDIR /go/src/app
COPY . .

# Fetch dependencies.
WORKDIR /go/src/app/server
RUN go get -d -v

# Build the binary.
WORKDIR /go/src/app
RUN go build -o appstart ./server
RUN pwd && ls -l

# Copy to rease
FROM builder AS release
WORKDIR /go/src/app
COPY --from=builder /go/src/app/appstart .
COPY --from=builder /go/src/app/dist ./dist
RUN pwd && ls -l

EXPOSE 8080

CMD ["./appstart"]