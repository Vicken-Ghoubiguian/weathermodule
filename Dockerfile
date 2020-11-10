#Put the golang image as image's base
FROM golang

# 
LABEL maintainer="ericghoubiguian@live.fr"

#Copy all the files and directories in the newly directory weathermodule
COPY . /weathermodule

#Edit the environment variable value GOPATH for the weathermodule directory
ENV GOPATH /weathermodule

#Change work directory for the samples directory of weathermodule project
WORKDIR /weathermodule/samples

#Container instruction as command-line interface command: 'go run $wished_sample' if the sample exists...
CMD if [ -f "$wished_sample" ]; then go run $wished_sample; elif [ -z "$wished_sample" ]; then echo  ""; else echo "\033[31mError: there is no $wished_sample example\033[0m"; fi
