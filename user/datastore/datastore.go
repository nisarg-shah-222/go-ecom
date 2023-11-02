package datastore

import (
	"fmt"
	"os"
)

func Get() {
	if MySQLConn == nil {
		err := ConnectMySQL()
		if err != nil {
			fmt.Println("Error connecting to MySQL:", err)
			os.Exit(1)
		}
	}
}
