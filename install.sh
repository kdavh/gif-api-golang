#!/bin/sh

# I've been enjoying developing using Docker, so even though it's one unnecessary level for these particular requirements, I'm going to just install docker and run the project specified docker container here.
# TODO: don't do everything in root

apt-get update
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
apt-add-repository 'deb https://apt.dockerproject.org/repo ubuntu-xenial main'
apt-get update
apt-get install -y docker-engine
