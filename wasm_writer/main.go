package wasmwriter

import "bytes"

type Wasm_Writer struct {
	buffer bytes.Buffer
}

func NewWasmWriter() *Wasm_Writer {
	return &Wasm_Writer{}
}

func (w *Wasm_Writer) Init() {
	w.buffer.Write([]byte{0x00, 0x61, 0x73, 0x6d, 0x01, 0x00, 0x00, 0x00}) //magic number and version
}
