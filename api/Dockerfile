# build stage
FROM golang:alpine as build-env
#RUN apk add build-base
#LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn"

WORKDIR /src

COPY go.mod go.mod
COPY go.sum go.sum

RUN cd /src && go mod download

COPY . /src
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o bin/api api/api-v1.go


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/api /app/bin/api
COPY --from=build-env /app/etc /app/etc

ENV TZ Asia/Shanghai

EXPOSE 8888

CMD ["/app/bin/api","-f", "etc/api-v1.yaml"]
