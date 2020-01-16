package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaosongfu/eslgo"
	"github.com/xiaosongfu/freeswitch-helper/pkg/callcenter"
)

// @title FreeSWITCH 助手项目接口文档
// @version v0.1.0
// @description FreeSWITCH 助手，包括呼叫中心、CDR、用户配置、拨号计划配置等
// @termsOfService http://swagger.io/terms/
//
// @contact.name xiaosongfu
// @contact.url https://xiaosongfu.com
// @contact.email dev@xiaosongfu.com
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:12052
// @BasePath /
//
func main() {
	r := gin.New()

	// add middleware
	r.Use(gin.Recovery())

	// connect esl
	esl, err := eslgo.Dial("192.168.160.10:8021", "ClueCon")
	if err != nil {
		panic(err)
	}

	// register call center
	callCenter := callcenter.NewCallCenter(esl)
	callCenter.RegisterRoutes(r)

	_ = r.Run(":12052")
}
