FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN ls -la
RUN make go-api-server-build

CMD ["./server"]
