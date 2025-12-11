package w_util

import "fmt"

func Encode_UL128(n uint64) []byte {
	var result []byte
	for n != 0 {
		o_byte := n & 0x7f
		n >>= 7
		if n != 0 {
			o_byte |= 0x80
		}
		result = append(result, byte(o_byte))
	}
	return result
}

func Decode_UL128(data []byte) (uint64, int, error) {
	var result uint64
	var shift uint

	for i, b := range data {
		result |= uint64(b&0x7F) << shift

		// last byte (continuation bit cleared)
		if (b & 0x80) == 0 {
			return result, i + 1, nil
		}

		shift += 7
	}

	// No terminating byte found -> malformed input
	return 0, 0, fmt.Errorf("unterminated UL128 value")
}

func Encode_SL128(n int64, bitsAmount uint8) []byte {
	negative := n < 0
	size := bitsAmount

	result := make([]byte, 0, 10)

	for {
		oByte := byte(n & 0x7F)

		n >>= 7

		if negative {
			n |= int64(^0) << (size - 7)
		}

		signBit := oByte & 0x40
		last :=
			(!negative && n == 0 && signBit == 0) ||
				(negative && n == -1 && signBit != 0)

		if !last {
			oByte |= 0x80
		}

		result = append(result, oByte)

		if last {
			break
		}
	}

	return result
}

func Decode_SL128(data []byte) (int64, int, error) {
	var result int64
	var shift uint
	var b byte

	for i := 0; i < len(data); i++ {
		b = data[i]

		result |= int64(b&0x7F) << shift
		shift += 7

		if (b & 0x80) == 0 { // last byte
			if (shift < 64) && (b&0x40 != 0) {
				result |= ^0 << shift
			}
			return result, i + 1, nil
		}
	}

	return 0, 0, fmt.Errorf("unterminated SL128")
}
