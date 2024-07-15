mkdir basicbots
mkdir basicbots\robots
mkdir basicbots\docs
copy robots\* basicbots\robots
env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o basicbots/basicbots64.exe
zip windows64.zip basicbots\* basicbots\robots\* basicbots\docs\*
rmdir /Q /S basicbots
copy windows64.zip latest\
