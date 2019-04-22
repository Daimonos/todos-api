# TODO Service

## About

This is a just-for-fun project to play with Microservices and Docker. All components are written in Go

## Requirements

This project requires the go-users-ms located at: https://github.com/Daimonos/go-user-ms

You can make the docker image locally for both this project and the user microservice by using the `make docker` command in each directory.

## Startup

1. Run the go-users-ms and go-todos-api microservices in docker
2. Set the environment variables defined below

## Environment Variables

1. TODOS_DBNAME = The filename for the Bolt databse that the TodoService will create. The docker container expects a /data folder to be mounted
2. TODOS_AUTHSERVICE = The url of the users microservice. It will be http://[Your Local IP Address]/validateToken
3. TODOS_PORT = The port you want the app to run on. Defaults to 80

## Usage

For now, you can use the app through Postman. You can send some registration & logins to the authentication service and use the token to call the todos service. It's pretty low-level but again, this is mostly just for fun :)
