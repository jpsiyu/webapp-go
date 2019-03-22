################## Server ####################
FROM golang:alpine AS server-builder

# Install git
RUN apk update && apk add --no-cache git

# Copy project
WORKDIR /go/src/app
COPY ./server ./server

# Fetch server dependencies.
WORKDIR /go/src/app/server
RUN go get -d -v

# Build server binary
WORKDIR /go/src/app
RUN go build -o appstart ./server
RUN pwd && ls -l

################## Client ####################
FROM node:carbon AS client-builder
WORKDIR /app
COPY ./package*.json ./
COPY ./client ./client
# Fetch client dependencies.
#RUN npm install --only-production
RUN npm install 

# Buld client dist
RUN npm run build
RUN pwd && ls -l

################## Combine in release ####################
FROM golang:alpine AS release
WORKDIR /go/src/app
COPY --from=server-builder /go/src/app/appstart ./
COPY --from=client-builder /app/dist ./dist
RUN pwd && ls -l

EXPOSE 8080

CMD ["./appstart"]