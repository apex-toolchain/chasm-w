package opcodes

import "fmt"

// Lookup_Op returns the single-byte opcode for a given instruction name.
// Only instructions with no immediate operands are included.
func Lookup_Op(op string) (byte, error) {
	switch op {
	case "nop":
		return 0x01, nil

	// i32 comparisons
	case "i32.eqz":
		return 0x45, nil
	case "i32.eq":
		return 0x46, nil
	case "i32.ne":
		return 0x47, nil
	case "i32.lt_s":
		return 0x48, nil
	case "i32.lt_u":
		return 0x49, nil
	case "i32.gt_s":
		return 0x4A, nil
	case "i32.gt_u":
		return 0x4B, nil
	case "i32.le_s":
		return 0x4C, nil
	case "i32.le_u":
		return 0x4D, nil
	case "i32.ge_s":
		return 0x4E, nil
	case "i32.ge_u":
		return 0x4F, nil

	// i64 comparisons
	case "i64.eqz":
		return 0x50, nil
	case "i64.eq":
		return 0x51, nil
	case "i64.ne":
		return 0x52, nil
	case "i64.lt_s":
		return 0x53, nil
	case "i64.lt_u":
		return 0x54, nil
	case "i64.gt_s":
		return 0x55, nil
	case "i64.gt_u":
		return 0x56, nil
	case "i64.le_s":
		return 0x57, nil
	case "i64.le_u":
		return 0x58, nil
	case "i64.ge_s":
		return 0x59, nil
	case "i64.ge_u":
		return 0x5A, nil

	// i32 math
	case "i32.clz":
		return 0x67, nil
	case "i32.ctz":
		return 0x68, nil
	case "i32.popcnt":
		return 0x69, nil
	case "i32.add":
		return 0x6A, nil
	case "i32.sub":
		return 0x6B, nil
	case "i32.mul":
		return 0x6C, nil
	case "i32.div_s":
		return 0x6D, nil
	case "i32.div_u":
		return 0x6E, nil
	case "i32.rem_s":
		return 0x6F, nil
	case "i32.rem_u":
		return 0x70, nil
	case "i32.and":
		return 0x71, nil
	case "i32.or":
		return 0x72, nil
	case "i32.xor":
		return 0x73, nil
	case "i32.shl":
		return 0x74, nil
	case "i32.shr_s":
		return 0x75, nil
	case "i32.shr_u":
		return 0x76, nil
	case "i32.rotl":
		return 0x77, nil
	case "i32.rotr":
		return 0x78, nil

	// i64 math
	case "i64.clz":
		return 0x79, nil
	case "i64.ctz":
		return 0x7A, nil
	case "i64.popcnt":
		return 0x7B, nil
	case "i64.add":
		return 0x7C, nil
	case "i64.sub":
		return 0x7D, nil
	case "i64.mul":
		return 0x7E, nil
	case "i64.div_s":
		return 0x7F, nil
	case "i64.div_u":
		return 0x80, nil
	case "i64.rem_s":
		return 0x81, nil
	case "i64.rem_u":
		return 0x82, nil
	case "i64.and":
		return 0x83, nil
	case "i64.or":
		return 0x84, nil
	case "i64.xor":
		return 0x85, nil
	case "i64.shl":
		return 0x86, nil
	case "i64.shr_s":
		return 0x87, nil
	case "i64.shr_u":
		return 0x88, nil
	case "i64.rotl":
		return 0x89, nil
	case "i64.rotr":
		return 0x8A, nil

	// f32 math
	case "f32.abs":
		return 0x8B, nil
	case "f32.neg":
		return 0x8C, nil
	case "f32.ceil":
		return 0x8D, nil
	case "f32.floor":
		return 0x8E, nil
	case "f32.trunc":
		return 0x8F, nil
	case "f32.nearest":
		return 0x90, nil
	case "f32.sqrt":
		return 0x91, nil
	case "f32.add":
		return 0x92, nil
	case "f32.sub":
		return 0x93, nil
	case "f32.mul":
		return 0x94, nil
	case "f32.div":
		return 0x95, nil
	case "f32.min":
		return 0x96, nil
	case "f32.max":
		return 0x97, nil
	case "f32.copysign":
		return 0x98, nil

	// f64 math
	case "f64.abs":
		return 0x99, nil
	case "f64.neg":
		return 0x9A, nil
	case "f64.ceil":
		return 0x9B, nil
	case "f64.floor":
		return 0x9C, nil
	case "f64.trunc":
		return 0x9D, nil
	case "f64.nearest":
		return 0x9E, nil
	case "f64.sqrt":
		return 0x9F, nil
	case "f64.add":
		return 0xA0, nil
	case "f64.sub":
		return 0xA1, nil
	case "f64.mul":
		return 0xA2, nil
	case "f64.div":
		return 0xA3, nil
	case "f64.min":
		return 0xA4, nil
	case "f64.max":
		return 0xA5, nil
	case "f64.copysign":
		return 0xA6, nil
	}

	return 0, fmt.Errorf("unknown operation %s", op)
}

