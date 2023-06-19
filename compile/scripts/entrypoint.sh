#!/usr/bin/env bash
set -ex
info() {
    echo $(date) - [INFO] - "$*"
}
mkdir -p ${BIN}
info get redis-cli
go get -v github.com/go-redis/redis
info get cassandra-cli
go get -v github.com/gocql/gocql
info build get 
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BIN}/get ${SRC}/get.go
info build put 
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ${BIN}/put ${SRC}/put.go
info create ConfigMap ${BIN_CM}
kubectl -n ${POD_NAMESPACE} create configmap ${BIN_CM} --from-file ${BIN}
