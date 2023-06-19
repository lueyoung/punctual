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

[[ -z $TIME ]] && TIME=10

CMD="/usr/local/bin/monitor -i ${POD_IP}"
if [[ -n ${THRESHOLD} ]]; then
  CMD+=" -t ${THRESHOLD}"
fi
CMD+=" $*"
info Running: \"$CMD\"
while :; do
  eval $CMD
  sleep $TIME 
done
