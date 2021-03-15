package config

import "time"

const LocalAddress string = "10.206.0.11"

// MaxMessageCount simply specified as number of messages for now, in the future
// we may want to allow this to be specified by size in bytes
const MaxMessageCount uint32 = 100

// AbsoluteMaxBytes The byte count of the serialized messages in a batch cannot
// exceed this value.
const AbsoluteMaxBytes uint32 = 10 * 1024 * 1024

// PreferredMaxBytes The byte count of the serialized messages in a batch should not
// exceed this value.
const PreferredMaxBytes uint32 = 2 * 1024 * 1024

// MainNodeAddressLocal the local mainNode's address
const MainNodeAddressLocal string = LocalAddress + ":8001"

// MainNodeAddressMsg the mainNode's address for transmit PrevBlock
const MainNodeAddressMsg string = ":8003"

// MainNodeAddressBlock the mainNode's address for transmit PrevBlock
const MainNodeAddressBlock string = LocalAddress + ":8002"

// BathchTimeout the time for cutting a batch if there are not enough Msg
const BathchTimeout time.Duration = 2 * time.Second

// OrdererAddress the orderer consensus dht address
const OrdererAddress string = "127.0.0.1:6666"

// GrpcTimeout grpc time out
const GrpcTimeout time.Duration = 2 * time.Second