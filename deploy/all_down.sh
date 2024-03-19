#!/bin/bash

cd third
#docker-compose down --rmi all
docker-compose down
cd ../service
docker-compose down --rmi all
cd ../simulator
docker-compose down --rmi all
cd ../web
docker-compose down --rmi all

