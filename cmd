

docker run --name mysql_todo -e MYSQL_ROOT_PASSWORD=123456 -d mariadb

docker run --name adminer_todo --link mysql_todo:db -d -p 8081:8080 adminer

docker run --name app_todo --link mysql_todo:db -d -p 8080:8080 todo_app


CREATE USER 'non-root'@'db' IDENTIFIED BY '123456';
GRANT ALL PRIVILEGES ON * . * TO 'non-root'@'db';
flush privileges