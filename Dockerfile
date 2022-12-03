FROM ubuntu:kinetic-20221101

RUN apt update 
RUN apt upgrade -y
RUN apt install -y golang-go 
