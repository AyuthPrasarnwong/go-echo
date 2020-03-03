package bootstrap

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// _ defined non use
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type (
	// MySQL database management
	MySQL struct {
	}
)

// dbsMySQL variable for define connection
var dbsMySQL = make(map[string]*gorm.DB)

// CreateMySQLConnection make connection
func CreateMySQLConnection() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "jhi_" + defaultTableName
	// }
	for k, v := range viper.Get("database").(map[string]interface{}) {
		x := v.(map[string]interface{})
		port := 3306
		if val, found := x["port"]; found {
			port = val.(int)
		}
		host := x["host"].(string)

		switch driver := x["driver"].(string); driver {
		case "mysql":
			connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
				x["username"].(string),
				x["password"].(string),
				host,
				port,
				x["db"].(string),
			)
			db, err := gorm.Open("mysql", connection)
			if err != nil {
				panic(fmt.Sprintf("failed to connect database of %s connection", k))
			}
			db.LogMode(viper.GetBool("debug"))
			// defer db.Close()
			db.DB().SetMaxIdleConns(10)

			dbsMySQL[k] = db
		default:
		}
	}
}

// DB get mysql connection
func (ctl *MySQL) DB(x interface{}) *gorm.DB {
	if x == nil {
		return dbsMySQL["default"]
	}
	if connection, found := dbsMySQL[x.(string)]; found {
		return connection
	}
	panic(fmt.Sprintf("connection %s not found", x.(string)))
}
