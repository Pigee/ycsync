package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"log"
)

func getdata(tsk Taskele, DBS *sql.DB, DBT *sql.DB) {

	if len(tsk.Psql) > 5 {
		DBT.Exec(tsk.Psql) //truncate target table
	}

	var instr bytes.Buffer

	zkuo := []byte("(")
	ykuo := []byte(")")
	dhao := []byte(",")
	fhao := []byte(";")
	dyhao := []byte("'")

	// 	fmt.Println(zkuo)
	instr.WriteString(tsk.Isql)

	rows, e := DBS.Query(tsk.Ssql)
	if e == nil {
		errors.New("query incur error")
	}
	//获取列名cols
	cols, _ := rows.Columns()
	//	fmt.Println(rows)
	linec := 0
	if len(cols) > 0 {

		var ret []map[string]string

		for rows.Next() {

			linec++
			instr.Write(zkuo)
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols)) //数据库中的NULL值可以扫描到字节中
			for i, _ := range buff {
				buff[i] = &data[i]
			}
			rows.Scan(buff...) //扫描到buff接口中，实际是字符串类型data中
			dataKv := make(map[string]string, len(cols))
			for k, col := range data { //k是index，col是对应的值
				instr.Write(dyhao)
				instr.Write(col)
				//instr.WriteByte(string(col))
				instr.Write(dyhao)
				instr.Write(dhao)
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)

			// instr = strings.Join([]string{instr[0 : len(instr)-1], "),"}, "")
			instr.Truncate(instr.Len() - 1)
			instr.Write(ykuo)
			instr.Write(dhao)

			if linec == tsk.Irow { // row for each
				// 	instr = strings.Join([]string{instr[0 : len(instr)-1], ";"}, "")
				instr.Truncate(instr.Len() - 1)
				instr.Write(fhao)
				//				fmt.Println(instr.String())
				DBT.Exec(instr.String())
				log.Println("insert into table", tsk.Name, linec, " rows successfully...")
				//					fmt.Println(instr[20])
				instr.Reset()
				instr.WriteString(tsk.Isql)
				linec = 0
			}

		}
	} else {

		// return nil
	}

	if linec > 0 {
		instr.Truncate(instr.Len() - 1)
		instr.Write(fhao)
		//		fmt.Println(instr.String())
		DBT.Exec(instr.String())
		log.Println("insert into table", tsk.Name, linec, " rows successfully...")
	}

	//	instr = strings.Join([]string{instr[0 : len(instr)-1], ";"}, "")
	log.Println("Refreshing table", tsk.Name, " successfully...")
	//	DBS.Close()
	//	DBT.Close()

}

func exec() {

	fmt.Println("Hello.Go..")

	var dbc DbConfig
	var DBS, DBT *sql.DB
	DBS, DBT = dbc.initcon()
	c := cron.New()

	task := inittask()

	for _, ta := range task {
		ta := ta
		c.AddFunc(ta.Cronstr, func() { getdata(ta, DBS, DBT) })

	}

	c.Start()
	select {}
	DBS.Close()
	DBT.Close()
}
