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
ENV presentationSamples "\033[0;36mAll available samples: sample_1.go, sample_2.go, sample_3.go, sample_4.go, sample_5.go, sample_6.go and sample.7.go\033[0m"
ENV moreInfos "\033[0;36mIf you want more informations about this API, If you want more informations about these api, you can check the project github repo at https://github.com/Vicken-Ghoubiguian/weathermodule.\033[0m"
ENV errorMessage "\033[31mError: sample not found\033[0m"

#Container instruction as command-line interface command: 'go run $wished_sample' if the sample exists...
CMD echo "${welcomeMessage}"; \
    echo "${presentationSamples}"; \
    echo "${moreInfos}"; \
    echo "\e[94m"; \
    read -p 'What is your whished sample ? ' wished_sample; \
    read -p 'What is your wished city ? ' wished_city; \
    if [ -f "$wished_sample" ]; then \
	if [ $wished_sample = "sample_5.go" ]; then \
		read -p 'What is your whished country ? ' wished_country; \
		read -p 'What is your whished temperature scale ("Celsius", "Fahrenheit" or "Kelvin") ? ' whished_temp_scale; \
		echo "\033[0m"; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --countryCode="$wished_country" --temperatureScale="$whished_temp_scale"; \
	elif [ $wished_sample = "sample_4.go" ]; then \
		read -p 'What is your whished country ? ' wished_country; \
		echo "\033[0m"; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --countryCode="$wished_country"; \
	elif [ $wished_sample = "sample_6.go" ]; then \
		read -p 'What is your whished time format ("DMYHMS", "YMDHMS", "MDYHMS", "Timestamp") ? ' wished_time_format; \
		read -p 'What is your whished separator (Any char or string you want) ? ' wished_separator; \
		echo "\033[0m"; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --timeFormat="$wished_time_format" --dateSeparator="$wished_separator"; \
	elif [ $wished_sample = "sample_7.go" ]; then \
		read -p 'What is your whished pressure scale ("HectoPascal", "Pascal", "Bar", "Atmosphere", "Torr" or "PoundsPerSquareInch") ? ' wished_press_scale; \
		echo "\033[0m"; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}" --pressureScale="$wished_press_scale"; \
	else \
		echo "\033[0m"; \
		go run $wished_sample --city="$wished_city" --apiKey="${api_owm_key}"; \
	fi \
    elif [ -z "$wished_sample" ]; then \
	echo "\033[0m"; \
	echo  ""; \
    else \
	echo ""; \
	echo "${errorMessage}"; \
    fi
