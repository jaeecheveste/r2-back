FROM golang:1.18 as builder
WORKDIR /go/src/github.com/jaeecheveste/r2-back
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o r2-back ./cmd

FROM alpine:3.8
RUN apk --no-cache add ca-certificates tzdata
ENV SERVER_PORT=8080
ENV LOG_LEVEL=INFO
ENV PPROF_ENABLED=false
COPY --from=builder /go/src/github.com/jaeecheveste/r2-back .
COPY /docs/r2-back-openapi.yml docs/r2-back-openapi.yml

ENTRYPOINT ["./r2-back"]