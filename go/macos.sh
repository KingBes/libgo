# 当前文件所在的目录
cd $(dirname $0)

# 判断架构是否为amd64
if [ "$(uname -m)" != "x86_64" ]; then
arch="arm64"
else
arch="amd64"
fi

# 编译
GOOS=darwin GOARCH=$arch \
CGO_ENABLED=1 \
go build -o libgo.a -buildmode=c-archive .
