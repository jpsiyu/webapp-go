FROM golang:alpine

WORKDIR /go/src/app
COPY ./package.json /go/src/app
COPY ./server /go/src/app/server
COPY ./dist /go/src/app/dist

RUN ls -alR

EXPOSE 80

CMD ["go", "run", "./server/main.go"]