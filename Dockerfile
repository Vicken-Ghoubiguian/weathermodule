#
FROM golang

#
LABEL maintainer="ericghoubiguian@live.fr"

#
ARG sample_id

#
ENV number=$sample_id

#
COPY . /weathermodule

#
ENV GOPATH /weathermodule

#
WORKDIR /weathermodule/samples/samplesFolder

#
CMD go run sample_${number}.go
