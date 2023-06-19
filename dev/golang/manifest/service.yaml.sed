apiVersion: v1
kind: Service
metadata:
  namespace: {{.namespace}}
  labels:
    component: {{.name}}
  name: {{.name}}
spec:
  type: NodePort 
  selector:
    component: {{.name}}
  ports:
    - port: 80 
      targetPort: {{.port}} 
      nodePort: 10080
      name: http 
