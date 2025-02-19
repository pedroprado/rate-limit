docker-compose down
docker build -t firestore ./tests/firestore/.
docker build -t firestore_local -f ./tests/firestore/Dockerfile.local ./tests/firestore/.
docker-compose up --force-recreate -d --build --remove-orphans