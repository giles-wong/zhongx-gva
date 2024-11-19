#!/bin/bash

# 设置环境变量
export GOOS=linux
export GOARCH=amd64

# 构建应用程序
go build -o myapp-linux-amd64

# 检查构建结果
if [ $? -ne 0 ]; then
    echo "Build failed."
    exit 1
fi

# 查看文件信息
file myapp-linux-amd64

# 复制文件到目标目录
cp myapp-linux-amd64 ./runGo

# 检查复制结果
if [ $? -ne 0 ]; then
    echo "Copy failed."
    exit 1
fi
# 删除构建文件
rm myapp-linux-amd64

echo "Build and copy completed successfully."