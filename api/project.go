package api

import (
	"GoCodeGPT/config"
	"GoCodeGPT/project"
	"alicode.yjkj.ink/yjkj.ink/work/http"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	http2 "net/http"
	"os"
	"path"
)

type MakeModelInitRequest struct {
	UserId    string `json:"userId" binding:"required"`
	ProjectId string `json:"projectId" binding:"required"`
}
type MakeModelCodeRequest struct {
	UserId     string `json:"userId" binding:"required"`
	ProjectId  string `json:"projectId" binding:"required"`
	ModelIndex int    `json:"modelIndex" binding:"required"`
}
type MakeFunctionCodeRequest struct {
	UserId    string `json:"userId" binding:"required"`
	ProjectId string `json:"projectId" binding:"required"`
	Function1 int    `json:"function1" binding:"required"`
	Function2 int    `json:"function2" binding:"required"`
}

type CodeResponse struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
}

func createProject(c *gin.Context) {
	var req struct {
		UserId      string `json:"userId"  binding:"-"`
		ProjectName string `json:"projectName" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http2.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(config.SharePrivateConfigInstance().Uri+"project/create", req)
	c.String(http2.StatusOK, string(resp.Byte()))
}

func code(c *gin.Context) {
	fmt.Println("test")
	var req struct {
		UserId    string `json:"userId" binding:"-"`
		ProjectId string `json:"projectId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http2.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserId = config.SharePrivateConfigInstance().UserId

	resp := http.POSTJson(config.SharePrivateConfigInstance().Uri+"project/get", req)
	var proj *project.Project
	resp.Resp(&proj)

	if proj != nil {
		os.MkdirAll(path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "api"), 0777)
		os.MkdirAll(path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model"), 0777)
		for index, m := range proj.Models {
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model", m.NameEn+".go")
			codeReq := &MakeModelCodeRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = req.ProjectId
			codeReq.ModelIndex = index
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project/model/code", codeReq)
		}
		{
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model", "init.go")
			codeReq := &MakeModelInitRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = req.ProjectId
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project/model/init", codeReq)
		}
		for function1, m := range proj.FunctionsList {
			for function2, api := range m.Apis {
				fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "api", m.NameEn+"_"+api.NameEn+".go")
				codeReq := &MakeFunctionCodeRequest{}
				codeReq.UserId = config.SharePrivateConfigInstance().UserId
				codeReq.ProjectId = req.ProjectId
				codeReq.Function1 = function1
				codeReq.Function2 = function2
				genCode(fileName, config.SharePrivateConfigInstance().Uri+"project/function/code", codeReq)
			}
		}

		{
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "main.go")
			codeReq := &MakeModelInitRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = req.ProjectId
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project/code", codeReq)
		}
	}
	c.String(http2.StatusOK, string(resp.Byte()))
}
func genCode(fileName, uri string, req interface{}) error {
	if err := genCode2(fileName, uri, req); err != nil {
		return genCode(fileName, uri, req)
	}
	return nil
}
func genCode2(fileName, uri string, req interface{}) error {
	_, err := os.Stat(fileName)
	if err == nil {
		return err
	}

	respCode := http.POSTJson(uri, req)
	var codeReponse *CodeResponse
	respCode.Resp(&codeReponse)
	if codeReponse == nil {
		fmt.Println("生成代码错误：", respCode.Error(), string(respCode.Byte()), fileName)
		return fmt.Errorf("错误")
	}
	if codeReponse.Success {
		data, _ := base64.StdEncoding.DecodeString(codeReponse.Code)
		err = os.WriteFile(fileName, data, 0777)
		if err != nil {
			fmt.Println("写文件失败：", err, fileName)
			return err
		}
	} else {
		fmt.Println("生成代码失败：", string(respCode.Byte()), fileName)
		return fmt.Errorf("生成失败")
	}

	return nil
}
