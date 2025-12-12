package assembler

import (
	"chasm_w/opcodes"
	"strings"
)

func Assemble_Instruction(op string) []byte {
	var ob []byte
	fields := strings.Fields(op)
	switch fields[0] {
	case "i32.const", "i64.const", "f32.const", "f64.const":
		ob = assemble_const(fields[0], fields[1:])
	case "i32.add", "i64.add", "f32.add", "f64.add", "nop":
		b, e := opcodes.Lookup_Op(op)
		if e != nil {
			panic(e)
		}
		ob = []byte{b}
	}

	return ob
}
