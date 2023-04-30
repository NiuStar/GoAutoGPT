package app

import (
	"alicode.yjkj.ink/yjkj.ink/work/http"
	"encoding/base64"
	"fmt"
	"os"
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
