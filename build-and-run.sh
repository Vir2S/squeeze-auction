#!/usr/bin/env bash

docker build -t ypeskov/auction:latest . --no-cache && 
docker run -p 3000:3000 ypeskov/auction:latest
 