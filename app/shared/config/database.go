package config

import (
	"fmt"

	"github.com/gomodul/envy"
)

type databaseItem struct {
	DriverName     string
	DataSourceName string
}

type database struct {
	MySQL databaseItem
}

var DB = database{
	MySQL: databaseItem{
		DriverName: envy.Get("MYSQL_DRIVER"),
		DataSourceName: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8",
			envy.Get("MYSQL_USERNAME", "root"),
			envy.Get("MYSQL_PASSWORD", "abcd"),
			envy.Get("MYSQL_HOST", "localhost"),
			envy.Get("MYSQL_PORT", "3306"),
			envy.Get("MYSQL_NAME", "hacktiv"),
		),
	},
}
