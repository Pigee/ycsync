package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func main() {

	fmt.Println("Hello.Go..")
	var dbc DbConfig
	var DBS,DBT *sql.DB
	DBS,DBT = dbc.initcon()

	rows, e := DBS.Query("select t1.id,t3.name AreaName,t1.kh ClientNo,t1.hm ClientName,t2.watermeter_no MeterSeal,t2.address MeterAddress,t5.watert_name FeeKind,now() CreateDate from ys_cust_userbase t1 join ys_cust_watermeter t2 on t1.id = t2.cust_userbase_id join ys_cb_area t3 on t1.area_no = t3.area_no join ys_cust_yspz t4 on t1.id = t4.cust_userbase_id join ys_price_watertype t5 on t4.fyhlx = t5.watert_no limit 1")
	if e == nil {
        errors.New("query incur error")
    }
	fmt.Println(rows)

	DBS.Close()
	DBT.Close()
}
