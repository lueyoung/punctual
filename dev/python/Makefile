SHELL=/bin/bash
NAME=py-dev
NAMESPACE=database
MANIFEST=./manifest
CONF=./conf
SCRIPTS=./scripts
LABELS_VALUE=${NAME}
IMAGE_PULL_POLICY=IfNotPresent
IMAGE_PULL_POLICY2=Always
LOCAL_REGISTRY=punctual.reg.me/test
SCRIPTS_CM=${NAME}-scripts
CONF_CM=${NAME}-conf
SCHEDULE=*/15 * * * *
DISCOVERY=${NAMESPACE}-headless
TAG=v3
IMAGE=${LOCAL_REGISTRY}/${NAME}:${TAG}
ADMIN_CONF=/etc/kubernetes/admin.conf
ADMIN=admin-conf

all: build push deploy

build:
	@docker build -t ${IMAGE} .

push:
	@docker push ${IMAGE}

cp:
	@find ${MANIFEST} -type f -name "*.sed" | sed s?".sed"?""?g | xargs -I {} cp {}.sed {}

sed:
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.name}}"?"${NAME}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.namespace}}"?"${NAMESPACE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.port}}"?"${PORT}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.url}}"?"${URL}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.image}}"?"${IMAGE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.image.pull.policy}}"?"${IMAGE_PULL_POLICY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.image.pull.policy2}}"?"${IMAGE_PULL_POLICY2}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.labels.key}}"?"${LABELS_KEY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.labels.value}}"?"${LABELS_VALUE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.scripts.cm}}"?"${SCRIPTS_CM}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.conf.cm}}"?"${CONF_CM}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.proxy}}"?"${PROXY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.schedule}}"?"${SCHEDULE}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.discovery}}"?"${DISCOVERY}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.admin}}"?"${ADMIN}"?g
	@find ${MANIFEST} -type f -name "*.yaml" | xargs sed -i s?"{{.admin.conf}}"?"${ADMIN_CONF}"?g

deploy: deploy-main deploy-cm

deploy-main: OP=create
deploy-main: cp sed
	#@kubectl -n ${NAMESPACE} ${OP} configmap $(SCRIPTS_CM) --from-file ${SCRIPTS}/.
	@kubectl ${OP} -f ${MANIFEST}/.

deploy-cm: OP=create
deploy-cm:
	@kubectl -n ${NAMESPACE} ${OP} configmap ${ADMIN} --from-file=config=${ADMIN_CONF}

clean-cm: OP=delete
clean-cm:
	@kubectl -n ${NAMESPACE} ${OP} configmap ${ADMIN}

clean: clean-main clean-cm

clean-main: OP=delete
clean-main:
	#@kubectl -n ${NAMESPACE} ${OP} configmap $(SCRIPTS_CM)
	@kubectl ${OP} -f ${MANIFEST}/.
