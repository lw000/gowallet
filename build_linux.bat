set prjPath=%cd%
echo %prjPath%
cd ../../../
set GOPATH=%cd%
set GOARCH=amd64
set GOOS=linux
cd %prjPath%

rem go mod vendor

go build -a -v -ldflags="-s -w"