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
    "enableTrace": false,
    "enableMetric": false,
    "enablePprof": false,
    "jaeger": {
      "serviceName": "rpc-account",
      "sampler": {
        "type": "const",
        "param": 1
      },
      "reporter": {
        "logSpans": false
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