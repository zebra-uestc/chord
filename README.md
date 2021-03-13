# Chord

## 配置说明

### Main主机

MainNode与Orderer跑在同一台主机上，启动（重启）需要输入
`./script/setup.sh up`
在Main主机上启动Orderer、MainDhtNode、Deliver。

关闭MainNode与Orderer需要输入
`./script/setup.sh down`。

注意：Main主机需要fabric的工程与chord在同一级目录下，启动前修改[config.go](config/config.go)中的
`LocalAddress`
为Main主机在子网的IP地址。

## DhtNode主机

可在一台启动多个DhtNode加入到MainNode的环中，启动（重启）需要输入
`./script/join_node.sh join`。

关闭当前机器上的DhtNode需要输入
`./script/setup.sh close`。

注意：启动前需要修改[join_node.sh](script/join_node.sh)中的
`ADDR`
为当前DhtNode主机在子网的IP地址，
`PORTS`与`NODEIDS`分别修改为要定义的DhtNode的监听端口和ID，其数量要一致，且所有DhtNode的ID不能一致，如果有Dht节点挂掉再重新加入，则需要重启MainNode与所有的DhtNode。

## Broadcast主机

在Main主机与DhtNode主机启动完成后，需要在Broadcast主机上输入
`./script/setup.sh msg [参数消息数量] [并发数] [消息大小(单位KB)]`。

注意：启动前修改[setup.sh](script/setup.sh)中的
`ORDERER_ADDRESS`
为Main主机的地址 + Orderer接受消息的监听端口。