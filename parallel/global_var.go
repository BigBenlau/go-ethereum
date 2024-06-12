package parallel

import "fmt"

var Total_op_count_map = make(map[string]int64)

var Total_op_time_map = make(map[string]int64)

var op_channel = make(chan map[string]int64, 100000)

func Update_total_op_count_and_time(opcode string, run_time int64) {
	map_value := map[string]int64{
		opcode: run_time,
	}
	op_channel <- map_value
}

func Start_channel() {
	go collect_data_from_channel()
}

func collect_data_from_channel() {
	for {
		ch_data := <-op_channel
		for opcode, run_time := range ch_data {
			_, op_ok := Total_op_count_map[opcode]
			if !op_ok {
				Total_op_count_map[opcode] = 0
				Total_op_time_map[opcode] = 0
			}
			Total_op_count_map[opcode] += 1
			Total_op_time_map[opcode] += run_time
		}
	}
}

func Print_total_op_count_and_time() {
	for op_code, op_count := range Total_op_count_map {
		op_run_time := Total_op_time_map[op_code]
		fmt.Println("Opcode name is: ", op_code, ". Total Run time as nanos: ", op_run_time, ". Total Count is: ", op_count)
	}
}
