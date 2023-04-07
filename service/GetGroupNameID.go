package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetGroupNameID(groupname string) int {
	client := &http.Client{}

	// 通过API查询group列表，获取已存在的groupname和对应的groupid，并根据输入的groupname返回对应的groupid，如果不存在输入的groupname，返回-1
	req, err := http.NewRequest("GET", GitLabUrl+"/api/v4/groups?search="+groupname, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Add("PRIVATE-TOKEN", GitLabToken)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	groupinfo := TempInfo{}
	json.Unmarshal(body, &groupinfo)

	if len(groupinfo) == 0 {
		return -1
	} else {
		return groupinfo[0].ID
	}
}
