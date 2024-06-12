package parallel

var Total_op_count_and_time map[string]map[string]int64

func Update_total_op_count_and_time(opcode string, run_time int64) {
	_, op_ok := Total_op_count_and_time[opcode]
	if !op_ok {
		Total_op_count_and_time[opcode] = map[string]int64{
			"count":      1,
			"total_time": run_time,
		}
	} else {
		Total_op_count_and_time[opcode]["count"] += 1
		Total_op_count_and_time[opcode]["total_time"] += run_time
	}
}
