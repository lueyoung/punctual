#!/bin/sh
info() {
    echo $(date) - [INFO] - "$*"
}
CMD="/usr/local/bin/init.py -i ${REGISTER} -n ${POD_NAMESPACE} $*"
info Running: \"$CMD\"
exec $CMD
