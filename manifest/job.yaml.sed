apiVersion: batch/v1
kind: Job
metadata:
  name: {{.name0}} 
  namespace: {{.namespace}} 
  labels:
    component: {{.name0}}
    {{.labels.key}}: {{.labels.value}}
spec:
  template:
    metadata:
      labels:
        component: {{.name0}}
        {{.labels.key}}: {{.labels.value}}
    spec:
      serviceAccountName: {{.service.account}}
      initContainers:
        - name: sleeper 
          image: busybox:latest
          imagePullPolicy: {{.image.pull.policy}}
          command: ['sh', '-c', 'echo Waiting ... && sleep 30'] 
      containers:
        - name: {{.name0}} 
          image: {{.image0}} 
          command: ["/usr/local/bin/init.sh"]
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
          envFrom:
            - configMapRef:
                name: {{.env.cm}}
          volumeMounts:
            - name: host-time
              mountPath: /etc/localtime
              readOnly: true
            - name: runable 
              mountPath: /usr/local/bin/init.sh 
              subPath: init.sh
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
