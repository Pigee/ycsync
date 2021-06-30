package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DbConfig struct {
	Source struct {
		Db       string `json:"db"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"source"`
	Target struct {
		Db       string `json:"db"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"target"`
}

func main() {

	jsonFile, err := os.Open("db.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dbc DbConfig
	json.Unmarshal([]byte(byteValue), &dbc)
	fmt.Println(dbc)

}
