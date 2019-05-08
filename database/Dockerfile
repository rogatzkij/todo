FROM mysql
RUN mkdir /db
ADD ./DB.sql /docker-entrypoint-initdb.d/init.sql
RUN mysql
