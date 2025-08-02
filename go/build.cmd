go build -o demo.dll -buildmode=c-shared .\main.go
go build -o demo.a -buildmode=c-archive .\main.go