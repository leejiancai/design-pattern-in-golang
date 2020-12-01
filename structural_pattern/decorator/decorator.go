package decorator

import "encoding/base64"

type InputStream interface {
	read(off int64, length int64) ([]byte, error)
	write(data []byte, off int64, length int64) (int64, error)
}

type FilterInputStream struct {
	in InputStream
}

func(fi *FilterInputStream) read(off int64, length int64) ([]byte, error) {
	return fi.in.read(off, length)
}

func(fi *FilterInputStream)write(data []byte, off int64, length int64)  (int64, error) {
	return fi.in.write(data, off, length)
}

type Base64DecodeInputStream struct {
	FilterInputStream
}

func NewBase64DecodeInputStream(in InputStream) *Base64DecodeInputStream {
	return &Base64DecodeInputStream{
		FilterInputStream{
			in:in ,
		},
	}
}

// 实现base64 Decode的功能
func (rb *Base64DecodeInputStream)read(off int64, length int64) ([]byte, error){
	ret, err := rb.in.read(off, length)
	if err != nil {
		return nil, err
	}

	n := base64.StdEncoding.DecodedLen(len(ret))
	dst := make([]byte, n)
	_, err = base64.StdEncoding.Decode(dst, ret)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

type Base64EncodeInputStream struct {
	FilterInputStream
}

func NewBase64EncodeInputStream(in InputStream) *Base64EncodeInputStream {
	return &Base64EncodeInputStream{
		FilterInputStream{
			in:in ,
		},
	}
}

// 实现base64 Encode的功能
func (rb *Base64EncodeInputStream)write(data []byte, off int64, length int64)  (int64, error) {
	n := base64.StdEncoding.EncodedLen(len(data))
	dst := make([]byte, n)
	base64.StdEncoding.Encode(dst, data)
	return rb.in.write(dst, off, length)
}


// ByteSaveInput才是业务原有的业务逻辑。上述的Base64EncodeInputStream和Base64DecodeInputStream是装饰类，是对原有功能进行增强
type ByteSaveInput struct {
	copy []byte
}

func(bi *ByteSaveInput) read(off int64, length int64) ([]byte, error) {
	return bi.copy, nil
}

func(bi *ByteSaveInput)write(data []byte, off int64, length int64)  (int64, error) {
	bi.copy = data
	return int64(len(data)), nil
}


