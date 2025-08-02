@REM go build -o libgo.dll -buildmode=c-shared .\libgo.go
@REM 当前文件所在的目录
cd %~dp0
go build -o libgo.a -buildmode=c-archive .\libgo.go