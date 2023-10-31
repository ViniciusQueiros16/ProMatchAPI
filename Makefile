.PHONY: build clean deploy

clean:
	go clean
	rm -rf ./bin


build: clean
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o cmd/match/promatch-create-match cmd/match/promatch-create-match/main.go
	


deploy: clean build
	sls deploy --verbose


start:
	sls offline 