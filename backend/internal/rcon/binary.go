package rcon

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type binaryOperator struct {
	buf *bytes.Buffer
	err error
}

func (b *binaryOperator) Write(v any) {
	err := binary.Write(b.buf, binary.LittleEndian, v)
	if err != nil {
		b.err = fmt.Errorf("fail to write %w", err)
	}
}

func (b *binaryOperator) Read(v any) {
	err := binary.Read(b.buf, binary.LittleEndian, v)
	if err != nil {
		b.err = fmt.Errorf("fail to read %w", err)
	}
}
