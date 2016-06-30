FROM golang:1.5.3

ADD . /duproprio

WORKDIR /duproprio

ENTRYPOINT ["go","run"]