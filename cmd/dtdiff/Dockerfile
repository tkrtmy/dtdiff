FROM golang:1.11.2 AS builder
WORKDIR /go/src/cmd
RUN go get -d -u github.com/tkrtmy/dtdiff
COPY . .
RUN go build -a -o dtdiff .

FROM scratch
COPY --from=builder /go/src/cmd/dtdiff /dtdiff
ENTRYPOINT ["./dtdiff"]
