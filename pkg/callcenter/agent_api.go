package callcenter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaosongfu/freeswitch-helper/pkg/callcenter/freeswitch"
)

// listAgent list all agents
//
// @Summary list all agents
// @Tags Agent
// @Router /agent/list [get]
//
func (cc *CallCenter) listAgent(ctx *gin.Context) {
	evt, err := cc.Esl.ApiCommand(freeswitch.AgentList)
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

// addAgent add a agent
//
// @Summary add a agent
// @Tags Agent
// @Param name path string true "the agent name,always is a number,such as 1010"
// @Router /agent/add/{name} [get]
//
func (cc *CallCenter) addAgent(ctx *gin.Context) {
	name := ctx.Param("name")

	evt, err := cc.Esl.ApiCommand(fmt.Sprintf(freeswitch.AgentAdd, name))
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

// delAgent delete a agent
//
// @Summary delete a agent
// @Tags Agent
// @Param name path string true "the agent name,always is a number,such as 1010"
// @Router /agent/del/{name} [get]
//
func (cc *CallCenter) delAgent(ctx *gin.Context) {
	name := ctx.Param("name")

	evt, err := cc.Esl.ApiCommand(fmt.Sprintf(freeswitch.AgentDel, name))
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

// updateAgentStatus update agent's status
//
// @Summary update agent's status
// @Tags Agent
// @Param name path string true "the agent name,always is a number,such as 1010"
// @Param status body string true "the new status of the agent,available value is 'Available' and 'Logged Out'"
// @Router /agent/updateAgentStatus/{name} [post]
//
func (cc *CallCenter) updateAgentStatus(ctx *gin.Context) {
	name := ctx.Param("name")
	status := ctx.PostForm("status")

	evt, err := cc.Esl.ApiCommand(fmt.Sprintf(freeswitch.AgentSetStatus, name, status))
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

// updateAgentContact update agent's contact
//
// @Summary update agent's contact
// @Tags Agent
// @Param name path string true "the agent name,always is a number,such as 1010"
// @Param contact body string true "the contact of the agent,such as 'user/1010'"
// @Router /agent/updateAgentContact/{name} [post]
//
func (cc *CallCenter) updateAgentContact(ctx *gin.Context) {
	name := ctx.Param("name")
	contact := ctx.PostForm("contact")

	evt, err := cc.Esl.ApiCommand(fmt.Sprintf(freeswitch.AgentSetContact, name, contact))
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
