linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/expand-env-amd64 

.PHONY: linux
