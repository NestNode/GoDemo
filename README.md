# GeDemo

## 编译

```bash
go env -w GOPROXY=https://goproxy.cn,direct # (可选) 解决国内网络问题
go mod tidy

go run main.go # 运行

go build -o ./bin/demo.exe # 生成可执行文件 (windows加exe扩展名, linux不加)
go build -o ./bin/demo     # 生成可执行文件 (windows加exe扩展名, linux不加)
```
