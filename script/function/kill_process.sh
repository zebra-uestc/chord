KillProcessForPort(){
    for port in $*
    do
        echo "check $port"
        pid=`lsof -i:$port | grep "$port (LISTEN)" | awk '{print $2}'`
        if [ -n "${pid}" ]
        then
            echo "kill $pid for port $port"
            kill -9 $pid
        fi
    done
}