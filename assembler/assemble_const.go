package assembler

import (
	"chasm_w/opcodes"
	"chasm_w/w_util"
	"strconv"
	"strings"
)

func assemble_const(opcode string, rf []string) []byte {
	opcode_b, err := opcodes.Lookup_Op(opcode)
	if err != nil {
		panic("error looking up opcode, " + opcode)
	}

	var result []byte
	switch opcode_b {
	case 0x41, 0x42: // i32/i64
		str := rf[0]
		n, e := strconv.ParseInt(str, 10, 64)
		if e != nil {
			panic("bru, please make ur number numberable - " + str)
		}
		encoded := w_util.Encode_SL128(n, uint8(opcode_b-0x41)*32)
		result = make([]byte, 1+len(encoded))
		result[0] = opcode_b
		copy(result[1:], encoded)

	case 0x43, 0x44: // f32/f64
		str := rf[0]
		n, e := strconv.ParseFloat(str, 64)
		if e != nil {
			panic("bru, please make ur number numberable - " + str)
		}
		b2, e2 := w_util.EncodeFloat(n, int(opcode_b-0x42)*32)
		if e2 != nil {
			panic(e2)
		}
		result = make([]byte, 1+len(b2))
		result[0] = opcode_b
		copy(result[1:], b2)
	}

	return result
}

var lkp_b = []int{32, 64}

func disassemble_const(b []byte) string {
	if len(b) == 0 {
		panic("empty instruction")
	}

	op := b[0]
	outs := make([]string, 0, 2) // pre-allocate for speed
	o, e := opcodes.Lookup_Op_Back(op)
	if e != nil {
		panic(e)
	}
	outs = append(outs, o)

	rest := b[1:]
	if len(rest) == 0 {
		panic("no bytes left to decode")
	}

	if op >= 0x43 { // f32/f64
		decoded, e := w_util.DecodeFloat(rest, lkp_b[op-0x43])
		if e != nil {
			panic(e)
		}
		// use strconv for faster float to string
		outs = append(outs, strconv.FormatFloat(decoded, 'f', -1, 64))
	} else { // i32/i64
		decoded, _, e := w_util.Decode_UL128(rest)
		if e != nil {
			panic(e)
		}
		// use strconv for faster int to string
		outs = append(outs, strconv.FormatInt(int64(decoded), 10))
	}

	return strings.Join(outs, " ")
}
