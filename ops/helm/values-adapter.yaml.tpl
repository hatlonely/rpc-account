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
          "Filename": "config/app.json"
        }
      }
    }
  app.json: |
    {
      "grpcGateway": {
        "httpPort": 80,
        "grpcPort": 6080,
        "exitTimeout": "20s",
        "validators": [
          "Default"
        ],
        "usePascalNameLogKey": false,
        "usePascalNameErrKey": false,
        "marshalUseProtoNames": true,
        "marshalEmitUnpopulated": false,
        "unmarshalDiscardUnknown": true,
        "enableTrace": true,
        "enableMetric": true,
        "enablePprof": true,
        "jaeger": {
          "serviceName": "rpc-account",
          "sampler": {
            "type": "const",
            "param": 1,
            "samplingServerURL": "http://jaeger-agent.monitoring:5778/sampling"
          },
          "reporter": {
            "logSpans": false,
            "localAgentHostPort": "jaeger-agent.monitoring:6831"
          }
        }
      },
      "service": {
        "accountExpiration": "5m",
        "captchaExpiration": "30m"
      },
      "redis": {
        "redis": {
          "addr": "${REDIS_ADDRESS}",
          "password": "${REDIS_PASSWORD}",
          "dialTimeout": "200ms",
          "readTimeout": "200ms",
          "writeTimeout": "200ms",
          "maxRetries": 3,
          "poolSize": 20,
          "db": 0
        },
        "wrapper": {
          "name": "accountdb",
          "enableTrace": true,
          "enableMetric": true,
        },
        "retry": {
          "attempt": 3,
          "delay": "1s",
          "lastErrorOnly": true,
          "delayType": "BackOff"
        }
      },
      "mysql": {
        "gorm": {
          "username": "${MYSQL_USERNAME}",
          "password": "${MYSQL_PASSWORD}",
          "database": "${MYSQL_DATABASE}",
          "host": "${MYSQL_ENDPOINT}",
          "port": 3306,
          "connMaxLifeTime": "60s",
          "maxIdleConns": 10,
          "maxOpenConns": 20
        },
        "wrapper": {
          "name": "accountdb",
          "enableTrace": true,
          "enableMetric": true,
        },
        "retry": {
          "attempt": 3,
          "delay": "1s",
          "lastErrorOnly": true,
          "delayType": "BackOff"
        }
      },
      "email": {
        "from": "${EMAIL_FROM}",
        "password": "${EMAIL_PASSWORD}",
        "server": "${EMAIL_SERVER}",
        "port": ${EMAIL_PORT}
      },
      "logger": {
        "grpc": {
          "level": "Info",
          "writers": [{
            "type": "RotateFile",
            "options": {
              "filename": "log/${NAME}.rpc",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
              }
            }
          }, {
            "type": "ElasticSearch",
            "options": {
              "index": "grpc",
              "idField": "requestID",
              "timeout": "200ms",
              "msgChanLen": 200,
              "workerNum": 2,
              "es": {
                "es": {
                  "uri": "${ELASTICSEARCH_ENDPOINT}",
                  "username": "elastic",
                  "password": "${ELASTICSEARCH_PASSWORD}"
                },
                "retry": {
                  "attempt": 3,
                  "delay": "1s",
                  "lastErrorOnly": true,
                  "delayType": "BackOff"
                }
              }
            }
          }]
        },
        "info": {
          "level": "Info",
          "writers": [{
            "type": "RotateFile",
            "options": {
              "filename": "log/${NAME}.rpc",
              "maxAge": "24h",
              "formatter": {
                "type": "Json"
              }
            }
          }]
        }
      }
    }