#!/usr/bin/env bash
info() {
    echo $(date) - [INFO] - "$*"
}
warn() {
    echo $(date) - [WARNING] - "$*"
}
error() {
    echo $(date) - [ERROR] - "$*"
}
debug() {
    echo $(date) - [DEBUG] - "$*"
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
CMD="/usr/local/bin/batch -i $HOSTS"
if [[ -n ${BATCH_SIZE} ]]; then
  CMD+=" -n ${BATCH_SIZE}"
fi
CMD+=" $*"
info Running: \"$CMD\"
exec $CMD
exit $?