// Lookup_Op_Back returns the instruction name for a single-byte opcode.
func Lookup_Op_Back(op byte) (string, error) {
	switch op {
	case 0x01:
		return "nop", nil

	// i32 comparisons
	case 0x45:
		return "i32.eqz", nil
	case 0x46:
		return "i32.eq", nil
	case 0x47:
		return "i32.ne", nil
	case 0x48:
		return "i32.lt_s", nil
	case 0x49:
		return "i32.lt_u", nil
	case 0x4A:
		return "i32.gt_s", nil
	case 0x4B:
		return "i32.gt_u", nil
	case 0x4C:
		return "i32.le_s", nil
	case 0x4D:
		return "i32.le_u", nil
	case 0x4E:
		return "i32.ge_s", nil
	case 0x4F:
		return "i32.ge_u", nil

	// i64 comparisons
	case 0x50:
		return "i64.eqz", nil
	case 0x51:
		return "i64.eq", nil
	case 0x52:
		return "i64.ne", nil
	case 0x53:
		return "i64.lt_s", nil
	case 0x54:
		return "i64.lt_u", nil
	case 0x55:
		return "i64.gt_s", nil
	case 0x56:
		return "i64.gt_u", nil
	case 0x57:
		return "i64.le_s", nil
	case 0x58:
		return "i64.le_u", nil
	case 0x59:
		return "i64.ge_s", nil
	case 0x5A:
		return "i64.ge_u", nil

	// i32 math
	case 0x67:
		return "i32.clz", nil
	case 0x68:
		return "i32.ctz", nil
	case 0x69:
		return "i32.popcnt", nil
	case 0x6A:
		return "i32.add", nil
	case 0x6B:
		return "i32.sub", nil
	case 0x6C:
		return "i32.mul", nil
	case 0x6D:
		return "i32.div_s", nil
	case 0x6E:
		return "i32.div_u", nil
	case 0x6F:
		return "i32.rem_s", nil
	case 0x70:
		return "i32.rem_u", nil
	case 0x71:
		return "i32.and", nil
	case 0x72:
		return "i32.or", nil
	case 0x73:
		return "i32.xor", nil
	case 0x74:
		return "i32.shl", nil
	case 0x75:
		return "i32.shr_s", nil
	case 0x76:
		return "i32.shr_u", nil
	case 0x77:
		return "i32.rotl", nil
	case 0x78:
		return "i32.rotr", nil

	// i64 math
	case 0x79:
		return "i64.clz", nil
	case 0x7A:
		return "i64.ctz", nil
	case 0x7B:
		return "i64.popcnt", nil
	case 0x7C:
		return "i64.add", nil
	case 0x7D:
		return "i64.sub", nil
	case 0x7E:
		return "i64.mul", nil
	case 0x7F:
		return "i64.div_s", nil
	case 0x80:
		return "i64.div_u", nil
	case 0x81:
		return "i64.rem_s", nil
	case 0x82:
		return "i64.rem_u", nil
	case 0x83:
		return "i64.and", nil
	case 0x84:
		return "i64.or", nil
	case 0x85:
		return "i64.xor", nil
	case 0x86:
		return "i64.shl", nil
	case 0x87:
		return "i64.shr_s", nil
	case 0x88:
		return "i64.shr_u", nil
	case 0x89:
		return "i64.rotl", nil
	case 0x8A:
		return "i64.rotr", nil

	// f32 math
	case 0x8B:
		return "f32.abs", nil
	case 0x8C:
		return "f32.neg", nil
	case 0x8D:
		return "f32.ceil", nil
	case 0x8E:
		return "f32.floor", nil
	case 0x8F:
		return "f32.trunc", nil
	case 0x90:
		return "f32.nearest", nil
	case 0x91:
		return "f32.sqrt", nil
	case 0x92:
		return "f32.add", nil
	case 0x93:
		return "f32.sub", nil
	case 0x94:
		return "f32.mul", nil
	case 0x95:
		return "f32.div", nil
	case 0x96:
		return "f32.min", nil
	case 0x97:
		return "f32.max", nil
	case 0x98:
		return "f32.copysign", nil

	// f64 math
	case 0x99:
		return "f64.abs", nil
	case 0x9A:
		return "f64.neg", nil
	case 0x9B:
		return "f64.ceil", nil
	case 0x9C:
		return "f64.floor", nil
	case 0x9D:
		return "f64.trunc", nil
	case 0x9E:
		return "f64.nearest", nil
	case 0x9F:
		return "f64.sqrt", nil
	case 0xA0:
		return "f64.add", nil
	case 0xA1:
		return "f64.sub", nil
	case 0xA2:
		return "f64.mul", nil
	case 0xA3:
		return "f64.div", nil
	case 0xA4:
		return "f64.min", nil
	case 0xA5:
		return "f64.max", nil
	case 0xA6:
		return "f64.copysign", nil
	}

	return "", fmt.Errorf("unknown opcode 0x%X", op)
}
