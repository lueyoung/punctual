#!/usr/bin/env bash
set -ex
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
[[ -z $TIME ]] && TIME=60
[[ -z $BIT ]] && BIT=6
#HOST_HASH=$(echo -n $HOST_IP | sha1sum | awk -F ' ' '{print $1}')
#HOST_HASH=$(echo -n $HOST_IP | md5sum | awk -F ' ' '{print $1}')
HOST_HASH=$(echo -n $HOST_IP | sha256sum | awk -F ' ' '{print $1}')
N=${#HOST_HASH}
SHORT=${HOST_HASH:$[N-${BIT}]:$N}
HOST_ID=$((16#${SHORT}))
info host hash: HOST_HASH
#HOST_ID=$((16#${HOST_HASH}))
CMD="/usr/local/bin/run -i ${REGISTER} -n ${POD_NAMESPACE} -1 ${POD_IP} -2 ${HOST_IP} -3 ${HOST_HASH} -4 ${HOST_ID} $*"
info Running: \"$CMD\"
exec $CMD
sleep 30
exit $? 
