@echo off

@REM gocov / gocov-html

go get github.com/axw/gocov/gocov
go get github.com/matm/gocov-html
go mod tidy

%GOBIN%\gocov test ./... | %GOBIN%\gocov-html > coverage.html
del %GOBIN%\gocov.exe /Q /S
del %GOBIN%\gocov-html.exe /Q /S


@REM codecov

go test ./... -coverprofile=coverage.out
curl.exe --progress-bar -Lo codecov.exe https://uploader.codecov.io/latest/windows/codecov.exe 

@REM  codecov.exe -t ${CODECOV_TOKEN}
codecov.exe -t %1
del codecov.exe /Q /S