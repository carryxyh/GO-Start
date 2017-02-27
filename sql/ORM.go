package main

import (
	// "database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//注册驱动
	// orm.RegisterDriver("mysql", orm.DRMySQL)

	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/testgo?charset=utf8", 30)

	//注册定义的model
	orm.RegisterModel(new(User))

	// 创建table
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm := orm.NewOrm()

	user := User{Name: "ziyuan"}

	//插入表
	id, err := orm.Insert(&user)
	fmt.Println("Id : %d, ERR: %v \n", id, err)
	fmt.Println(user.Id)

	//更新表
	user.Name = "ZiMUEI"
	num, err := orm.Update(&user)
	fmt.Println("Num : %d,Err : %v \n", num, err)

	// 读取 one
	u := User{Id: 2}
	err = orm.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	fmt.Println("Id: ", u.Id, "Name: ", u.Name)

	//删除
	num, err = orm.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}
