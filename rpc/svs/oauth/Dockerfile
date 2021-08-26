# build stage
FROM golang:alpine as build-env

ENV CGO_ENABLED 0
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /src

RUN cd /src && \
    go mod download

COPY . /src

COPY rpc/svs/oauth/etc /app/etc
RUN go build -ldflags="-s -w" -o bin/oauth rpc/svs/oauth/oauth.go


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/oauth /app/bin/oauth
COPY --from=build-env /app/etc /app/etc

ENV TZ Asia/Shanghai

EXPOSE 8889

CMD ["/app/bin/oauth","-f", "etc/oauth.yaml"]
