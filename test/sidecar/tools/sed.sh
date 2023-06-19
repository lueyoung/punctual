#!/usr/bin/env bash
show_help () {
cat << USAGE
usage: $0 [ -m MANIFEST-PATH ] [ -t TARGET-TO-SED ] [ -v VALUE ]
    -m : Specify the path of manifest.
    -t : Specify the target to sed.
    -v : Specify the value.
USAGE
exit 0
}
N=$#
if [[ $[N%2] == 1 ]]; then
  if [[ $* == -h ]]; then
    show_help
  fi
  exit 0
fi 
# Get Opts
while getopts "hm:t:v:" opt; do # 选项后面的冒号表示该选项需要参数
    case "$opt" in
    h)  show_help
        ;;
    m)  MANIFEST=$OPTARG
        ;;
    t)  TARGET=$OPTARG
        ;;
    v)  VALUE=$OPTARG
        ;;
    ?)  # 当有不认识的选项的时候arg为?
        echo "unkonw argument"
        exit 1
        ;;
    esac
done
[[ -z $* ]] && show_help
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
chk_var () {
if [ -z "$2" ]; then
  error no input for \"$1\", try \"$0 -h\".
  sleep 3
  exit 1
fi
}
chk_var -m $MANIFEST 
chk_var -t $TARGET 
#echo $MANIFEST
#ls ${MANIFEST}
#info $TARGET
#info $VALUE
CMD=$(cat <<EOF
find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"${TARGET}"?"${VALUE}"?g
EOF
)
#info Running: \"$CMD\"
eval $CMD
exit $?
