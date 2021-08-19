#!/bin/bash
## author by wjl
## wechat wjl695788976

##### 删除none镜像
# shellcheck disable=SC2046
docker rmi $(docker images | grep "none" | awk '{print $3}')

#### nginx 服务
docker stop nginx
docker rm nginx
docker run -itd --net=host \
  -p 80:80 \
  -p 443:443 \
  --name nginx --restart=always -e TZ="Asia/Shanghai" \
  -v /docker/nginx/html:/usr/share/nginx/html:ro \
  -v /docker/nginx/conf.d:/etc/nginx/conf.d:ro \
  -v /docker/nginx/log:/var/log/nginx \
  nginx

#### mysql 服务
docker stop mysql
docker rm mysql
docker run --net=host --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v /docker/mysql:/var/lib/mysql -itd mysql:5.7

##### redis 服务
docker stop redis
docker rm redis
echo "Start Redis Service..."
docker run --net=host --name redis -itd \
  --publish 6379:6379 \
  --env 'REDIS_PASSWORD=123456' \
  --volume /docker/redis:/var/lib/redis \
  sameersbn/redis:latest

#### etcd 服务
docker stop etcd && \
docker rm etcd || true && \
docker run --net=host -itd \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/docker/etcd,destination=/etcd-data \
  --name etcd \
  quay.io/coreos/etcd \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new

#### jaeger服务
docker stop jaeger
docker rm jaeger
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:latest



