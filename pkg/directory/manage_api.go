package directory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// exportDirectory provide directory xml config content for FreeSWTICH
//
// @Summary provide directory xml config content for FreeSWTICH
// @Tag Directory
// @Router /directory/add [post]
//
func (d *Directory) addDirectory(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"abc": "abc",
	})
}
