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
if [ ! -x "$(command -v getent)" ]; then
  fail getent: no such file!
  exit 1
fi
WAIT=10
TRIES=100
i=0
while true; do
  if [[ $i -gt $TRIES ]]; then
    echo "=== Cannot resolve the DNS entry for $DISCOVERY. Has the service been created yet, and is SkyDNS functional?"
    echo "=== See http://kubernetes.io/v1.1/docs/admin/dns.html for more details on DNS integration."
    echo "=== Sleeping ${WAIT}s before pod exit."
    sleep $WAIT
    exit 0
  fi
  if ! getent hosts $DISCOVERY; then
    i=$[i+1]
    warn not get $DISCOVERY for $i
    sleep 3
  else
    break
  fi
done

INFO=''
while [[ -z $INFO ]]; do
  INFO=$(getent hosts $DISCOVERY)
done
debug info: $INFO
REG='^([0-9]{1,3}.){3}[0-9]{1,3}$'
HOSTS=''
SEP=''
for IP in $INFO; do
  if [[ $IP =~ $REG ]]; then
    HOSTS+=$SEP
    HOSTS+=$IP
    SEP=','
  fi
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
