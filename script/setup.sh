#!/bin/bash
# fabric要与chord在同一层级
# 参数说明：
#   up：启动order、deliver、main_node
#   msg: 启动boardcaster发送消息，后跟参数消息数量
#   down: 清理所有有关进程
KillProcessForPort(){
    NEED_PORT=(6666 7050 8001 8002 8003 8004)
    for port in "${NEED_PORT[@]}"
    do
        pid=`lsof -i:$port | grep "localhost:$port (LISTEN)" | awk '{print $2}'`
        if [ -n "${pid}" ]
        then
            echo "kill $pid for port $port"
            kill -9 $pid
        fi
    done
}

if [ "$1" == "up" ] || [ "$1" == "msg" ]
then
    CURRENT_DIR="$(pwd)"
    SCRIPT_DIR="$CURRENT_DIR/$(dirname $0)"
    CHORD_PATH="$SCRIPT_DIR/.."
    FABRIC_PATH="$CHORD_PATH/../fabric"
    if [ "$1" == "up" ]
    then
        KillProcessForPort

        mkdir -p $SCRIPT_DIR/tmp
        cd $FABRIC_PATH
        FABRIC_CFG_PATH="$FABRIC_PATH/sampleconfig"
        echo "配置文件路径FABRIC_CFG_PATH=$FABRIC_CFG_PATH"
        echo "启动Orderer..."
        go run cmd/orderer/main.go > $SCRIPT_DIR/tmp/Orderer.log 2>&1 &
        sleep 1s
        cd $CHORD_PATH
        echo "启动DHT-MainNode..."
        go run server/main.go > $SCRIPT_DIR/tmp/DHT-MainNode.log 2>&1 &
        cd $FABRIC_PATH
        echo "启动Deliver..."
        go run orderer/sample_clients/deliver_stdout/client.go -channelID system-channel -quiet > $SCRIPT_DIR/tmp/Deliver.log 2>&1 &
    else
        if [ $# -eq 2 ]
        then
            echo "启动Broadcaster..."
            cd $FABRIC_PATH
            go run orderer/sample_clients/broadcast_msg/client.go -channelID system-channel -messages $2
        else
            echo "请输入要发送的Msg数量"
        fi
    fi
fi

if [ "$1" == "down" ]
then
    echo "closing..."
    KillProcessForPort
    
    END_PROCESS_NAME=("cmd/orderer/main.go" "deliver_stdout/client.go" "broadcast_msg/client.go" "go run server/main.go")
    for name in "${END_PROCESS_NAME[@]}"
    do
        pid=`ps -ef | grep "$name" | grep -v grep | awk '{print $2}'`
        if [ -n "${pid}" ]
        then
            echo "kill $pid for process $name"
            kill -9 $pid
        fi
    done
    echo "close over"
fi