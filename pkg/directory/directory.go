package directory

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Directory struct {
	db *sql.DB
}

func NewDirectory(db *sql.DB) *Directory {
	return &Directory{
		db: db,
	}
}

func (d *Directory) RegisterRoutes(engine *gin.Engine) {
	dr := engine.Group("/directory")
	{
		dr.POST("/export", d.exportDirectory)

		dr.POST("/add", d.addDirectory)
	}
}
