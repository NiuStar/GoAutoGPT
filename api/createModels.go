package api

import (
	"GoCodeGPT/config"
	"GoCodeGPT/project"
	"alicode.yjkj.ink/yjkj.ink/work/http"
	"github.com/gin-gonic/gin"
	http2 "net/http"
)

func createProjectModels(c *gin.Context) {
	var req struct {
		UserId      string `json:"userId"  binding:"-"`
		ProjectId   string `json:"projectId" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http2.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(config.SharePrivateConfigInstance().Uri+"project/model/create", req)
	c.String(http2.StatusOK, string(resp.Byte()))
}

func makeSureProjectModels(c *gin.Context) {
	var req struct {
		UserId      string           `json:"userId"  binding:"-"`
		ProjectId   string           `json:"projectId" binding:"required"`
		Description string           `json:"description"`
		Models      []*project.Model `json:"models" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http2.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(config.SharePrivateConfigInstance().Uri+"project/model/create", req)
	c.String(http2.StatusOK, string(resp.Byte()))
}
