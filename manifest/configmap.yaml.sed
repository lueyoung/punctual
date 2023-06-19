apiVersion: v1
kind: ConfigMap
metadata:
  name: {{.env.cm}} 
  namespace: {{.namespace}}
data:
  DISCOVERY_NAME: {{.discovery.name}}
  DISCOVERY_NAMESPACE: {{.discovery.namespace}}
  OBJECT: {{.object}}
