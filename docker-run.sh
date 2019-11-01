PORT=8080
APP=testapi1
docker run --rm --name $APP -d -p $PORT:$PORT -e "PORT=$PORT" integrational/$APP:latest
