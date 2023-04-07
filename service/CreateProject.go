package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateProject(args *Commond, nameSpaceID int) {
	// 创建项目
	client := &http.Client{}
	newProjectInfo := ProjectInfo{
		Name:        args.ProjectName,
		Description: args.Description,
		Path:        args.ProjectName,
		NameSpaceID: nameSpaceID,
		Visibility:  args.Visibility,
		ImportUrl:   GitLabUrl + "/init/bare.git",
	}
	dataByte, err := json.Marshal(newProjectInfo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dataByte))
	req, err := http.NewRequest("POST", GitLabUrl+"/api/v4/projects", bytes.NewReader(dataByte))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("PRIVATE-TOKEN", GitLabToken)
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}
