apiVersion: v1
kind: Namespace
metadata:
  name: {{.namespace}}
  labels:
    {{.labels.key}}: {{.labels.value}}
