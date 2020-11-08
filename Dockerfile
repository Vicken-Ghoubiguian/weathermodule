#
FROM golang

#
LABEL maintainer="ericghoubiguian@live.fr"

#
COPY . /weathermodule

#
ENV GOPATH /weathermodule

#
WORKDIR /weathermodule/samples/samplesFolder

#
ENTRYPOINT ["go", "run", "sample_1.go"]
