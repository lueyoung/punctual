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
        # copier
        - name: {{.name}}
          image: {{.image}}
          imagePullPolicy: {{.image.pull.policy}}
          command:
            - tail
            #- /usr/local/bin/copier.sh
          args:
            - -f 
            - /dev/null
          env:
            - name: DISCOVERY 
              value: {{.discovery}}.{{.namespace}}
            - name: KUBECONFIG
              value: {{.admin.conf}}
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
      volumes:
        - name: host-time
          hostPath:
            path: /etc/localtime
