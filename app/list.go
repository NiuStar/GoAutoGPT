package app

import (
	"GoCodeGPT/config"
	"alicode.yjkj.ink/yjkj.ink/work/http"
	"fmt"
	http2 "net/http"
)

type Project2 struct {
	Uuid           string
	Name           string
	NameEn         string
	PMDescription  string
	ArcDescription string
	DBADescription string
}

func List() error {

	projs, err := projectList(config.SharePrivateConfigInstance().UserId)
	for _, p := range projs {
		fmt.Println("name:", p.Name, "id", p.Uuid, "nameEn", p.NameEn)
	}
	return err
}

func projectList(userId string) ([]*Project2, error) {
	var req struct {
		UserId string `json:"userId"  binding:"-"`
	}
	req.UserId = config.SharePrivateConfigInstance().UserId
	resp := http.POSTJson(fmt.Sprintf(config.SharePrivateConfigInstance().Uri+"project2/list"), req)
	if resp.StatusCode != http2.StatusOK {
		return nil, fmt.Errorf("list project error", resp.Error())
	}

	var proj []*Project2
	if err := resp.Resp(&proj); err != nil {
		return nil, fmt.Errorf("list project error:%s %s", err.Error(), string(resp.Byte()))
	}

	return proj, nil
}
