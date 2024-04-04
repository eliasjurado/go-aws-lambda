set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0
go build -tags lambda.norpc -o bootstrap main.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o myFunction.zip bootstrap