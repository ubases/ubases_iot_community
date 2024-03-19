#!/bin/bash

cd service
docker-compose down --rmi all
docker-compose up -d


