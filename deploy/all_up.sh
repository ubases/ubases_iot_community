#!/bin/bash

cd third
#chmod ugo+w vernemq
#chmod ugo+w redis
#chmod ugo+w haproxy
#chmod ugo-w ./mysql/conf/my.cnf
#mkdir ./mysql/data
#chmod ugo+w ./mysql/data
#docker-compose down --rmi all
docker-compose up -d

cd ../service
docker-compose down --rmi all
docker-compose up -d

cd ../simulator
docker-compose down --rmi all
docker-compose up -d

cd ../web
docker-compose down --rmi all
docker-compose up -d

