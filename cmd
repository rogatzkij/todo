
docker build -t todo_app .
docker container prune -f

docker run --name mysql_todo -e MYSQL_ROOT_PASSWORD=ioUTRA -p3306:3306 -d mariadb

docker run --name adminer_todo --link mysql_todo:db -d -p 8080:8080 adminer

docker run --name app_todo --link mysql_todo:db -d -p 80:80 todo_app