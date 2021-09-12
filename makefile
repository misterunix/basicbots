docs: 
	gomarkdoc -u > documentation/basicbots-dev.md

linux:
	env GOOS=linux GOARCH=amd64 go build -o binaries/basicbots-linux64

windows:
	mkdir basicbots
	mkdir basicbots/robots
	mkdir basicbots/docs
	cp robots/* basicbots/robots
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o basicbots/basicbots64.exe
	zip windows64.zip basicbots/* basicbots/robots/* basicbots/docs/*
	rmdir /Q /S basicbots
  
