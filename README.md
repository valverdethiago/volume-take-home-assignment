# volume-take-home-assignment

## Description
This project aims to purpose a solution for the problem described [here](./STATEMENT.md) as a take home assignment 
for a Senior Golang Developer Position at [Volume](https://volume.finance/)

## Solution
This software was written in [golang](https://go.dev/) (version 1.19) and can be run locally using go cli or docker.
Please make sure you have anyone of these tools properly setup. 

## API
The API input and output payloads can be found [here](./API.md)

## Makefile
There's a [Makefile](./Makefile) with some useful task to build, test or run the project:
- **tests**: Run all the tests with coverage
- **run**: Runs a local server
- **build**: Builds the code 
- **docker-build**: Builds a docker image with the app 
- **docker-run**: Runs the image built on *docker-build* using the default port 8080