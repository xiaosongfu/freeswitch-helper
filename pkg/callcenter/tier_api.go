package callcenter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaosongfu/freeswitch-helper/pkg/callcenter/freeswitch"
)

// listTier list all tiers
//
// @Summary list all tiers
// @Tags Tier
// @Router /tier/list [get]
//
func (cc *CallCenter) listTier(ctx *gin.Context) {
	evt, err := cc.esl.ApiCommand(freeswitch.TierList)
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

// addTier add a tier
//
// @Summary add a tier
// @Tags Tier
// @Param queue path string true "the queue of which you want to add to"
// @Param agent path string true "the agent of which you want to add"
// @Router /tier/add/{queue}/{agent} [get]
//
func (cc *CallCenter) addTier(ctx *gin.Context) {
	queue := ctx.Param("queue")
	agent := ctx.Param("agent")

	evt, err := cc.esl.ApiCommand(fmt.Sprintf(freeswitch.TierAdd, queue, agent))
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

// delTier delete a tier
//
// @Summary delete a tier
// @Tags Tier
// @Param queue path string true "the queue of which you want to add to"
// @Param agent path string true "the agent of which you want to add"
// @Router /tier/del/{queue}/{agent} [get]
//
func (cc *CallCenter) delTier(ctx *gin.Context) {
	queue := ctx.Param("queue")
	agent := ctx.Param("agent")

	evt, err := cc.esl.ApiCommand(fmt.Sprintf(freeswitch.TierDel, queue, agent))
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
