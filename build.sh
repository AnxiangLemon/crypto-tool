#!/bin/bash

# 常量定义
MAIN_FILE="./main.go"
OUTPUT_DIR="./bin"
PROGRAM_NAME="lgotools"  # 修改这里以设置程序名称

# 设置要编译的平台和架构 "darwin/amd64" "darwin/arm64" IOS编译失败
PLATFORMS=("linux/amd64" "linux/arm64" "windows/amd64" "darwin/amd64" "darwin/arm64")

# 确保输出目录存在
mkdir -p $OUTPUT_DIR

# 遍历每个平台并编译
for PLATFORM in "${PLATFORMS[@]}"
do
    # 分割平台和架构
    IFS="/" read -r -a SPLIT <<< "$PLATFORM"
    GOOS=${SPLIT[0]}
    GOARCH=${SPLIT[1]}

    # 设置输出文件的名称，针对Windows平台添加.exe扩展名
    OUTPUT_NAME="$OUTPUT_DIR/${PROGRAM_NAME}-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi

    # 编译
    echo "Compiling for $GOOS/$GOARCH ..."
  #  GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -ldflags "-s -w" -o "$OUTPUT_NAME" $MAIN_FILE
    GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -ldflags "-s -w" -o "$OUTPUT_NAME" $MAIN_FILE

    if [ $? -ne 0 ]; then
        echo "Error: Failed to build for $GOOS/$GOARCH"
        exit 1
    fi
done

echo "所有指定平台的构建已完成。二进制文件在 $OUTPUT_DIR 目录."
