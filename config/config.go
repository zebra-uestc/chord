package config

import "time"

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
const MainNodeAddressLocal string = "127.0.0.1:8001"

// MainNodeAddressMsg the mainNode's address for transmit PrevBlock
const MainNodeAddressMsg string = "127.0.0.1:8002"

// BathchTimeout the time for cutting a batch if there are not enough Msg
const BathchTimeout time.Duration = 2 * time.Second
