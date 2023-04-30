package app

import (
	"GoCodeGPT/config"
	"GoCodeGPT/project"
	"alicode.yjkj.ink/yjkj.ink/work/http"
	"fmt"
	http2 "net/http"
	"os"
	"path"
)

type Project struct {
	Uuid           string
	Name           string
	NameEn         string
	PMDescription  string
	ArcDescription string
	DBADescription string
	Functions      string
	Models         string
	Error          *string `json:"error"`
}

type Response struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	NameEn    string `json:"name_en"`
	Models    []*project.Model
	Functions []*project.Functions
}

func Genarate(projectId string) error {
	fmt.Println("第二步：初始化项目")

	proj, err := projectReq(projectId, "init")
	if err != nil {
		return err
	}
	fmt.Println("第三步：功能分解")

	var functions []*project.Functions
	resp, err := projectMake(projectId, "function/create")
	if err != nil {
		return err
	}
	var nameEn = resp.NameEn
	functions = append(functions, resp.Functions...)
	var models []*project.Model
	fmt.Println("第四步：数据库设计")

	resp, err = projectMake(projectId, "model/create")
	if err != nil {
		return err
	}
	models = append(models, resp.Models...)
	proj.NameEn = nameEn
	fmt.Println("第五步：生成代码")

	if proj != nil {
		err = os.MkdirAll(path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "api"), 0777)
		if err != nil {
			return err
		}
		err = os.MkdirAll(path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model"), 0777)
		if err != nil {
			return err
		}
		total := len(models) + 2
		for _, m := range functions {
			total += len(m.Apis)
		}
		pos := 0
		for index, m := range models {
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model", m.NameEn+".go")
			codeReq := &MakeModelCodeRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = projectId
			codeReq.ModelIndex = index
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project2/model/code", codeReq)
			pos++
			fmt.Println("生成进度如下：", pos, "/", total)
		}
		{
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "model", "init.go")
			codeReq := &MakeModelInitRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = projectId
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project2/model/init", codeReq)
			pos++
			fmt.Println("生成进度如下：", pos, "/", total)
		}
		for function1, m := range functions {
			for function2, api := range m.Apis {
				fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "api", m.NameEn+"_"+api.NameEn+".go")
				codeReq := &MakeFunctionCodeRequest{}
				codeReq.UserId = config.SharePrivateConfigInstance().UserId
				codeReq.ProjectId = projectId
				codeReq.Function1 = function1
				codeReq.Function2 = function2
				genCode(fileName, config.SharePrivateConfigInstance().Uri+"project2/function/code", codeReq)
				pos++
				fmt.Println("生成进度如下：", pos, "/", total)
			}
		}

		{
			fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "main.go")
			codeReq := &MakeModelInitRequest{}
			codeReq.UserId = config.SharePrivateConfigInstance().UserId
			codeReq.ProjectId = projectId
			genCode(fileName, config.SharePrivateConfigInstance().Uri+"project2/code", codeReq)
			pos++
			fmt.Println("生成进度如下：", pos, "/", total)
		}
	} else {
		return fmt.Errorf("项目id为kong")
	}
	{
		fileName := path.Join(config.SharePrivateConfigInstance().Src, proj.NameEn, "go.mod")
		err = os.WriteFile(fileName, createGoMod(proj.NameEn), 0777)
		if err != nil {
			fmt.Println("写文件失败：", err, fileName)
			return err
		}
	}
	fmt.Println("第六步：生成成功")

	return nil
}
func Create(appName, description string) error {
	fmt.Println("第一步：创建项目")
	projectId, err := createProject(appName, description)
	if err != nil {
		return err
	}
	return Genarate(projectId)
}

func createProject(projectName, description string) (string, error) {
	var req struct {
		UserId      string `json:"userId"  binding:"-"`
		ProjectName string `json:"projectName" binding:"required"`
		Description string `json:"description" binding:"-"`
	}
	req.ProjectName = projectName
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(config.SharePrivateConfigInstance().Uri+"project2/create", req)
	if resp.StatusCode != http2.StatusOK {
		return "", fmt.Errorf("create project error:%s", resp.Error())
	}

	var proj *Project
	if err := resp.Resp(&proj); err != nil {
		return "", err
	}

	return proj.Uuid, nil
}

func projectReq(projectId, method string) (*Project, error) {
	var req struct {
		UserId    string `json:"userId"  binding:"-"`
		ProjectId string `json:"projectId" binding:"required"`
	}
	req.ProjectId = projectId
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(fmt.Sprintf(config.SharePrivateConfigInstance().Uri+"project2/%s", method), req)
	if resp.StatusCode != http2.StatusOK {
		return nil, fmt.Errorf("%s project error:%s", method, resp.Error())
	}

	var proj *Project
	if err := resp.Resp(&proj); err != nil {
		return nil, fmt.Errorf("%s project error:%s %s", method, err.Error(), string(resp.Byte()))
	}
	if proj.Error != nil {
		return nil, fmt.Errorf(*proj.Error)
	}
	return proj, nil
}

func projectMake(projectId, method string) (*Response, error) {
	var req struct {
		UserId    string `json:"userId"  binding:"-"`
		ProjectId string `json:"projectId" binding:"required"`
	}
	req.ProjectId = projectId
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(fmt.Sprintf(config.SharePrivateConfigInstance().Uri+"project2/%s", method), req)
	if resp.StatusCode != http2.StatusOK {
		return nil, fmt.Errorf("%s project make error", method, resp.Error())
	}

	var respSuccess *Response
	if err := resp.Resp(&respSuccess); err != nil {
		return nil, fmt.Errorf("%s project make error:%s %s", method, err.Error(), string(resp.Byte()))
	}
	if respSuccess.Success {
		return respSuccess, nil
	}
	return nil, fmt.Errorf(respSuccess.Error)
}
