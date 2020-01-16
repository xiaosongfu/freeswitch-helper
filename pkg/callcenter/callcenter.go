package callcenter

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaosongfu/eslgo"
)

type CallCenter struct {
	Esl *eslgo.Connection
}

func NewCallCenter(esl *eslgo.Connection) *CallCenter {
	return &CallCenter{Esl: esl}
}

func (cc *CallCenter) RegisterRoutes(engine *gin.Engine) {
	// 坐席相关的接口
	agent := engine.Group("/agent")
	{
		agent.GET("/list", cc.listAgent)
		agent.GET("/add/:name", cc.addAgent)
		agent.GET("/del/:name", cc.delAgent)
		agent.POST("/updateStatus/:name", cc.updateAgentStatus)
		agent.POST("/updateContact/:name", cc.updateAgentContact)
	}

	// 队列相关的接口
	queue := engine.Group("/queue")
	{
		queue.GET("/list", cc.listQueue)
		queue.GET("/listQueueAgents/:queue", cc.listQueueAgents)
		queue.GET("/listQueueMembers/:queue", cc.listQueueMembers)
	}

	// 梯队相关的接口
	tier := engine.Group("/tier")
	{
		tier.GET("/list", cc.listTier)
		tier.GET("/add/:queue/:agent", cc.addTier)
		tier.GET("/del/:queue/:agent", cc.delTier)
	}
}
