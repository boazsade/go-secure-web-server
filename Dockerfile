FROM ubuntu:latest

ENV SERVER_PORT=8443

RUN apt-get update && apt-get -y install golang-go
RUN mkdir -p /opt/webserver
WORKDIR /opt/webserver
COPY server.go /opt/webserver
COPY .certs /opt/webserver
RUN go build /opt/webserver/server.go
COPY staticfiles /opt/webserver
EXPOSE $SERVER_PORT:8443
CMD [ "./server", "8443" ]
