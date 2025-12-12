package assembler

import (
	"chasm_w/opcodes"
	"fmt"
)

func Disassemble_Instruction(b []byte) string {
	var ob string
	switch b[0] {
	case 0x41, 0x42, 0x43, 0x44:
		ob = disassemble_const(b)
	//special case - single op instructions
	case 0x1, 0x6A, 0x7C, 0x92, 0xA0:
		v, e := opcodes.Lookup_Op_Back(b[0])
		if e != nil {
			panic(fmt.Errorf("not found - opcode byte %s", b[0]))
		}
		ob = v
	}

	return ob
}
