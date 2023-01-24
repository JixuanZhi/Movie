FROM ubuntu:kinetic-20221101

RUN apt update 
RUN apt upgrade -y
RUN apt install -y golang-go

EXPOSE 80

VOLUME /go_source

WORKDIR /go_source 

CMD go run server.go
#CMD /bin/bash
