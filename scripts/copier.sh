#!/usr/bin/env bash
info() {
    echo $(date) - [INFO] - "$*"
}

error() {
    echo $(date) - [ERROR] - "$*"
}

fail() {
    error "$@"
    exit 1
}
CMD="/usr/local/bin/svc -m ${OBJECT} -n ${DISCOVERY_NAMESPACE} -s ${DISCOVERY_NAME} -e \" \""
info Discovering: \"${CMD}\"
IPS=$(eval $CMD)
SEP=""
HOSTS=""
for IP in $IPS; do
    HOSTS+=$SEP
    HOSTS+=$IP
    SEP=','
done
info hosts: $HOSTS
CMD="/usr/local/bin/copy -r ${HOSTS} -c ${HOSTS} $*"
info Running: \"$CMD\"
eval $CMD
exit $?
