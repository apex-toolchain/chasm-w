package assembler

import (
	"strings"
)

func Assemble_Instruction(op string) []byte {
	var ob []byte
	fields := strings.Fields(op)
	switch fields[0] {
	case "i32.const", "i64.const", "f32.const", "f64.const":
		ob = assemble_const(fields[0], fields[1:])
	}

	return ob
}
