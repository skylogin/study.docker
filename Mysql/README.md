# build
docker image build -t example/mysql-data:latest .

# run (volume: mysql-data)
docker container run -d --name mysql-data example/mysql-data:latest

# run (mysql)
docker container run -d --rm --name mysql \
  -e "MYSQL_ALLOW_EMPTY_PASSWORD=yes" \
  -e "MYSQL_DATABASE=volume_test" \
  -e "MYSQL_USER=example" \
  -e "MYSQL_PASSWORD=example" \
  --volumes-from mysql-data \
  mysql:5.7
  
# connect ROOT
docker container exec -it mysql mysql -u root -p volume_test


# stop mysql
docker container stop mysql

# compose data file with tar
docker continer run -v ${PWD}:/tmp \
  --volumes-from mysql-data \
  busybox \
  tar zcvf /tmp/mysql-backup.tar.gz /var/lib/mysql
