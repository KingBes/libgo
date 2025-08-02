# 当前文件所在的目录
cd $(dirname $0)

# 编译
go build -o libgo.a -buildmode=c-archive ./libgo.go
