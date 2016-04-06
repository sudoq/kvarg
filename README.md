#Kvarg
A simple key-value service that uses redis for storage.
Can be used as a simple interface to redis but also for testing service discovery, configuration management, etc.

##Usage
The following example sets a key and then retrieves it
```
curl -XPUT http://127.0.0.1:4711/mykey -d value="kvarg for breakfast"
curl http://127.0.0.1:4711/mykey
```

##Running with docker-compose
```
docker-compose up -d
```

##Running with docker
```
docker run --name redis -d redis
docker run --port 4711:8080 --link redis:db sudoq/kvarg
```
