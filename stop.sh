#!/bin/bash

# 查找运行在 8080 端口的进程
PID=$(lsof -t -i:8080)

if [ -z "$PID" ]; then
    echo "没有找到运行在 8080 端口的进程"
else
    echo "正在停止进程 $PID"
    kill -9 $PID
    echo "后端服务已停止"
fi 