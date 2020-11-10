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

#Edit environment variables to contain messages to display
ENV welcomeMessage "\033[0;36mWelcome to the Docker tester of the weathermodule API examples...\033[0m"
ENV moreInfos "\033[0;36mIf you want more informations about this API, If you want more informations about these api, you can check the project github repo at https://github.com/Vicken-Ghoubiguian/weathermodule.\033[0m"

#Container instruction as command-line interface command: 'go run $wished_sample' if the sample exists...
CMD echo "${welcomeMessage}"; echo "${moreInfos}"; echo ""; if [ -f "$wished_sample" ]; then go run $wished_sample; elif [ -z "$wished_sample" ]; then echo  ""; else echo "\033[31mError: there is no $wished_sample example\033[0m"; fi
