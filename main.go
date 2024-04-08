package main

import (
	_ "apiproject/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

/*
DB_PORT=3318
DB_CONNECTION=mysql
DB_PASSWORD=1QAZ2wsx!
DB_USERNAME=dev
DB_HOST=49.234.16.160
DB_DATABASE=yehwang
*/

func init() {
	// 这里注册一个default默认数据库，数据库驱动是mysql.
	// 第三个参数是数据库dsn, 配置数据库的账号密码，数据库名等参数
	// dsn参数说明：
	// username - mysql账号
	// password - mysql密码
	// db_name - 数据库名
	// 127.0.0.1:3306 - 数据库的地址和端口
	orm.RegisterDataBase("default", "mysql", "dev:1QAZ2wsx!@tcp(49.234.16.160:3318)/yehwang?charset=utf8&parseTime=true&loc=Local")
	// 打开调试模式，开发的时候方便查看orm生成什么样子的sql语句
	orm.Debug = true
}

func main() {
	web.Run()
}
