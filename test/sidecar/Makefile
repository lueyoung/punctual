SHELL=/bin/bash
NAME=sidecar-test
NAMESPACE=database
MANIFEST=./manifest
CONF=./conf
SCRIPTS=./scripts
TOOLS=./tools
LABELS_VALUE=${NAME}
IMAGE_PULL_POLICY=IfNotPresent
IMAGE_PULL_POLICY2=Always
LOCAL_REGISTRY=punctual.reg.me/test
SCRIPTS_CM=${NAME}-scripts
CONF_CM=${NAME}-conf
TAG=v3
SERVICE_ACCOUNT=admin
PORT=8080

NAME1=test
IMAGE1=centos:7

NAME2=kubectl
IMAGE2=lachlanevenson/k8s-kubectl:v1.12.4

all: deploy

build:
	@docker build -t ${IMAGE} .

push:
	@docker push ${IMAGE}

cp:
	@find ${MANIFEST} -type f -name "*.sed" | sed s?".sed"?""?g | xargs -I {} cp {}.sed {}

sed:
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.name}} -v ${NAME}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.name1}} -v ${NAME1}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.name2}} -v ${NAME2}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.namespace}} -v ${NAMESPACE}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.port}} -v ${PORT}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.image1}} -v ${IMAGE1}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.image2}} -v ${IMAGE2}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.image.pull.policy}} -v ${IMAGE_PULL_POLICY}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.image.pull.policy2}} -v ${IMAGE_PULL_POLICY2}
	@${TOOLS}/sed.sh -m ${MANIFEST} -t {{.service.account}} -v ${SERVICE_ACCOUNT}

deploy: deploy-main

deploy-main: OP=create
deploy-main: cp sed
	@kubectl ${OP} -f ${MANIFEST}/.

clean: clean-main

clean-main: OP=delete
clean-main:
	@kubectl ${OP} -f ${MANIFEST}/.
