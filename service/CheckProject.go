package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func CheckProject(projectname string) {
	// 检查 ProjectName 是否符合规则
	if matched, _ := regexp.MatchString("^[a-z0-9-]+$", projectname); !matched {
		fmt.Fprintf(os.Stderr, "Error: ProjectName 只能使用小写字母、数字或-\n")
		os.Exit(11)
	}

	// 检查 ProjectName 是否存在
	client := &http.Client{}
	req, err := http.NewRequest("GET", GitLabUrl+"/api/v4/projects?search="+projectname, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(12)
	}

	req.Header.Add("PRIVATE-TOKEN", GitLabToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	projectinfo := TempInfo{}
	json.Unmarshal(body, &projectinfo)

	if len(projectinfo) != 0 {
		fmt.Println("ERROR: 已存在此项目！")
		os.Exit(14)
	}
}
