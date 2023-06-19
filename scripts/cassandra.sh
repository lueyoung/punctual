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
CMD="/usr/local/bin/svc -m ${OBJECT} -n ${DISCOVERY_NAMESPACE} -s ${DISCOVERY_NAME} -e \" \""
info Discovering: \"${CMD}\"
IPS=$(eval $CMD)
SEEDS=""
SEP=""
for IP in $IPS; do
    SEEDS+=$SEP
    SEEDS+=$IP
    SEP=','
done

CONF=/etc/cassandra/cassandra.yaml
cat /tmp/cassandra.yaml > $CONF 

THIS_IP=$POD_IP
info cluster name: $CLUSTER_NAME
info IP: ${THIS_IP}
info pod namespace: ${POD_NAMESPACE}
info seeds: ${SEEDS}

# 1 cluster name
sed -i s?"{{.cluster.name}}"?"${CLUSTER_NAME}"?g $CONF

# 2 seed
sed -i s?"{{.seeds}}"?"$SEEDS"?g $CONF

# 3 address
sed -i s?"{{.pod.ip}}"?"${POD_IP}"?g $CONF

/usr/sbin/cassandra -R -f
