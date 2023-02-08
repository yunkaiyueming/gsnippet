package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
  "time"
  "errors"
)

var engine *xorm.Engine

func DBInit(){
	//获取conf.ini-mysql配置
	// // sqlConfig, err := config.File.GetSection("mysql")
	// if err != nil {
	// 	fmt.Println("数据库配置获取出错",err)
	// 	panic(err)
	// }
  sqlConfig:=map[string]string{
    "user":"root",
    "password":"root",
    "host":"127.0.0.1",
    "port":"3306",
    "dbname":"test",
  }

	// userName:password@tcp(host:port)/dbname?charset=utf8&parseTime=True&loc=Local
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		sqlConfig["user"],
		sqlConfig["password"],
		sqlConfig["host"],
		sqlConfig["port"],
		sqlConfig["dbname"],
	)
	//新建xorm单表引擎
	err := errors.New("")
	engine, err = xorm.NewEngine("mysql", dbConn)
	if err != nil{
		fmt.Println("数据库连接失败",err)
		panic(err)
	}

	//测试数据库是否连接成功
	err = engine.Ping()
	if err != nil{
		fmt.Println("数据库ping失败",err)
		panic(err)
	}

	//最大空闲连接数
	// maxIdle := config.File.MustInt("mysql", "max_idle", 2)
	//最大打开连接数
	// maxConn := config.File.MustInt("mysql", "max_conn", 10)

	engine.SetMaxIdleConns(100)
	engine.SetMaxOpenConns(100)
	fmt.Println("数据库初始化成功！")
}

type User struct{
  Id int64 `xorm:"id pk autoincr"`
  Name string `xorm:"name varchar(200)"`
  Salt string `xorm:"salt varchar(200)"`
  Age int `xorm:"age int(11)"`
  Passwd string `xorm:"passwd varchar(200)"`
  Created time.Time `xorm:"created datetime"`
  Updated time.Time `xorm:"updated datetime"`
  Info string `xorm:"info varchar(200)"`
  ServantInfo string `xorm:"sinfo varchar(200)"`
  MapServant map[string]Servant `xorm:"servant_info varchar(500)"`
}

type Servant struct{
	lv int64 
	exp int64 
	aura string
}

// CREATE TABLE IF NOT EXISTS `user` (
// 	`id` int(10) unsigned NOT NULL,
// 	`name` varchar(200) NOT NULL DEFAULT '',
// 	`salt` varchar(200) NOT NULL DEFAULT '',
// 	`age` int(11) NOT NULL DEFAULT '',
// 	`passwd` varchar(200) NOT NULL DEFAULT '',
// 	`created` datetime NOT NULL DEFAULT '0',
// 	`updated` datetime NOT NULL DEFAULT '0',
// 	`info` varchar(200) NOT NULL DEFAULT '0',
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

func (user *User) GetUserinfo(uid int){
  results, err := engine.QueryString("select * from user where Id = ?", 8)
  fmt.Println("获取数据",results,err)
  for k, v := range results {
    fmt.Println(k, v, fmt.Sprintf("%T", v))
}
}

func (user *User) UpdateExecStr(){
  affected, err := engine.Exec("update user set name = ? where id = ?", "aa1", 1)
  fmt.Println(affected,err)
}

func (user *User) InsertUser(){
  affected, err := engine.Insert(user)
  fmt.Println(affected,err)
}

func (user *User) getAll(){
  var users []User
	err := engine.Where("name like ?", "%aa%").And("age > 1").Limit(10, 0).Find(&users)
// SELECT * FROM user WHERE name = ? AND age > 10 limit 10 offset 0
	fmt.Println(users,err)
}


func main(){
	DBInit()
  	err := engine.Sync(new(User))
	fmt.Println(err)

	firstServant := Servant{lv:1,exp:1,aura:"{}"}
	user1:=User{
		Name:"aa",
		Salt:"12312",
		Age:10,
		Passwd:"jodfsajfosda",
		Created:time.Now(),
		Updated:time.Now(),
		Info:"{}",
		ServantInfo:"{}",
		MapServant: map[string]Servant{"1001":firstServant},
	}

	user1.InsertUser()

  	new(User).GetUserinfo(1)
	new(User).UpdateExecStr()
	new(User).getAll()

}