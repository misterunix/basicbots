docs: 
	gomarkdoc -u > documentation/basicbots-dev.md

linux:
	mkdir basicbots
	mkdir basicbots/docs
	mkdir basicbots/robots
	cp robots/* basicbots/robots
	env GOOS=linux GOARCH=amd64 go build -o basicbots/basicbots64
	tar cvfz latest/linux64.tgz basicbots
	rm -rf basicbots
	
