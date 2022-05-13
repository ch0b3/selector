build:
	gofmt -w .
	GOOS=linux GOARCH=amd64 go build -o bin/main main.go

deploy-dev: build
	sls deploy --verbose --stage dev

deploy-prod: build
	sls deploy --verbose --stage prod

invoke:
	sls invoke local --function selector --path event.json

offline:
	sls offline start
