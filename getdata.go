package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	//"strings"
)

type TaskList []struct {
	Name      string `json:"name"`
	Ssql      string `json:"ssql"`
	Isql      string `json:"isql"`
	Irow      int    `json:"irow"`
	Startdate string `json:"startdate"`
	Interval  int    `json:"interval"`
}

func (TaskList) initTask() (re string) {

	jsonFile, err := os.Open("task.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var ts TaskList
	json.Unmarshal([]byte(byteValue), &ts)
	fmt.Println(ts[0].Isql)
	fmt.Println(len(ts))
	return "ok"

}
