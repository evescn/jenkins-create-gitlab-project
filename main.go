package main

import (
	"create_gitlab_project/service"
	"flag"
	"fmt"
	"os"
)

func main() {
	c := new(service.Commond)
	flag.StringVar(&c.GroupName, "groupname", "", "GitLab 组名 (required)")
	flag.StringVar(&c.ProjectName, "projectname", "", "GitLab 项目名 (required)")
	flag.StringVar(&c.Visibility, "visibility", "", "GitLab 项目权限 (required)")
	flag.StringVar(&c.Description, "desc", "", "GitLab 项目描述")

	flag.Parse()

	// 判断传参是否正确
	if c.GroupName == "" || c.ProjectName == "" || c.Visibility == "" {
		flag.Usage()
		os.Exit(1)
	}

	// 检查 GroupName 是否存在
	groupID := service.GetGroupNameID(c.GroupName)
	if groupID == -1 {
		fmt.Printf("ERROR: GitLab 不存在 %s 的 GroupName\n", c.GroupName)
		os.Exit(-2)
	}

	// 检查 Visibility
	service.CheckVisibility(c.Visibility)

	// 检查 ProjectName 是否存在
	service.CheckProject(c.ProjectName)

	// 创建项目
	service.CreateProject(c, groupID)

}
