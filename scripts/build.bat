@echo=off
cd ..\cmd\git-tools
go build -o C:\Users\Gerard\go\bin\git-tools.exe
set GOARCH=amd64
set GOOS=linux
go build -o C:\Users\Gerard\go\bin\git-tools-linux
set GOARCH=amd64
set GOOS=darwin
go build -o C:\Users\Gerard\go\bin\git-tools-darwin
