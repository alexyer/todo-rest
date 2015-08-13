FROM scratch
MAINTAINER Olexander Yermakov "olexander.yermakov@gmail.com"

ADD ca-certificates.crt /etc/ssl/certs
ADD todo-rest /

EXPOSE 8080

CMD ["/todo-rest"]
