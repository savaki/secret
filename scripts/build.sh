#!/bin/bash

export GOPATH=${HOME}/vagrant
mkdir ${GOPATH}

export PATH=/usr/local/go/bin:${PATH}
cd /vagrant/secret-tool

# install the dependencies
sudo apt-get install -y mercurial
go get code.google.com/p/go.crypto/nacl/secretbox

# build the project
go build
