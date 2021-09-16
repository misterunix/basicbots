all : build

docs: 
	gomarkdoc -u > documentation/basicbots-dev.md

build:
	mkdir basicbots
	mkdir basicbots/docs
	mkdir basicbots/robots
	mkdir basicbots/binaries
	cp documentation/BASIC.md basicbots/docs/
	cp documentation/BASICBOTS.md basicbots/docs/
	cp robots/* basicbots/robots
	env GOOS=linux GOARCH=amd64 go build -o basicbots/binaries/basicbots64-linux
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o basicbots/binaries/basicbots64-windows.exe
	tar cvfz latest/linux64.tgz basicbots basicbots/docs/ basicbots/robots/ basicbots/binaries/basicbots64-linux
	zip latest/windows64.zip basicbots/robots/* basicbots/docs/* basicbots/binaries/basicbots64-windows.exe
	rm -rf basicbots