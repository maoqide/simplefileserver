From golang:1.7-alpine

MAINTAINER MAOQI maoqidemail@gmail.com

WORKDIR /root/
EXPOSE 18080
ADD ./simplefileserver.go /root
RUN go build -o simplefileserver && mkdir -p "/root/servedir"
#VOLUME /root/servedir

ENTRYPOINT ["/root/simplefileserver"]
