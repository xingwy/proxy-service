#!/bin/bash

# 编译指定的Go文件
build_go_file() {
    local go_file="$1"
    local output_dir="$2"
    
    echo "Building $go_file..."
   # go mod tidy
    go build -o $output_dir/service_app $go_file

    # 检查编译是否成功
    if [ $? -ne 0 ]; then
        echo "Build failed for $go_file"
        exit 1
    fi
}

# 检查是否提供了参数
if [ $# -eq 0 ]; then
    echo "No arguments provided. Usage: $0 [api]"
    exit 1
fi

# 根据参数值确定Go文件和目标目录
if [ "$1" == "api" ]; then
    go_file="./api/build/main.go"
    target_dir="./api/build"
else
    echo "Invalid argument. Usage: $0 [api]"
    exit 1
fi

# 确保目标目录存在
mkdir -p $target_dir

# 编译Go文件
build_go_file $go_file $target_dir

# 赋权限
chmod +x $target_dir/service_app

# 进入目标目录并启动服务
cd $target_dir
./service_app
