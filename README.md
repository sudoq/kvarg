#Kvarg
A simple key-value service

##API
Send GET requests to the following endpoints to get and set values

|API endpoint|Description|
|------------|-----------|
|`/{key}/{value}`|Sets the value for provided key|
|`/{key}`|Gets the value for provided key|


##Running with docker-compose
```
docker-compose up -d
```

##Running with docker
```
docker run --name redis -d redis
docker run --port 4711:8080 --link redis:db sudoq/kvarg
```

##Example usage
```
curl http://localhost:4711/exampleKey/exampleValue
curl http://localhost:4711/exampleKey
```
