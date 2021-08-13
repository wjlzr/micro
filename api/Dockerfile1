FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS windows
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/micro
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/micro api/api-v1.go


FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app.
COPY --from=builder /app/micro /app/micro
COPY --from=builder /app/etc /app/etc

CMD ["./micro", "-f", "etc/api-v1.yaml"]
