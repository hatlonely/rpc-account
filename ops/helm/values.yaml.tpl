name: "${NAME}"
replicaCount: "${REPLICA_COUNT}"

image:
  repository: "${REGISTRY_ENDPOINT}/${REGISTRY_NAMESPACE}/${IMAGE_REPOSITORY}"
  tag: "${IMAGE_TAG}"
  pullPolicy: Always
  pullSecret: "${PULL_SECRET_NAME}"

ingress:
  enable: true
  host: "${INGRESS_HOST}"
  secretName: "${SECRET_NAME}"

config:
  base.json: |
    {
      "Decoder": {
        "type": "Json"
      },
      "Provider": {
        "Type": "Local",
        "Options": {
          "Filename": "app.json"
        }
      }
    }
