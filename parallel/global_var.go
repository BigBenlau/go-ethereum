package parallel

import "fmt"

var Total_op_count_and_time map[string]map[string]int64

func Update_total_op_count_and_time(opcode string, run_time int64) {
	_, op_ok := Total_op_count_and_time[opcode]
	fmt.Println("opcode: ", opcode, "op_ok: ", op_ok)
	if !op_ok {
		Total_op_count_and_time[opcode]["count"] = 1
		Total_op_count_and_time[opcode]["total_time"] = run_time
	} else {
		Total_op_count_and_time[opcode]["count"] += 1
		Total_op_count_and_time[opcode]["total_time"] += run_time
	}
}

func Print_total_op_count_and_time() {
	for op_code, time_value_list := range Total_op_count_and_time {
		// op_count := time_value_list["count"]
		op_run_time := time_value_list["total_time"]
		fmt.Println("Opcode name is", op_code, "Run time as nanos: ", op_run_time)
	}
}
