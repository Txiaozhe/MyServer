#! /bin/sh

linuxRelease=myserver.linux

rm -f $linuxRelease

GOOS=linux GOARCH=amd64 go build -o $linuxRelease main.go

scp $linuxRelease root@120.24.168.127:~/MyServer
