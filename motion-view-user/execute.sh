#!/bin/bash

docker run --rm \
	-v "$PWD":/var/task \
	-e DB_USER='musicpost_dist' -e DB_PASSWORD='LoPb3MUJHCXMTtdymHkNXWB8crOdwRBz' -e DB_NAME='rightr' -e DB_PORT=3362 -e DB_HOST='db.writer.dev.musicpost.com.br' \
	-e AWS_ACCESS_KEY_ID='AKIA3OZCVV4342LFW7NR' -e AWS_SECRET_ACCESS_KEY='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' -e AWS_DEFAULT_REGION='us-east-2' \
	-e ENV='dev' -e CACHE_HOST='18.116.179.147' -e CACHE_DB='1' -e CACHE_PORT='6379' -e CACHE_PASSWORD='aB5uzjw6S6So0tD344UuoM5xYXhnGfTz' -e MAX_RECORDS_PAGE=50\
	lambci/lambda:go1.x app \
	"$( cat event.json )"
