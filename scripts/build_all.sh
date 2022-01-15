#!/bin/sh

# disable go modules
export GOPATH=""

# disable cgo
export CGO_ENABLED=0

set -e
set -x
go mod tidy
# linux
GOOS=linux GOARCH=amd64 go build -o release/linux/amd64/dns-client
GOOS=linux GOARCH=arm64 go build -o release/linux/arm64/dns-client
GOOS=linux GOARCH=arm   go build -o release/linux/arm/dns-client
GOOS=linux GOARCH=386   go build -o release/linux/386/dns-client

# windows
GOOS=windows GOARCH=amd64 go build -o release/windows/amd64/dns-client.exe
GOOS=windows GOARCH=386   go build -o release/windows/386/dns-client.exe

# darwin
GOOS=darwin GOARCH=amd64 go build -o release/darwin/amd64/dns-client

# freebsd
GOOS=freebsd GOARCH=amd64 go build -o release/freebsd/amd64/dns-client
GOOS=freebsd GOARCH=arm   go build -o release/freebsd/arm/dns-client
GOOS=freebsd GOARCH=386   go build -o release/freebsd/386/dns-client

# netbsd
GOOS=netbsd GOARCH=amd64 go build -o release/netbsd/amd64/dns-client
GOOS=netbsd GOARCH=arm   go build -o release/netbsd/arm/dns-client

# openbsd
GOOS=openbsd GOARCH=amd64 go build -o release/openbsd/amd64/dns-client
GOOS=openbsd GOARCH=arm   go build -o release/openbsd/arm/dns-client
GOOS=openbsd GOARCH=386   go build -o release/openbsd/386/dns-client

# dragonfly
GOOS=dragonfly GOARCH=amd64 go build -o release/dragonfly/amd64/dns-client

# solaris
GOOS=solaris GOARCH=amd64 go build -o release/solaris/amd64/dns-client
