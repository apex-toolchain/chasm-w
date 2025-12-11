package main

import (
	"chasm_w/assembler"
	"fmt"
	"math/rand"
	"time"
)

func benchmarkAssemblerDisassembler(n int) {
	rand.Seed(time.Now().UnixNano())

	types := []string{"i32.const", "i64.const", "f32.const", "f64.const"}
	instructions := make([]string, n)

	for i := 0; i < n; i++ {
		instType := types[rand.Intn(len(types))]

		var value string
		switch instType {
		case "i32.const":
			value = fmt.Sprintf("%d", rand.Int31())
		case "i64.const":
			value = fmt.Sprintf("%d", rand.Int63())
		case "f32.const":
			value = fmt.Sprintf("%f", rand.Float32()*1e6)
		case "f64.const":
			value = fmt.Sprintf("%f", rand.Float64()*1e6)
		}

		instructions[i] = fmt.Sprintf("%s %s", instType, value)
	}

	start := time.Now()

	for _, inst := range instructions {
		asmed := assembler.Assemble_Instruction(inst)
		_ = assembler.Disassemble_Instruction(asmed)
	}

	elapsed := time.Since(start)
	timePerInst := elapsed / time.Duration(n)
	instPerSec := float64(n) / elapsed.Seconds()
	minstPerSec := instPerSec / 1e6

	fmt.Printf("Total instructions: %d\n", n)
	fmt.Printf("Total time: %v\n", elapsed)
	fmt.Printf("Time per instruction: %v\n", timePerInst)
	fmt.Printf("Throughput: %.3f MInst/s\n", minstPerSec)
}

func main() {
	benchmarkAssemblerDisassembler(1_000_000)
}
