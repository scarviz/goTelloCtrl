@echo off
SET DEF_GOOS=%GOOS%
SET DEF_GOARCH=%GOARCH%

::SET GOOS=linux
::SET GOARCH=amd64
::go build -o goAmazonVideo main.go
go build -o tello.exe main.go

SET GOOS=%DEF_GOOS%
SET GOARCH=%DEF_GOARCH%