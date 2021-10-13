all : build

docs:
	gomarkdoc -u > documentation/basicbots-dev.md

build:
	mkdir -p basicbots
	mkdir -p basicbots/docs
	mkdir -p basicbots/robots
	mkdir -p basicbots/binaries
	cp documentation/BASIC.md basicbots/docs/
	cp documentation/BASICBOTS.md basicbots/docs/
	cp robots/* basicbots/robots
	env GOOD=osx GOARCH=amd64 go build -o basicbots/binaries/basicbots
	env GOOS=linux GOARCH=amd64 go build -o basicbots/binaries/basicbots64-linux
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o basicbots/binaries/basicbots64-windows.exe
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o basicbots/binaries/basicbots64-darwin
	tar cvfz latest/linux64-0.0.2b.tgz basicbots/docs/ basicbots/robots/ basicbots/binaries/basicbots64-linux
	tar cvfz latest/darwin64-0.0.2b.tgz basicbots/docs/ basicbots/robots/ basicbots/binaries/basicbots64-darwin
	zip latest/windows64-0.0.2b.zip basicbots/robots/* basicbots/docs/* basicbots/binaries/basicbots64-windows.exe
	rm -rf basicbots
