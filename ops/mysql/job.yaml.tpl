apiVersion: batch/v1
kind: Job
metadata:
  namespace: ${NAMESPACE}
  name: init-db-${NAME}
spec:
  template:
    metadata:
      namespace: ${NAMESPACE}
      name: init-db-${NAME}
    spec:
      imagePullSecrets:
        - name: ${PULL_SECRET_NAME}
      containers:
        - name: init-db-${NAME}
          image: ${REGISTRY_ENDPOINT}/${REGISTRY_NAMESPACE}/mysql:8.0.25-debian-10-r37
          command:
            - /bin/sh
            - -ec
            - |
              mysql -uroot -h"${MYSQL_HOST}" -p"${MYSQL_ROOT_PASSWORD}" -e "
                CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};
                CREATE USER IF NOT EXISTS '${MYSQL_USERNAME}'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';
                GRANT ALL PRIVILEGES ON ${MYSQL_DATABASE}.* TO '${MYSQL_USERNAME}'@'%';
              "
      restartPolicy: Never