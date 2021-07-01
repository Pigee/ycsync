package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"strings"
)

type DbConfig struct {
	Source struct {
		Db       string `json:"db"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		State    int    `json:"state"`
	} `json:"source"`
	Target struct {
		Db       string `json:"db"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		State    int    `json:"state"`
	} `json:"target"`
}

func (DbConfig) initcon() (*sql.DB,*sql.DB){

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
	fmt.Println(dbc.Source.State)

	paths := strings.Join([]string{dbc.Source.User, ":", dbc.Source.Password, "@tcp(", dbc.Source.IP, ":", dbc.Source.Port, ")/", dbc.Source.Db, "?charset=utf8"}, "")
	DBs, _ := sql.Open("mysql", paths)
	DBs.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DBs.SetMaxIdleConns(10)
	//验证连接
	if err := DBs.Ping(); err != nil {
		fmt.Println("open DBs database fail")
   		return nil,nil
	}
	fmt.Println("connnect Dbs success")

	patht := strings.Join([]string{dbc.Target.User, ":", dbc.Target.Password, "@tcp(", dbc.Target.IP, ":", dbc.Target.Port, ")/", dbc.Target.Db, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DBt, _ := sql.Open("mysql", patht)
	DBt.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DBt.SetMaxIdleConns(10)
	//验证连接
	if err := DBt.Ping(); err != nil {
		fmt.Println("open Dbt database fail")
		return nil,nil
	}
	fmt.Println("connnect DBt success")
	
	return DBs,DBt
}
