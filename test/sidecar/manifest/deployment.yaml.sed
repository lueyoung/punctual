kind: Deployment 
apiVersion: extensions/v1beta1
metadata:
  namespace: {{.namespace}} 
  name: {{.name}} 
spec:
  replicas: 1
  template:
    metadata:
      labels:
        component: {{.name}}
        addonmanager.kubernetes.io/mode: Reconcile
    spec:
      serviceAccountName: {{.service.account}}
      containers:
        # 1 centos 
        - name: {{.name1}}
          image: {{.image1}}
          imagePullPolicy: {{.image.pull.policy}}
          command:
            - tail
            #- /usr/local/bin/copier.sh
          args:
            - -f 
            - /dev/null
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
          volumeMounts:
            - name: host-time
              mountPath: /etc/localtime
              readOnly: true
        # 2 kubectl 
        - name: {{.name2}}
          image: {{.image2}}
          imagePullPolicy: {{.image.pull.policy}}
          args:
            - proxy 
            - --port 
            - "{{.port}}"
          volumeMounts:
            - name: host-time
              mountPath: /etc/localtime
              readOnly: true
      volumes:
        - name: host-time
          hostPath:
            path: /etc/localtime
