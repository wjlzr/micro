#!/bin/bash
## author by wjl
## wechat wjl695788976

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



