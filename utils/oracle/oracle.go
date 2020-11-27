package oracle

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"time"
)

//https://blogs.oracle.com/developers/how-to-connect-a-go-program-to-oracle-database-using-goracle
//https://oracle.github.io/odpi/doc/installation.html#id3
//ln -s /opt/oracle/instantclient_19_8/libclntsh.dylib /usr/local/lib/

func TestCon() (*sql.DB, error) {
	db, err := sql.Open("godror", "apps/apps@eb-arp-dev-fah.otr.ru:1529/fah")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	//defer db.Close()

	rows, err := db.Query("select doc_guid from XXT_BS_DOC_REGISTRY where rownum = 1")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
	return db, nil
}
