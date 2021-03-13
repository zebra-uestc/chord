ADDR="10.206.0.11"
PORTS=(9001 9002 9003)
NODEIDS=("1-1" "1-2" "1-3")

CURRENT_DIR="$(pwd)"
SCRIPT_DIR="$CURRENT_DIR/$(dirname $0)"
source $SCRIPT_DIR/function/kill_process.sh

if [ "$1" == "join" ]
then
    CHORD_PATH="$SCRIPT_DIR/.."
    LOG_PATH="$SCRIPT_DIR/tmp/sample_nodes"
    cd $CHORD_PATH
    mkdir -p $LOG_PATH
    NODENUM=${#PORTS[@]}
    if [ $NODENUM -ne ${#NODEIDS[@]} ]
    then
        echo "Port与ID数量要一致"
    else
        KillProcessForPort ${PORTS[*]}
        for(( i=0;i<$NODENUM;i++)) 
        do   
            echo "join node: go run dhtnode/samplenode/main.go -addr $ADDR:${PORTS[$i]} -id ${NODEIDS[$i]} > $LOG_PATH/node${NODEIDS[$i]}.log 2>&1 &"
            go run dhtnode/samplenode/main.go -addr $ADDR:${PORTS[$i]} -id ${NODEIDS[$i]} > $LOG_PATH/node${NODEIDS[$i]}.log 2>&1 &
        done;
    fi
else
    if [ "$1" == "close" ]
    then
        KillProcessForPort ${PORTS[*]}
    else
        echo "请输入参数（join、close）"
    fi
fi
