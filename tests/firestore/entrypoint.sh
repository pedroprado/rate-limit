#!/bin/bash

set -xe

gcloud config set project "${FIRESTORE_PROJECT_ID}"
gcloud beta emulators firestore start --host-port="0.0.0.0:${PORT}"