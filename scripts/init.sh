#!/usr/bin/env bash
set -ex
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
CMD="/usr/local/bin/svc -m ${OBJECT} -n ${DISCOVERY_NAMESPACE} -s ${DISCOVERY_NAME} -e \" \""
info Discovering: \"${CMD}\"
IPS=$(eval $CMD)
N=0
SEP=""
HOSTS=""
for IP in $IPS; do
    N=$[N+1]
    HOSTS+=$SEP
    HOSTS+="${IP}:6379 ${IP}:6380"
    SEP=' '
done
info Redis hosts info: ${HOSTS}
info Num of hosts: ${N}
#CMD="redis-trib.rb create --replicas ${N} ${HOSTS}"
CMD="echo yes | redis-cli --cluster create ${HOSTS} --cluster-replicas 1"
info Running: \"${CMD}\"
eval $CMD
exit $?
