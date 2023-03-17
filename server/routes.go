package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schlucht/fhxreader/models"
)

var u = models.UnitList{}

func ReadUnitName(c *gin.Context) {
	u := models.UnitList{}
	u.NewUnitList()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"header":    "Startseite",
		"unitNames": u.UnitNames,
	})
}

func ReadUnitDetail(c *gin.Context) {
	name := c.Query("unit")

	if name != "" {
		li := u.UPNames(name)
		c.HTML(http.StatusOK, "up.html", gin.H{
			"header":  "Unit Procedures",
			"upnames": li,
		})
	}
}
