package parallel

var total_op_count_and_time map[string]map[string]int64 = make(map[string]map[string]int64)

func update_total_op_count_and_time(opcode string, run_time int64) {
	_, op_ok := total_op_count_and_time[opcode]
	if !op_ok {
		total_op_count_and_time[opcode] = map[string]int64{
			"count":      1,
			"total_time": run_time,
		}
	} else {
		total_op_count_and_time[opcode]["count"] += 1
		total_op_count_and_time[opcode]["total_time"] += run_time
	}
}
