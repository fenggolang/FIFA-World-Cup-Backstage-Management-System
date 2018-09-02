package init

import (
	"FIFA-World-Cup-Backstage-Management-System/infra/config"
	"fmt"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" //GORM已经包装了一些驱动程序，以便更容易记住驱动程序的导入路径,这里导入postgres数据库驱动
)

var POSTGRES *gorm.DB

func init() {
	// db init
	connectString := config.GetPostgreConfig()
	fmt.Println(connectString)
	connect, err := gorm.Open(
		"postgres",
		connectString,
	)
	connect.LogMode(true)
	if err != nil {
		fmt.Println("connect postgres database failed!", err)
		return
	}
	fmt.Println("login postgres database success!")
	POSTGRES = connect
}
