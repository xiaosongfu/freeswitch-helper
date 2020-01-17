package callcenter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaosongfu/freeswitch-helper/pkg/callcenter/freeswitch"
)

// listQueue list all queues
//
// @Summary list all queues
// @Tags Queue
// @Router /queue/list [get]
//
func (cc *CallCenter) listQueue(ctx *gin.Context) {
	evt, err := cc.esl.ApiCommand(freeswitch.QueueList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": evt.Body,
	})
}

// listQueueAgents list all agents of a queue
//
// @Summary list all agents of a queue
// @Tags Queue
// @Param queue path string true "the queue of which you want to query"
// @Router /queue/listQueueAgents/{queue} [get]
//
func (cc *CallCenter) listQueueAgents(ctx *gin.Context) {
	queue := ctx.Param("queue")

	evt, err := cc.esl.ApiCommand(fmt.Sprintf(freeswitch.QueueAgentsList, queue))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": evt.Body,
	})
}

// listQueueMembers list all members of a queue
//
// @Summary list all members of a queue
// @Tags Queue
// @Param queue path string true "the queue of which you want to query"
// @Router /queue/listQueueMembers/{queue} [get]
//
func (cc *CallCenter) listQueueMembers(ctx *gin.Context) {
	queue := ctx.Param("queue")

	evt, err := cc.esl.ApiCommand(fmt.Sprintf(freeswitch.QueueMembersList, queue))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": evt.Body,
	})
}

// listQueueTiers list all tiers of a queue
//
// @Summary list all tiers of a queue
// @Tags Queue
// @Param queue path string true "the queue of which you want to query"
// @Router /queue/listQueueTiers/{queue} [get]
//
func (cc *CallCenter) listQueueTiers(ctx *gin.Context) {
	queue := ctx.Param("queue")

	evt, err := cc.esl.ApiCommand(fmt.Sprintf(freeswitch.QueueTiersList, queue))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": evt.Body,
	})
}
