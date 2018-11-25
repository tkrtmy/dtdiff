FROM golang:1.11.2
WORKDIR /go/src/cmd
RUN go get -d -u github.com/tkrtmy/dtdiff
COPY . .
RUN go build -a -o dtdiff .

ENTRYPOINT ["./dtdiff"]
