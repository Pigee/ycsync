package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func main() {

	fmt.Println("Hello.Go..")
	var dbc DbConfig
	var DBS, DBT *sql.DB
	var instr string
	instr = "insert into tsw_meterread values"
	DBS, DBT = dbc.initcon()

	rows, e := DBS.Query("select t1.id,t3.name AreaName,t1.kh ClientNo,t1.hm ClientName,t2.watermeter_no MeterSeal,t2.address MeterAddress,t5.watert_name FeeKind,now() CreateDate from ys_cust_userbase t1 join ys_cust_watermeter t2 on t1.id = t2.cust_userbase_id join ys_cb_area t3 on t1.area_no = t3.area_no join ys_cust_yspz t4 on t1.id = t4.cust_userbase_id join ys_price_watertype t5 on t4.fyhlx = t5.watert_no  limit 20")
	if e == nil {
		errors.New("query incur error")
	}
	//获取列名cols
	cols, _ := rows.Columns()
	fmt.Println(rows)
	if len(cols) > 0 {

		DBS, DBT = dbc.initcon()
		var ret []map[string]string
		for rows.Next() {

			fmt.Println(rows)
			instr = strings.Join([]string{instr, "("}, "")
			buff := make([]interface{}, len(cols))
			data := make([][]byte, len(cols)) //数据库中的NULL值可以扫描到字节中
			for i, _ := range buff {
				buff[i] = &data[i]
			}
			rows.Scan(buff...) //扫描到buff接口中，实际是字符串类型data中
			dataKv := make(map[string]string, len(cols))
			for k, col := range data { //k是index，col是对应的值
				//            fmt.Printf("%30s:\t%s\n", cols[k], col)
				//				fmt.Printf("'%s',", col)
				instr = strings.Join([]string{instr, "'", string(col), "',"}, "")
				dataKv[cols[k]] = string(col)
			}
			ret = append(ret, dataKv)

			instr = strings.Join([]string{instr[0 : len(instr)-1], "),"}, "")
		}
		//  return ret
		//         fmt.Println(ret)
	} else {
		// return nil
	}

	instr = strings.Join([]string{instr[0 : len(instr)-1], ";"}, "")

	fmt.Println(len(instr))
	fmt.Println(instr)
	// fmt.Printf(instr)
	fmt.Printf("\n")

	// fmt.Println(rows)

	DBS.Close()
	DBT.Close()
}
