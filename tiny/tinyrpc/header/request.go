package header

import (
	"encoding/binary"
	"sync"
)

type CompressType uint16

type RequestHeader struct {
	sync.RWMutex
	CompressType CompressType
	Method       string // 方法名
	ID           uint64 // 请求ID
	RequestLen   uint32 // 请求体长度
	Checksum     uint32 // 请求体校验和，crc32算法
}

// Marshal will encode request header into a byte slice
func (r *RequestHeader) Marshal() []byte {
	r.RLock()
	defer r.RUnlock()

	idx := 0
	header := make([]byte, MaxHeaderSize+len(r.Method))
	// 写入uint16类型的压缩类型
	binary.LittleEndian.PutUint16(header[idx:], uint16(r.CompressType))
	idx += Uint16Size

	idx += writeString(header[idex:], r.Method)
	idx += binary.PutUvarint(header[idx:], r.ID)
	idx += binary.PutUvarint(header[idx:], uint64(r.RequestLen))

	binary.LittleEndian.PutUint32(header[idx:], r.Checksum)
	idx += Uint32Size
	return header[:idx]
}

// Unmarshal
func (r *RequestHeader) Unmashal(data []byte) (err error) {
	r.Lock()
	defer r.Unlock()
	if len(data) == 0 {
		return UnmarshalError
	}

	defer func() {
		if r := recover(); r != nil {
			err = UnmashalError
		}

	}()

	idx, size := 0, 0
	r.CompressType = CompressType(binary.LittleEndian.Uint16(data[idx:]))
	idx += Uint16Size

	r.Method, size = readString(data[idx:])
	idx += size

	r.ID, size = binary.Uvarint(data[idx:])
	idx += size

	length, size := binary.Uvarint(data[idx:])
	r.RequestLen = uint32(length)
	idx += size

	r.Checksum = binary.LittleEndian.Uint32(data[idx:])
	return nil
}

func readString(data []byte) (string, int) {
	idx := 0
	length, size := binary.Uvarint(data)
	idx += size
	str := string(data[idx : idx+int(length)])
	idx += len(str)
	return str, idx
}

func writeString(data []byte, str string) int {
	idx := 0
	idx += binary.PutUvarint(data, uint64(len(str)))
	copy(data[idx:], str)
	idx += len(str)
	return idx
}
