#!/bin/bash
docker run --name redis -d redis
docker run --port 4711:8080 --link redis:db -i -t sudoq/keva
