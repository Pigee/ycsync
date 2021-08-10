package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	//	"strings"
	"bytes"
)

func main() {

	fmt.Println("Hello.Go..")

	task := inittask()
	for _, ta := range task {
		fmt.Println(ta)
		fmt.Println("\n")
	}

	//	fmt.Println(task)

	var dbc DbConfig
	var DBS, DBT *sql.DB
	DBS, DBT = dbc.initcon()
	//	var instr strings.Builder
	var instr bytes.Buffer

	zkuo := []byte("(")
	ykuo := []byte(")")
	dhao := []byte(",")
	fhao := []byte(";")
	dyhao := []byte("'")

	fmt.Println(zkuo)
	instr.WriteString("insert into tsw_meterread values")

	rows, e := DBS.Query("select t1.id,t3.name AreaName,t1.kh ClientNo,t1.hm ClientName,t2.watermeter_no MeterSeal,t2.address MeterAddress,t5.watert_name FeeKind,now() CreateDate from ys_cust_userbase t1 join ys_cust_watermeter t2 on t1.id = t2.cust_userbase_id join ys_cb_area t3 on t1.area_no = t3.area_no join ys_cust_yspz t4 on t1.id = t4.cust_userbase_id join ys_price_watertype t5 on t4.fyhlx = t5.watert_no limit 11")
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

			if linec > 2 { // row for each
				// 	instr = strings.Join([]string{instr[0 : len(instr)-1], ";"}, "")
				instr.Truncate(instr.Len() - 1)
				instr.Write(fhao)
				fmt.Println(instr.String())
				//					fmt.Println(instr[20])
				instr.Reset()
				instr.WriteString("insert into tsw_meterread values")
				linec = 0
			}

		}
	} else {

		// return nil
	}

	if linec > 0 {
		instr.Truncate(instr.Len() - 1)
		instr.Write(fhao)
		fmt.Println(instr.String())
	}

	//	instr = strings.Join([]string{instr[0 : len(instr)-1], ";"}, "")
	fmt.Printf("\n")

	DBS.Close()
	DBT.Close()
}
