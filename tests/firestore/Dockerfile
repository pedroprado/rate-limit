FROM google/cloud-sdk:335.0.0-alpine

RUN apk add --update --no-cache openjdk8-jre
RUN gcloud components install cloud-firestore-emulator beta --quiet
COPY entrypoint.sh .

ENV PORT 8080
EXPOSE "$PORT"

ENV FIRESTORE_PROJECT_ID "local-project"

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]