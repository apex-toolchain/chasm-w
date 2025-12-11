package opcodes

import "fmt"

func Lookup_Op(op string) (byte, error) {
	switch op {
	case "i32.const":
		return 0x41, nil
	case "i64.const":
		return 0x42, nil
	case "f32.const":
		return 0x43, nil
	case "f64.const":
		return 0x44, nil
	case "i32.add":
		return 0x6A, nil
	case "i64.add":
		return 0x7C, nil
	case "f32.add":
		return 0x92, nil
	case "f64.add":
		return 0xA0, nil
	}
	return 0, fmt.Errorf("unknown operation %s", op)
}

func Lookup_Op_Back(op byte) (string, error) {
	switch op {
	case 0x41:
		return "i32.const", nil
	case 0x42:
		return "i64.const", nil
	case 0x43:
		return "f32.const", nil
	case 0x44:
		return "f64.const", nil
	case 0x6A:
		return "i32.add", nil
	case 0x7C:
		return "i64.add", nil
	case 0x92:
		return "f32.add", nil
	case 0xA0:
		return "f64.add", nil
	}
	return "", fmt.Errorf("unknown opcode 0x%X", op)
}
