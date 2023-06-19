apiVersion: v1
kind: Service
metadata:
  namespace: {{.namespace}}
  labels:
    {{.labels.key}}: {{.labels.value}}
  name: {{.svc1}} 
spec:
  clusterIP: None 
  selector:
    component: {{.name}}
  ports:
    - port: 6379
      targetPort: 6379 
      name: redis 
    - port: 9042 
      targetPort: 9042 
      name: cassandra 
---
apiVersion: v1
kind: Service
metadata:
  namespace: {{.namespace}}
  labels:
    {{.labels.key}}: {{.labels.value}}
  name: {{.svc2}}
spec:
  type: NodePort 
  #type: ClusterIP 
  selector:
    component: {{.name}}
  ports:
    - port: {{.port}} 
      targetPort: {{.port}} 
      nodePort: 2{{.port}} 
      name: http 
