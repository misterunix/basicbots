#!/bin/sh


# Build the project
echo "cleaning bin directory"
rm -rf bin/*

echo "compiling linux amd64"
GOOS=linux GOARCH=amd64 go build -o bin/basicbots-linux_amd64

echo "compiling linux arm64"
GOOS=linux GOARCH=arm64 go build -o bin/basicbots-linux_arm64

echo "compiling windows amd64"
GOOS=windows GOARCH=amd64 go build -o bin/basicbots-windows_amd64.exe

echo "compiling darwin amd64"
GOOS=darwin GOARCH=amd64 go build -o bin/basicbots-darwin_amd64

echo "compiling darwin arm64"
GOOS=darwin GOARCH=arm64 go build -o bin/basicbots-darwin_arm64

# this is where I have my basicbots binaries
# you can change this to your own path and machine type
# cp bin/basicbots-linux_amd64 ~/basicbots/bin/
