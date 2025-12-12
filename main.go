package main

import (
	"chasm_w/assembler"
	"fmt"
	"time"
)

func main() {
	// Array of 10 instructions to test
	instructions := []string{
		"f32.const 2.2",
		"i32.const 10",
		"i32.add",
		"i32.add",
		"f32.const 3.14",
		"i32.add",
		"i32.const 7",
		"f32.const 1.1",
		"i32.add",
		"f32.const 0.5",
	}

	repeats := 2_000
	totalInst := len(instructions) * repeats

	// Precompute assembled instructions to avoid repeated allocations
	assembled := make([][]byte, len(instructions))
	for i, instr := range instructions {
		assembled[i] = assembler.Assemble_Instruction(instr)
	}

	start := time.Now()

	for r := 0; r < repeats; r++ {
		for _, asmed := range assembled {
			_ = assembler.Disassemble_Instruction(asmed)
		}
	}

	duration := time.Since(start)
	totalTimeSec := duration.Seconds()
	minstPerSec := float64(totalInst) / 1e6 / totalTimeSec
	timePerInstNs := duration.Nanoseconds() / int64(totalInst)

	fmt.Println("Total instructions:", totalInst)
	fmt.Printf("Total time: %.6f sec\n", totalTimeSec)
	fmt.Printf("Throughput: %.3f MInst/s\n", minstPerSec)
	fmt.Printf("Time per instruction: %d ns\n", timePerInstNs)
}
