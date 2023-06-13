#!/bin/bash

FOLDER=$PWD

go mod init app

for folder in $( ls /go/src/ )
do
	if [ -f /go/src/$folder/$folder.go ]
	then
		echo "||-------------------------------module-$folder------------------------------------||"
		
		cd /go/src/$folder
		go mod init $folder
		cp -Rap /go/src/$folder /usr/local/go/src/$folder
		cd $FOLDER
	fi
done

if [ "$1" == "run-get" ]; then
	go get -t ./...
fi

echo "||-------------------------------Building $FOLDER------------------------------------||"

go vet .
go build -v