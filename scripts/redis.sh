#!/bin/bash
set -e
info() {
    echo $(date -d today +'%Y-%m-%d %H:%M:%S') - [INFO] - "$*"
}
warn() {
    echo $(date -d today +'%Y-%m-%d %H:%M:%S') - [WARNING] - "$*"
}
error() {
    echo $(date -d today +'%Y-%m-%d %H:%M:%S') - [ERROR] - "$*"
}
debug() {
    echo $(date -d today +'%Y-%m-%d %H:%M:%S') - [DEBUG] - "$*"
}

fail() {
    error "$@"
    exit 1
}

MAX_MEM=${MAX_MEM:-"75"}
CONF=/usr/local/etc/redis
FILE=${CONF}/redis.conf
mkdir -p ${CONF} 
[[ -f ${FILE} ]] || touch ${FILE} 
cat /tmp/redis.conf > ${FILE} 
info IP: ${POD_IP}
info pod namespace: ${POD_NAMESPACE}
info port: ${PORT}
info max mem: ${MAX_MEM} %
MemTotal=$(cat /proc/meminfo | grep MemTotal: | awk -F ' ' '{print $2}')
BYTES=$[${MemTotal}*1024*${MAX_MEM}/100]
info max mem in trem of bytes: ${BYTES}

# 1 ip 
sed -i s?"{{.ip}}"?"${POD_IP}"?g ${FILE} 

# 2 port 
sed -i s?"{{.port}}"?"${PORT}"?g ${FILE} 

# 2 port 
sed -i s?"{{.bytes}}"?"${BYTES}"?g ${FILE} 

exec "$@"
