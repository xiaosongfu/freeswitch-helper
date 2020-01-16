package freeswitch

const (
	AgentList       = "callcenter_config agent list"               // 列出所有坐席
	AgentAdd        = "callcenter_config agent add %s callback"    // 添加坐席
	AgentDel        = "callcenter_config agent del %s"             // 删除坐席
	AgentSetStatus  = "callcenter_config agent set status %s '%s'" // 修改坐席状态
	AgentSetContact = "callcenter_config agent set contact %s %s"  // 修改坐席呼入字符串
)

const (
	QueueList        = "callcenter_config queue list"            // 列出所有队列
	QueueAgentsList  = "callcenter_config queue list agents %s"  // 列出指定队列的所有坐席
	QueueMembersList = "callcenter_config queue list members %s" // 列出指定队列的所有外线成员
	QueueTiersList   = "callcenter_config queue list tiers %s"   // 列出指定队列的所有梯队
)

const (
	TierList = "callcenter_config tier list"      // 列出所有梯队
	TierAdd  = "callcenter_config tier add %s %s" // 添加梯队
	TierDel  = "callcenter_config tier del %s %s" // 删除梯队
)
