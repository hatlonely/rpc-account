CREATE DATABASE IF NOT EXISTS ${MYSQL_DATABASE};
CREATE USER IF NOT EXISTS '${MYSQL_USERNAME}'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';
GRANT ALL PRIVILEGES ON ${MYSQL_DATABASE}.* TO '${MYSQL_USERNAME}'@'%';

USE ${MYSQL_DATABASE};
CREATE TABLE IF NOT EXISTS \`accounts\` (
  \`id\` bigint NOT NULL AUTO_INCREMENT,
  \`email\` varchar(64) DEFAULT NULL,
  \`phone\` varchar(64) DEFAULT NULL,
  \`name\` varchar(32) DEFAULT NULL,
  \`password\` varchar(32) DEFAULT NULL,
  \`birthday\` timestamp NULL DEFAULT '1970-01-02 00:00:00',
  \`gender\` int DEFAULT NULL,
  \`avatar\` varchar(512) DEFAULT NULL,
  PRIMARY KEY (\`id\`),
  UNIQUE KEY \`email_idx\` (\`email\`),
  UNIQUE KEY \`phone_idx\` (\`phone\`),
  UNIQUE KEY \`name_idx\` (\`name\`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;