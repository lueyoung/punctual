apiVersion: batch/v1
kind: Job
metadata:
  name: {{.name}} 
  namespace: {{.namespace}} 
  labels:
    component: {{.name}}
    {{.labels.key}}: {{.labels.value}}
spec:
  template:
    metadata:
      labels:
        component: {{.name}}
        {{.labels.key}}: {{.labels.value}}
    spec:
      containers:
        - name: {{.name}} 
          image: {{.image}} 
          command: ["/usr/local/bin/entrypoint.sh"]
          imagePullPolicy: {{.image.pull.policy}}
          env:
            - name: KUBECONFIG 
              value: {{.admin.conf}} 
            - name: BIN_CM 
              value: {{.bin.cm}} 
            - name: BIN 
              value: /workspace/bin
            - name: SRC 
              value: /workspace/src
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
          volumeMounts:
            - name: host-time
              mountPath: /etc/localtime
              readOnly: true
            - name: runable 
              mountPath: /usr/local/bin/entrypoint.sh 
              subPath: entrypoint.sh
              readOnly: true
            - mountPath: {{.admin.conf}} 
              name: admin-conf 
              readOnly: true
              subPath: config
            - mountPath: /workspace/src 
              name: src 
              readOnly: true
            - mountPath: /bin/kubectl
              name: kubectl-bin
              readOnly: true
      restartPolicy: Never
      volumes:
        - name: host-time
          hostPath:
            path: /etc/localtime
        - name: runable 
          configMap:
            name: {{.scripts.cm}}
            defaultMode: 0755
        - name: admin-conf
          configMap:
            name: {{.cli.cm}} 
        - name: src 
          configMap:
            name: {{.src.cm}} 
        - name: kubectl-bin
          hostPath:
            path: {{.kubectl.bin}} 
