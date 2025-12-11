package w_util

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func EncodeFloat(value float64, bitsize int) ([]byte, error) {
	buf := new(bytes.Buffer)

	switch bitsize {
	case 32:
		// Convert to float32 first
		f32 := float32(value)
		if err := binary.Write(buf, binary.LittleEndian, f32); err != nil {
			return nil, err
		}
	case 64:
		if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported bit size, must be 32 or 64")
	}

	return buf.Bytes(), nil
}

func DecodeFloat(data []byte, bitsize int) (float64, error) {
	buf := bytes.NewReader(data)

	switch bitsize {
	case 32:
		var f32 float32
		if err := binary.Read(buf, binary.LittleEndian, &f32); err != nil {
			return 0, err
		}
		return float64(f32), nil
	case 64:
		var f64 float64
		if err := binary.Read(buf, binary.LittleEndian, &f64); err != nil {
			return 0, err
		}
		return f64, nil
	default:
		return 0, errors.New("unsupported bit size, must be 32 or 64")
	}
}
