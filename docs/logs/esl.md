> 错误示例：

```
version

Content-Type: command/reply
Reply-Text: -ERR command not found
```

> 正确示例：

```
api version

Content-Type: api/response
Content-Length: 113

FreeSWITCH Version 1.10.1-release+git~20190820T165404Z~f9990221e6~64bit (git f999022 2019-08-20 16:54:04Z 64bit)
```

## callcenter_config xxx list

```
api callcenter_config agent list

Content-Type: api/response
Content-Length: 582

name|instance_id|uuid|type|contact|status|state|max_no_answer|wrap_up_time|reject_delay_time|busy_delay_time|no_answer_delay_time|last_bridge_start|last_bridge_end|last_offered_call|last_status_change|no_answer_count|calls_answered|talk_time|ready_time|external_calls_count
1018|single_box|d7fae5ff-221f-48c2-b497-de844d98f9c7|callback|user/1018|Available|In a queue call|0|0|0|0|0|1579140526|1579140140|1579140523|1579058496|0|63|6949|1579079168|0
1015|single_box||callback|user/1015|Available|Waiting|0|0|0|0|0|1579071874|1579071878|1579140523|1579067112|0|8|178|1579140528|0
+OK
```

```
api callcenter_config queue list

Content-Type: api/response
Content-Length: 634

name|strategy|moh_sound|time_base_score|tier_rules_apply|tier_rule_wait_second|tier_rule_wait_multiply_level|tier_rule_no_agent_no_wait|discard_abandoned_after|abandoned_resume_allowed|max_wait_time|max_wait_time_with_no_agent|max_wait_time_with_no_agent_time_reached|record_template|calls_answered|calls_abandoned|ring_progressively_delay|skip_agents_with_external_calls|agent_no_answer_status
support@default|longest-idle-agent|local_stream://moh|system|false|300|true|false|60|false|0|0|5||0|0|0|true|On Break
xiaoi_ivr_queue|longest-idle-agent|local_stream://moh|queue|false|0|false|true|60|false|0|0|5||81|10|0|true|On Break
+OK
```

```
api callcenter_config tier list

Content-Type: api/response
Content-Length: 99

queue|agent|state|level|position
xiaoi_ivr_queue|1018|Ready|1|1
xiaoi_ivr_queue|1015|Ready|1|1
+OK
```

## callcenter_config queue list xxx xxx

```
api callcenter_config queue list agents xiaoi_ivr_queue

Content-Type: api/response
Content-Length: 538

name|instance_id|uuid|type|contact|status|state|max_no_answer|wrap_up_time|reject_delay_time|busy_delay_time|no_answer_delay_time|last_bridge_start|last_bridge_end|last_offered_call|last_status_change|no_answer_count|calls_answered|talk_time|ready_time|external_calls_count
1018|single_box||callback|user/1018|Available|Waiting|0|0|0|0|0|1579144855|1579144887|1579144853|1579058496|0|70|9455|1579144139|0
1015|single_box||callback|user/1015|Available|Waiting|0|0|0|0|0|1579071874|1579071878|1579144853|1579067112|0|8|178|1579144858|0
+OK
```

```
api callcenter_config queue list members xiaoi_ivr_queue

Content-Type: api/response
Content-Length: 4

+OK
```

```
api callcenter_config queue list tiers xiaoi_ivr_queue

Content-Type: api/response
Content-Length: 99

queue|agent|state|level|position
xiaoi_ivr_queue|1018|Ready|1|1
xiaoi_ivr_queue|1015|Ready|1|1
+OK
```
