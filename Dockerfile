#Put the golang image as image's base
FROM golang

# 
LABEL maintainer="ericghoubiguian@live.fr"

#Define the Dockerfile argument 'sample_id' to specify the whished sample identifier
ARG sample_id

#Define the environment variable 'number' to take the 'sample_id' argument value and put it in the bottom CMD instruction
ENV number=$sample_id

#Copy all the files and directories in the newly directory weathermodule
COPY . /weathermodule

#Edit the environment variable value GOPATH for the weathermodule directory
ENV GOPATH /weathermodule

#Change work directory for the samples directory of weathermodule project
WORKDIR /weathermodule/samples/samplesFolder

#Container instruction as command-line interface command: 'go run sample_${number}.go' if the sample exists...
CMD if [ "${number}" -lt 7 ] && [ "${number}" -gt 0 ]; then go run sample_${number}.go; else echo "Error: there is no sample_${number}.go example"; fi
