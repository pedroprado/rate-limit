# If docker-compose is not installed, download it
docker-compose down
docker image rm notifications
docker image rm firestore-integration

docker-compose ps
if [ $? -ne 0 ]; then

  curl -L "https://github.com/docker/compose/releases/download/1.27.3/docker-compose-$(uname -s)-$(uname -m)" -o docker-compose
  if [ $? -ne 0 ]; then
      echo "Error getting docker compose"
      exit 1
  fi
  export PATH=$PATH:$PWD
  chmod +x docker-compose

fi

docker build -t notifications ../../
docker build -t firestore-integration ../firestore/.

docker-compose up --force-recreate -d --build --remove-orphans