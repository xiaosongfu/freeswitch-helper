package decode

import "testing"

func TestX(t *testing.T) {
	tests := []struct {
		txt  string
		want string
	}{
		{
			txt: "name|instance_id|uuid|type|contact|status|state|max_no_answer|wrap_up_time|reject_delay_time|busy_delay_time|no_answer_delay_time|last_bridge_start|last_bridge_end|last_offered_call|last_status_change|no_answer_count|calls_answered|talk_time|ready_time|external_calls_count\n1018|single_box|8b844a43-eb6b-4c47-84be-e37bb788ddec|callback|user/1018|Available|In a queue call|0|0|0|0|0|1579144725|1579144510|1579144722|1579058496|0|69|9385|1579144139|0\n1015|single_box||callback|user/1015|Available|Waiting|0|0|0|0|0|1579071874|1579071878|1579144722|1579067112|0|8|178|1579144727|0\n+OK\n",
		},
	}

	for _, tt := range tests {
		got := parse(tt.txt)
		if got != tt.want {
			t.Error("error")
		}
	}
}
