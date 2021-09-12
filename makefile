docs: 
	gomarkdoc -u > documentation/basicbots-dev.md

linux:
	env GOOS=linux GOARCH=amd64 go build -o binaries/basicbots-linux64
