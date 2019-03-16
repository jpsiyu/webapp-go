############################
# STEP 1 build executable binary
############################

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/src/app
COPY . .

# Fetch dependencies.
# Using go get.
WORKDIR /go/src/app/server
RUN go get -d -v

# Build the binary.
WORKDIR /go/src/app

# Using go mod with go 1.11
RUN go mod download

RUN go build server/main.go 
RUN pwd && ls -l

ENTRYPOINT ["/go/src/app/main"]
CMD ["go", "run", "./server/main.go"]