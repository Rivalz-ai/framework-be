package constant

// share, keyshare, failover, exclusive
type ConsumerType int

// Consumer position
type ConsumerPosition int

type BatcherBuilderType int
type CompressionType int
type CompressionLevel int
type HashingScheme int
type MessageID string

// consumer type
const (
	// pulsar
	Shared ConsumerType = iota
	// pulsar	// kafka
	Exclusive
	// pulsar
	Failover
	// pulsar
	KeyShared
)

// Consumer position
const (
	//pulsar | kafka (Newest)
	//- consume old messages not ACK yet: if consumer group | subscription name is old
	//- consume new mesage: if consumer group | subscription name is new
	Earliest ConsumerPosition = iota
	//consume mesasges from begining:  pulsar | kafka ; kafka is oldest
	Latest
)

const (
	//default: (k1,v1), (k2,v2), (k1,v2) = after batch: [(k1,v1), (k2,v2), (k1,v2)]
	DefaultBatchBuilder BatcherBuilderType = iota
	//batch base on Key: (k1,v1), (k2,v2), (k1,v2) = after batch: [(k1,v1), (k1,v2)] ;  [(k2,v2)]
	KeyBasedBatchBuilder
)

const (
	NoCompression CompressionType = iota
	LZ4
	ZLib
	ZSTD
)

const (
	Default CompressionLevel = iota
	Faster
	Better
)

const (
	JavaStringHash HashingScheme = iota
	Murmur3_32Hash
)

var LatestMessageID MessageID = "LATEST"
var EarliestMessageID MessageID = "EARLIEST"
