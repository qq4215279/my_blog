set GOPATH=D:\data\GoWorkSpace
: GOARCH 是目标处理器架构 支持 arm arm64 amd64 386 ppc64 ppc64le mips64 mips64le s390x
set GOARCH=amd64
: GOOS 是目标操作系统 支持 darwin freebsd linux windows android dragonfly netbsd openbsd plan9 solaris
set GOOS=linux
go build
