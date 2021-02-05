# weatherModule

This is a Go package using OpenWeatherMap to get and manipulate current weather datas from a wished city in a wished country.

## Contents

1. [What is this project ?](#what_is_this_project)

2. [How to use it ?](#how_to_use_it)

3. [A little example](#a_little_example)

4. [Running with Docker](#running_with_docker)

	* [Pull image from DockerHub](#pull_image_from_dockerhub)

	* [Build image from the Dockerfile](#build_image_from_dockerfile)

5. [A few usefull links](#a_few_usefull_links)

6. [Conclusion](#conclusion)

<a name="what_is_this_project"></a>
## What is this project ?

This project is about the development of a Go package using OpenWeatherMap to get and manipulate current weather datas from a wished city in a wished country.
This repository is composed of one Go file named *weathermodule.go* which contain the *WeatherModule* main type to get, to encapsulate and to manage weather, one Dockerfile to build your own Docker image to run examples, one *LICENCE* file to define the license (here the GNU version 3.0 one), the [*src*](https://github.com/Vicken-Ghoubiguian/weathermodule/tree/master/src) directory which contains all defined classes used in the WeatherModule type definition and the tthe [*samples*](https://github.com/Vicken-Ghoubiguian/weathermodule/tree/master/samples) directory which contain all defined samples to test locally the *WeatherModule* type, all of the other types used to define weather and their functions.

<a name="how_to_use_it"></a>
## How to use it ?

<a name="a_little_example"></a>
## A little example

<a name="running_with_docker"></a>
## Running with Docker

You can test all available samples with Docker. That will permit you to test the Go package, to study it, and to show how to integrate it in your own project. There are 2 ways to run the samples via Docker and to study their integration in it: with the current project's Dockerfile or the official Docker images hosted in the Docker Hub web service.

<a name="pull_image_from_dockerhub"></a>
### Pull image from Docker Hub

A Docker image to running `weathermodule` samples is available on Docker Hub service, just [here](https://hub.docker.com/r/wicken/weathermodule).
To running it and finaly enjoy all samples of the API, please follow the instructions below:

```bash
# Pulling the 'weathermodule' from Docker hub in the current machine...
docker pull wicken/weathermodule:latest

# Running the 'weathermodule' Docker image as a Docker container to test all of the Go samples...
docker container run -it wicken/weathermodule:latest
```

Now enjoy !

<a name="build_image_from_dockerfile"></a>
### Build image from the Dockerfile

The users can also build their own Docker image to running `weathermodule` samples using the current `Dockerfile`.
To build, run it and finaly enjoy all samples of the API, please follow the instructions below:

```bash
# Cloning the the 'weathermodule' project from GitHub...
git clone https://github.com/Vicken-Ghoubiguian/weathermodule

# Placing in the 'weathermodule' GitHub project's folder...
cd weathermodule

# Creating the weathermodule Docker image with the value of the user's OpenWeatherMap API key passed into the the parameter 'user_api_key'...
docker build . -t weathermodule:latest --build-arg user_api_key="<wished_openWeatherMap_API_key>"

# Running the 'weathermodule' Docker image as a Docker container to test all of the Go samples...
docker container run -it weathermodule:latest
```

Now enjoy !

<a name="a_few_usefull_links"></a>
## A few usefull links

<a name="conclusion"></a>
## Conclusion

It has been an exciting project, easily usable by anyone in any other project written in Go language regardless of its size (large or small) and which has broadened my knowledge in Go language development as potentially every other beginner or confirmed developer.
