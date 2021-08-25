package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/godror/godror"
	_ "github.com/mattn/go-adodb"
	"io/ioutil"
	"os"
	"strings"
)

type Task []struct {
	Name    string `json:"name"`
	Ssql    string `json:"ssql"`
	Isql    string `json:"isql"`
	Psql    string `json:"psql"`
	Irow    int    `json:"irow"`
	Cronstr string `json:"cronstr"`
}

type Taskele struct {
	Name    string `json:"name"`
	Ssql    string `json:"ssql"`
	Isql    string `json:"isql"`
	Psql    string `json:"psql"`
	Irow    int    `json:"irow"`
	Cronstr string `json:"cronstr"`
}

type DbConfig struct {
	Source struct {
		Db       string `json:"db"`
		Type     string `json:"type"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		State    int    `json:"state"`
	} `json:"source"`
	Target struct {
		Db       string `json:"db"`
		Type     string `json:"type"`
		IP       string `json:"ip"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		State    int    `json:"state"`
	} `json:"target"`
}

func inittask() Task {

	jsonFile, err := os.Open("task.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tsk Task
	json.Unmarshal([]byte(byteValue), &tsk)
	//	fmt.Println(tsk)

	fmt.Println("init Task success")
	return tsk

}

func (DbConfig) initcon() (*sql.DB, *sql.DB) {

	jsonFile, err := os.Open("db.json")

	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}

	// 要记得关闭
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var DBs, DBt *sql.DB
	var dbc DbConfig
	json.Unmarshal([]byte(byteValue), &dbc)
	fmt.Println(dbc.Source.State)

	switch dbc.Source.Type {
	case "mysql":
		DBs = initmysql(dbc.Source.User, dbc.Source.Password, dbc.Source.IP, dbc.Source.Port, dbc.Source.Db)
	case "oracle":
		DBs = initmysql(dbc.Source.User, dbc.Source.Password, dbc.Source.IP, dbc.Source.Port, dbc.Source.Db)
	case "mssql":
		DBs = initmysql(dbc.Source.User, dbc.Source.Password, dbc.Source.IP, dbc.Source.Port, dbc.Source.Db)
	default:
		DBs = initmysql(dbc.Source.User, dbc.Source.Password, dbc.Source.IP, dbc.Source.Port, dbc.Source.Db)
	}

	switch dbc.Target.Type {
	case "mysql":
		DBt = initmysql(dbc.Target.User, dbc.Target.Password, dbc.Target.IP, dbc.Target.Port, dbc.Target.Db)
	case "oracle":
		DBt = initmysql(dbc.Target.User, dbc.Target.Password, dbc.Target.IP, dbc.Target.Port, dbc.Target.Db)
	case "mssql":
		DBt = initmysql(dbc.Target.User, dbc.Target.Password, dbc.Target.IP, dbc.Target.Port, dbc.Target.Db)
	default:
		DBt = initmysql(dbc.Target.User, dbc.Target.Password, dbc.Target.IP, dbc.Target.Port, dbc.Target.Db)
	}

	return DBs, DBt
}

func initmysql(user, password, ip, port, db string) *sql.DB {

	paths := strings.Join([]string{user, ":", password, "@tcp(", ip, ":", port, ")/", db, "?charset=utf8"}, "")
	Mysqldb, _ := sql.Open("mysql", paths)
	Mysqldb.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	Mysqldb.SetMaxIdleConns(10)
	//验证连接
	if err := Mysqldb.Ping(); err != nil {
		fmt.Println("open ",db," database fail")
		return nil
	}
	fmt.Println("connnect ",db," success")
	return Mysqldb
}
