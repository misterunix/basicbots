all : build

docs:
	gomarkdoc -u > documentation/basicbots-dev.md

build:
	# mkdir -p basicbots
	mkdir -p basicbots/docs
	mkdir -p basicbots/robots
	mkdir -p basicbots/binaries
	cp documentation/BASIC.md basicbots/docs/
	cp documentation/BASICBOTS.md basicbots/docs/
	cp robots/* basicbots/robots
	env GOOS=linux GOARCH=amd64 go build -o basicbots/binaries/basicbots64-linux
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o basicbots/binaries/basicbots64-windows.exe
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o basicbots/binaries/basicbots64-darwin
	sha256sum basicbots/binaries/basicbots64-linux
	sha256sum basicbots/binaries/basicbots64-windows.exe
	sha256sum basicbots/binaries/basicbots64-darwin
	tar cvz latest/linux64-0.0.2c.tgz basicbots/docs/ basicbots/robots/ basicbots/binaries/basicbots64-linux
	tar cvz latest/darwin64-0.0.2c.tgz basicbots/docs/ basicbots/robots/ basicbots/binaries/basicbots64-darwin
	zip -q latest/windows64-0.0.2c.zip basicbots/robots/* basicbots/docs/* basicbots/binaries/basicbots64-windows.exe
	sha256sum latest/linux64-0.0.2c.tgz
	sha256sum latest/darwin64-0.0.2c.tgz
	sha256sum latest/windows64-0.0.2c.zip
	rm -rf basicbots
