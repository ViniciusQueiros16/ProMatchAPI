#!/bin/bash

FOLDER=/go/src/app
FOLDER_GO=/go
FOLDER_LAMBDA_FUNCTION=$PWD

if [ "$1" != "run-get" ]; then
	cd ../../libs && ./setup.sh && cd $FOLDER_LAMBDA_FUNCTION
fi

docker run -v "$PWD":$FOLDER -v "$PWD/../go":$FOLDER_GO -w $FOLDER golang:1.15 ./install_libs.sh "$1"