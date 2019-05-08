FROM ubuntu:latest
RUN mkdir /todo
ADD ./ /todo
WORKDIR /todo
RUN chmod +x /todo/main
CMD ["/todo/main"]