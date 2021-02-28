FROM golang:1.16.0 AS builder
WORKDIR /go/src/work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o dtdiff ./cmd/dtdiff

FROM scratch
COPY --from=builder /go/src/work/dtdiff /dtdiff
ENTRYPOINT ["./dtdiff"]
