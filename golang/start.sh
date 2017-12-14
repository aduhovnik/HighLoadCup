chown -R mysql:mysql /var/lib/mysql
service mysql start
/etc/init.d/mysql start
mysql -u root -p'root' -e 'CREATE DATABASE golang'
mysql -u root -p'root' golang < go/src/golang/dump.sql
mysql -u root -p'root' -e 'set global max_connections = 5001'
unzip /tmp/data/data.zip
./go/bin/golang