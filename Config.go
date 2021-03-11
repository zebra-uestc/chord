package dhtnode

import "time"

// Simply specified as number of messages for now, in the future
// we may want to allow this to be specified by size in bytes
const MAXMESSAGECOUNT uint32 = 100

// The byte count of the serialized messages in a batch cannot
// exceed this value.
const ABSOLUTEMAXBYTES uint32 = 10 * 1024 * 1024

// The byte count of the serialized messages in a batch should not
// exceed this value.
const PREFERREDMAXBYTES uint32 = 2 * 1024 * 1024

// the local mainNode's address
const MAINNODEADDRESSLOCAL string = "127.0.0.1:8001"

// the mainNode's address for transmit PrevBlock
const MAINNODEADDRESSMSG string = "127.0.0.1:8002"

const BATCHTIMEOUT time.Duration = 2 * time.Second
