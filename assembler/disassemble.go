package assembler

func Disassemble_Instruction(b []byte) []string {
	var ob []string
	switch b[0] {
	case 0x41, 0x42, 0x43, 0x44:
		ob = disassemble_const(b)
	}

	return ob
}
