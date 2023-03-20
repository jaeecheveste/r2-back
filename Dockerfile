FROM golang:1.18 as builder
WORKDIR /go/src/github.com/jaeecheveste/r2-back
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o r2-back ./cmd

FROM alpine:3.8
RUN apk --no-cache add ca-certificates tzdata
ENV MONGOURI=mongodb+srv://malomz:malomzPassword@cluster0.e5akf.mongodb.net/golangDB?retryWrites=true&w=majority
ENV SERVER_PORT=8080
ENV LOG_LEVEL=INFO
ENV PPROF_ENABLED=false
COPY --from=builder /go/src/github.com/jaeecheveste/r2-back .
ENTRYPOINT ["./r2-back"]