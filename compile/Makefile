SHELL=/bin/bash
KUBECTL_BIN=/usr/local/bin/kubectl
ADMIN_CONF=/etc/kubernetes/admin.conf
NAME=compile
IMAGE=golang:latest
IMAGE_PULL_POLICY=IfNotPresent
MANIFEST=./manifest
SCRIPTS=./scripts
SRC=./src
LABELS_KEY=app
LABELS_VALUE=${NAME}
SCRIPTS_CM=${NAME}-scripts
SRC_CM=${NAME}-src
CLI_CM=${NAME}-conf

all: echo deploy 

echo:
	@echo ${BIN_CM} 
	@echo ${NAMESPACE} 

cp:
	@find ${MANIFEST} -type f -name "*.sed" | sed s?".sed"?""?g | xargs -I {} cp {}.sed {}

sed:
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.name}}"?"${NAME}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.namespace}}"?"${NAMESPACE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.port}}"?"${PORT}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.image}}"?"${IMAGE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.image.pull.policy}}"?"${IMAGE_PULL_POLICY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.kubectl.bin}}"?"${KUBECTL_BIN}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.bin.cm}}"?"${BIN_CM}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.labels.key}}"?"${LABELS_KEY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.labels.value}}"?"${LABELS_VALUE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.scripts.cm}}"?"${SCRIPTS_CM}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.src.cm}}"?"${SRC_CM}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.admin.conf}}"?"${ADMIN_CONF}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.cli.cm}}"?"${CLI_CM}"?g

deploy: export OP=create
deploy: cp sed
	@kubectl -n ${NAMESPACE} ${OP} configmap ${SRC_CM} --from-file=${SRC}/.
	@kubectl -n ${NAMESPACE} ${OP} configmap ${SCRIPTS_CM} --from-file=${SCRIPTS}/.
	@kubectl -n ${NAMESPACE} ${OP} configmap ${CLI_CM} --from-file=config=${ADMIN_CONF}
	@kubectl ${OP} -f ${MANIFEST}/job.yaml

clean: export OP=delete
clean:
	@kubectl -n ${NAMESPACE} ${OP} configmap ${SRC_CM}
	@kubectl -n ${NAMESPACE} ${OP} configmap ${SCRIPTS_CM}
	@kubectl -n ${NAMESPACE} ${OP} configmap ${CLI_CM}
	-@kubectl -n ${NAMESPACE} ${OP} all,job,cm -l ${LABELS_KEY}=${LABELS_VALUE} 
