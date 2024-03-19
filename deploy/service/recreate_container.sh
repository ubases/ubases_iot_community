container_id=`docker ps -aq --filter "ancestor=service-$1"`
echo "container id：$container_id"
docker stop $container_id
docker rm $container_id

image_id=`docker image ls -q -f "reference=service-$1"`
echo "image id：$image_id"
docker rmi $image_id

echo "docker-compose install"
docker-compose up -d --force-recreate $1
echo "success"