build:
	GOOS=linux GOARCH=amd64 go build -o bin/main main.go

deploy:
	sls deploy --verbose

invoke:
	sls invoke local --function selector --path event.json

offline:
	sls offline start
