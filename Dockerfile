#Put the golang image as image's base
FROM golang

# 
LABEL maintainer="ericghoubiguian@live.fr"

#Copy all the files and directories in the newly directory weathermodule
COPY . /weathermodule

#Edit the environment variable value GOPATH for the weathermodule directory
ENV GOPATH /weathermodule

#Define the Dockerfile argument 'user_api_key' to specify the user's api key
ARG user_api_key

#Define the environment variable 'api_owm_key' to take the 'user_api_key' argument value and put it to run the samples
ENV api_owm_key=$user_api_key

#Change work directory for the samples directory of weathermodule project
WORKDIR /weathermodule/samples

#Edit environment variables to contain messages to display
ENV welcomeMessage "\033[0;36mWelcome to the Docker tester of the weathermodule API examples...\033[0m"
ENV moreInfos "\033[0;36mIf you want more informations about this API, If you want more informations about these api, you can check the project github repo at https://github.com/Vicken-Ghoubiguian/weathermodule.\033[0m"
ENV errorMessage "\033[31mError: sample not found\033[0m"

#Container instruction as command-line interface command: 'go run $wished_sample' if the sample exists...
CMD echo "${welcomeMessage}"; \
    echo "${moreInfos}"; \
    echo ""; \
    read -p 'What is your whished sample ? ' wished_sample; \
    read -p 'What is your wished city ? ' wished_city; echo ""; \
    if [ -f "$wished_sample" ]; then \
	if [ $wished_sample = "sample_5.go" ]; then \
		read -p 'What is your whished country ? ' wished_country; \
		read -p 'What is your whished temperature scale ? ' whished_temp_scale; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --countryCode="$wished_country" --temperatureScale="$whished_temp_scale"; \
	elif [ $wished_sample = "sample_4.go" ]; then \
		read -p 'What is your whished country ? ' wished_country; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --countryCode="$wished_country"; \
	else \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}"; \
	fi \
    elif [ -z "$wished_sample" ]; then \
	echo  ""; \
    else echo "${errorMessage}"; fi
