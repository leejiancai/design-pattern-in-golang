package decorator

import (
	"bytes"
	"testing"
)

func TestByteSaveInput(t *testing.T) {
	want := []byte("Golang")
	bs := &ByteSaveInput{}
	bs.write(want, 0, 0)
	got, _ := bs.read(0, 0)
	if !(bytes.Equal(want, got)) {
		t.Fatalf("Expected %q, got %q", want, got)
	}

}

func TestBase64(t *testing.T) {
	// 这里演示的就是两个装饰器类，对原来的ByteSaveInput的功能进行了增强，并且装饰器之前还可以嵌套使用
	want := []byte("Golang")
	bs := &ByteSaveInput{}
	base64Encoder := NewBase64EncodeInputStream(bs)
	base64Encoder.write(want, 0, 0)
	got, _ := base64Encoder.read(0, 0)
	if bytes.Equal(want, got) {
		t.Fatalf("Expected %q, got %q", want, got)
	}

	base64Decoder := NewBase64DecodeInputStream(base64Encoder)
	got, _ = base64Decoder.read(0, 0)
	if !bytes.Equal(want, got) {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
