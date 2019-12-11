package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
	"github.com/nameouyang/learning-go/conf"
	"strconv"
	"time"
)

var DBConnect *gorm.DB

func init() {
	//db, err := gorm.Open("mysql", "userName:password@127.0.0.1:3306/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()
	//db 链接参数
	var connectStr = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBConf.User,
		conf.DBConf.Password,
		conf.DBConf.Host+":"+strconv.Itoa(conf.DBConf.Port),
		conf.DBConf.DBName,
	)
	//fmt.Println(connectStr)
	var err error
	DBConnect, err = gorm.Open(conf.DBConf.DBType, connectStr)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		time.Sleep(10 * time.Second) // 若连接失败，则延时10秒重新连接
		DBConnect, err = gorm.Open(conf.DBConf.DBType, connectStr)
		if err != nil {
			panic(err.Error())
		}
	}
	//root:root@tcp(127.0.0.1:3306)/db_monitor?charset=utf8&parseTime=True&loc=Local
	//root:root@tcp(127.0.0.1:3306)/db_monitor?charset=utf8&parseTime=True&loc=Local

	if DBConnect.Error != nil {
		fmt.Printf("database error %v", DBConnect.Error)
	}
	//更改默认表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DBConf.TablePrefix + defaultTableName
	}
	DBConnect.LogMode(conf.DBConf.Debug)
	// 全局禁用表名复数 例如 user 会变成 users
	DBConnect.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
	DBConnect.DB().SetMaxIdleConns(conf.DBConf.MaxIdle)
	DBConnect.DB().SetMaxOpenConns(conf.DBConf.MaxOpen)
	//数据迁移使用的
	/*DBConnect.Set(
		"gorm:table_options", //创建表时添加表后缀
		"ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci",
	).AutoMigrate(&User{}, &Task{})
	DBConnect.Model(&User{}).AddUniqueIndex("uk_email", "email")
	*/
}
