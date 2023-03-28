package header

import "sync"

// ResponseHeader request header structure looks like:
// +--------------+---------+----------------+-------------+----------+
// | CompressType |    ID   |      Error     | ResponseLen | Checksum |
// +--------------+---------+----------------+-------------+----------+
// |    uint16    | uvarint | uvarint+string |    uvarint  |  uint32  |
// +--------------+---------+----------------+-------------+----------+
type ResponseHeader struct {
	sync.RWMutex
	CompressType CompressType
	ID           uint64
	Error        string
	ResponseLen  uint32
	Checksum     uint32
}

// Marshal will encode response header into a byte slice
func (r *ResponseHeader) Marshal() []byte {
	return []byte{}
}

// Unmarshal
func (r *ResponseHeader) Unmarshal() (*ResponseHeader, error) {
	return nil, nil
}
