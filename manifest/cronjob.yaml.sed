apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: {{.name6}} 
  namespace: {{.namespace}} 
  labels:
    component: {{.name6}}
    {{.labels.key}}: {{.labels.value}}
spec:
  schedule: "{{.schedule}}"
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: {{.service.account}}
          restartPolicy: OnFailure
          containers:
            # copier
            - name: {{.name6}}
              image: {{.image6}}
              imagePullPolicy: {{.image.pull.policy2}}
              command:
                - /usr/local/bin/copier.sh
              env:
                - name: POD_IP
                  valueFrom:
                    fieldRef:
                      fieldPath: status.podIP
                - name: NODE_NAME
                  valueFrom:
                    fieldRef:
                      fieldPath: spec.nodeName
                - name: POD_NAME
                  valueFrom:
                    fieldRef:
                      fieldPath: metadata.name
                - name: POD_NAMESPACE
                  valueFrom:
                    fieldRef:
                      fieldPath: metadata.namespace
                - name: HOST_IP
                  valueFrom:
                    fieldRef:
                      fieldPath: status.hostIP
              envFrom:
                - configMapRef:
                    name: {{.env.cm}}
              volumeMounts:
                - name: host-time
                  mountPath: /etc/localtime
                  readOnly: true
                - name: runable
                  mountPath: /usr/local/bin/copier.sh
                  subPath: copier.sh
                  readOnly: true
          volumes:
            - name: host-time
              hostPath:
                path: /etc/localtime
            - name: runable
              configMap:
                name: {{.scripts.cm}}
                defaultMode: 0755
