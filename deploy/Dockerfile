FROM ubuntu:latest
RUN mkdir /todo

ADD ../src/main/public/ /todo/public
ADD ../src/main/templates/ /todo/templates
ADD ../src/main/SQL.conf/ /todo/templates

ADD ../main /todo

WORKDIR /todo
RUN chmod +x /todo/main
CMD ["/todo/main"]