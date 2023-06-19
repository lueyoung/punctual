kind: Deployment 
apiVersion: extensions/v1beta1
metadata:
  namespace: {{.namespace}} 
  name: {{.name6}}-test 
spec:
  replicas: 1
  template:
    metadata:
      labels:
        component: {{.name6}}
    spec:
      containers:
        # copier
        - name: {{.name6}}
          image: {{.image6}}
          imagePullPolicy: {{.image.pull.policy2}}
          command:
            - tail
            #- /usr/local/bin/copier.sh
          args:
            - -f 
            - /dev/null
          env:
            - name: DISCOVERY 
              value: {{.name}}.{{.namespace}}
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
