FROM golang:1.5.3

ADD . /image

WORKDIR /image

ENTRYPOINT ["go","run"]
