#!/bin/bash
# fabric要与chord在同一层级
# 参数说明：
#   up：启动order、deliver、main_node
#   msg: 启动boardcaster发送消息，后跟参数消息数量（可选，默认数量为1）、并发数（可选，默认并发数为1）、与消息大小（可选，单位为KB，默认大小1KB）
#   down: 清理所有有关进程
ORDERER_ADDRESS="10.206.0.9:7050"

CURRENT_DIR="$(pwd)"
SCRIPT_DIR="$CURRENT_DIR/$(dirname $0)"
source $SCRIPT_DIR/function/kill_process.sh

ListenPort=(6666 7050 8001 8002 8003)
if [ "$1" == "up" ] || [ "$1" == "msg" ]
then
    CHORD_PATH="$SCRIPT_DIR/.."
    FABRIC_PATH="$CHORD_PATH/../fabric"

    export FABRIC_CFG_PATH="$FABRIC_PATH/sampleconfig"
    echo "配置文件路径FABRIC_CFG_PATH=$FABRIC_CFG_PATH"

    if [ "$1" == "up" ]
    then
        KillProcessForPort ${ListenPort[*]}

        mkdir -p $SCRIPT_DIR/tmp
        cd $FABRIC_PATH
        echo "启动Orderer..."
        echo "go run cmd/orderer/main.go > $SCRIPT_DIR/tmp/Orderer.log 2>&1 &"
        go run cmd/orderer/main.go > $SCRIPT_DIR/tmp/Orderer.log 2>&1 &
        sleep 1s
        cd $CHORD_PATH
        echo "启动DHT-MainNode..."
        echo "go run server/main.go > $SCRIPT_DIR/tmp/DHT-MainNode.log 2>&1 &"
        go run server/main.go > $SCRIPT_DIR/tmp/DHT-MainNode.log 2>&1 &
        cd $FABRIC_PATH
        echo "启动Deliver..."
        echo "go run orderer/sample_clients/deliver_stdout/client.go -channelID system-channel -quiet > $SCRIPT_DIR/tmp/Deliver.log 2>&1 &"
        go run orderer/sample_clients/deliver_stdout/client.go -channelID system-channel -quiet > $SCRIPT_DIR/tmp/Deliver.log 2>&1 &
    else
        ARGS="-server $ORDERER_ADDRESS -channelID system-channel"
        if [ $# -ge 2 ]
        then
            ARGS="$ARGS -messages $2"
            if [ $# -ge 3 ]
            then
                ARGS="$ARGS -goroutines $3"
                if [ $# -ge 4 ]
                then
                    ARGS="$ARGS -size $[$4*1024]"
                fi
            fi
        fi
        cd $FABRIC_PATH
        echo "go run orderer/sample_clients/broadcast_msg/client.go $ARGS"
        go run orderer/sample_clients/broadcast_msg/client.go $ARGS
    fi
elif [ "$1" == "down" ]
then
    echo "closing..."
    KillProcessForPort ${ListenPort[*]}
        
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
else
    echo "请输入参数（up、msg、down）"
fi
