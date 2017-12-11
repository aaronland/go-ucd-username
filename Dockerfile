# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t ucd-username .
# docker run -p 6161:8080 -e HOST='0.0.0.0' ucd-username

FROM golang

ADD . /go-ucd-username

RUN cd /go-ucd-username; make bin

EXPOSE 8080

CMD /go-ucd-username/bin/ucd-usernamed -host ${HOST}
