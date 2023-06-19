apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: {{.name}} 
  namespace: {{.namespace}} 
spec:
  schedule: "{{.schedule}}"
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: {{.service.account}}
          restartPolicy: OnFailure
          containers:
            - name: {{.name}} 
              image: {{.image}} 
              imagePullPolicy: {{.image.pull.policy2}}
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
                - name: BATCH_SIZE 
                  value: "100000"
              envFrom:
                - configMapRef:
                    name: {{.env.cm}}
              volumeMounts:
                - name: host-time
                  mountPath: /etc/localtime
                  readOnly: true
              volumeMounts:
                - name: host-time
                  mountPath: /etc/localtime
                  readOnly: true
                - name: runable
                  mountPath: /usr/local/bin/batch.sh
                  subPath: batch.sh
                  readOnly: true
          volumes:
            - name: host-time
              hostPath:
                path: /etc/localtime
            - name: runable
              configMap:
                name: {{.scripts.cm}}
                defaultMode: 0755
