#Keva
A simple key-value service

##API
Send GET requests to the following end points to get and set values
|API endpoint|Example|Description|
|------------|-------|-----------|
|`/{key}/{value}`|`http://localhost:8080/foo/bar`|Sets the value for provided key|
|`/{key}`|`http://localhost:8080/foo`|Gets the value for provided key|


##Running with docker
```
docker run --name redis -d redis
docker run --port 4711:8080 --link redis:db -i -t sudoq/keva
```
