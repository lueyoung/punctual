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
    - port: {{.port}} 
      targetPort: {{.port}} 
      nodePort: 1{{.port}}
      name: http 
