FROM golang:1.12-alpine as builder

WORKDIR /go/src/github.com/buzhiyun/gocron

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories && apk update \
    && apk add --no-cache git ca-certificates make bash nodejs yarn

RUN make install-vue \
    && make build-vue \
    && make statik \
    && CGO_ENABLED=0  GOPROXY=https://goproxy.io make gocron

FROM alpine:3.7

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/' /etc/apk/repositories \
    && apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -g app app

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /app

COPY --from=builder /go/src/github.com/buzhiyun/gocron/bin/gocron .

RUN chown -R app:app ./

EXPOSE 5920

USER app

ENTRYPOINT ["/app/gocron", "web"]
