#!/bin/bash

# Define variables
#REPO_DIR="./"
CONTAINER_NAME="idr-app"

# Navigate to the repository directory
#cd $REPO_DIR

# Pull the latest code
git pull origin master

# Build the new Docker image
docker compose build app

# Stop the running container
docker stop $CONTAINER_NAME

# Remove the old container
docker rm $CONTAINER_NAME

# Start the new container
docker compose up -d app

echo "Container updated successfully!"
