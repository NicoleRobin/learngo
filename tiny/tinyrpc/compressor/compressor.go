package compressor

type CompressType int32

const (
	Raw CompressType = iota
	Gzip
	Snappy
	Zlib
)

var Compressors = map[CompressType]Compressor{
	Raw:    &RawCompressor{},
	Gzip:   &GzipCompressor{},
	Snappy: &SnappyCompressor{},
	Zlib:   &ZlibCompressor{},
}

type Compressor interface {
	Zip([]byte) ([]byte, error)
	Unzip([]byte) ([]byte, error)
}

type RawCompressor struct {
}

func (r *RawCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (r *RawCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}

type GzipCompressor struct {
}

func (r *GzipCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (r *GzipCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}

type SnappyCompressor struct {
}

func (r *SnappyCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (r *SnappyCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}

type ZlibCompressor struct {
}

func (r *ZlibCompressor) Zip(data []byte) ([]byte, error) {
	return nil, nil
}

func (r *ZlibCompressor) Unzip(data []byte) ([]byte, error) {
	return nil, nil
}
