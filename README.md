# todo

## Что это?
Это моя курсовая работа. Представляет из себя веб-сервис по ведению дел (todo list).

## И как это запустить?

Рабочий вариант находится по адресу [todotodo.ml](http://todotodo.ml)

Что бы нам запустить веб-сервис локально потребуются 3 docker контейнера

1. MariaDB
~~~bash
docker run --name mysql_todo -e MYSQL_ROOT_PASSWORD=password -d mariadb
~~~

3. Adminer

~~~bash
docker run --name adminer_todo --link mysql_todo:db -d -p 8080:8080 adminer
~~~

5. Сама todo'шка
~~~bash
docker build -t todo_app ./source/
docker run --name app_todo --link mysql_todo:db -d -p 80:8081 todo_app
~~~

Настройки подключения todo'шки храняться в файле `source/SQL.config` 

## А что внутри?
### Backend
Написан на golang. 
### Frontend
html\css + немного js. 

 `source/public` - статический контент (js, css, стили)
  `source/template` - шаблоны html, в которые будет подставленна вся нужная информация 


## А есть ли недостатки?
Не весь функционал доделан, проект все еще находится на стадии разработки
Основные проблемы:
* не проработана многопоточность, надо добавить WaitGroup и Mutex
* не настроенно https соединение
* нет выдачи сообщения ошибки о неправильном логине\пароле при входе и при регистрации
* нет возможности редактировать созданные дела
* изначально планироволось выдавать пользователям ачивки за сделанные задания
* из БД пока не удаляются не активные сессии

## Что стоило бы переделать?
* отдавать статический контент не средствами go, а поручить это apache или nginx (как и https соединение)
* каждый раз при обновлении информации на странице происходит перезагрузка страницы, это плохо, надо переделать на ajax
* внешний вид окна добовления новых дел
