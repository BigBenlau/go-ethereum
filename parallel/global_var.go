package parallel

import "fmt"

var Total_op_count_map = make(map[string]int64)

var Total_op_time_map = make(map[string]int64)

func Update_total_op_count_and_time(opcode string, run_time int64) {
	_, op_ok := Total_op_count_map[opcode]
	fmt.Println("opcode: ", opcode, "op_ok: ", op_ok)
	if !op_ok {
		Total_op_count_map[opcode] = 0
		Total_op_time_map[opcode] = 0
	}
	Total_op_count_map[opcode] += 1
	Total_op_time_map[opcode] += run_time
}

func Print_total_op_count_and_time() {
	for op_code, _ := range Total_op_count_map {
		op_run_time := Total_op_time_map[op_code]
		fmt.Println("Opcode name is", op_code, "Run time as nanos: ", op_run_time)
	}
}
