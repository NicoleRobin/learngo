package header

import "sync"

var (
	RequestPool  sync.Pool
	ResponsePool sync.Pool
)

func init() {
	RequestPool = sync.Pool{
		New: func() any {
			return &RequestHeader{}
		},
	}
	ResponsePool = sync.Pool{
		New: func() any {
			return &ResponseHeader{}
		},
	}
}

// ResetHeader reset request header
func (r *RequestHeader) ResetHeader() {
	r.ID = 0
	r.Checksum = 0
	r.Method = ""
	r.CompressType = 0
	r.RequestLen = 0
}

// ResetHeader reset response header
func (r *ResponseHeader) ResetHeader() {
	r.Error = ""
	r.ID = 0
	r.CompressType = 0
	r.Checksum = 0
	r.ResponseLen = 0
}
