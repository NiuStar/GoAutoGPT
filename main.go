package main

import (
	"GoCodeGPT/app"
	"github.com/urfave/cli/v3"
	"os"
)

func main() {
	cApp := &cli.App{}
	cApp.Name = "Golang服务生成工具"
	cApp.Usage = "代码自动化生成工具，自动生成api接口、client、数据层代码"
	cApp.Version = "1.0.0"
	var (
		projectName string
	)
	cApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			Usage:       "项目名称",
			Aliases:     []string{"p"},
			Destination: &projectName,
		},
	}
	cApp.Commands = append(cApp.Commands, CreateCommand(), GenerateCommand(), ListCommand())
	cApp.Run(os.Args)
}
func GenerateCommand() *cli.Command {
	cmd := &cli.Command{}
	cmd.Aliases = []string{"g"}
	cmd.Name = "generate"
	cmd.Usage = "生成代码"
	var (
		projectId string
	)
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "projectId",
			Usage:       "项目Id",
			Aliases:     []string{"p"},
			Destination: &projectId,
		},
	}
	cmd.Action = func(c *cli.Context) error {
		return app.Genarate(projectId)
	}
	return cmd
}

func CreateCommand() *cli.Command {
	cmd := &cli.Command{}
	cmd.Aliases = []string{"c"}
	cmd.Name = "create"
	cmd.Usage = "创建项目"
	var (
		projectName string
		description string
	)
	cmd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "project",
			Usage:       "项目名称",
			Aliases:     []string{"p"},
			Destination: &projectName,
		},
		&cli.StringFlag{
			Name:        "description",
			Usage:       "项目概述",
			Aliases:     []string{"d"},
			Destination: &description,
		},
	}
	cmd.Action = func(c *cli.Context) error {
		return app.Create(projectName, description)
	}
	return cmd
}

func ListCommand() *cli.Command {
	cmd := &cli.Command{}
	cmd.Aliases = []string{"l"}
	cmd.Name = "list"
	cmd.Usage = "项目列表"
	cmd.Action = func(c *cli.Context) error {
		return app.List()
	}
	return cmd
}
