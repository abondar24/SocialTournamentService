FROM library/mysql:latest

MAINTAINER Alex Bondar <abondar1992@gmail.com>

ENV  MYSQL_ROOT_PASSWORD=alex21
ADD db.sql /docker-entrypoint-initdb.d
EXPOSE 3306
