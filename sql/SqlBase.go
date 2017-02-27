package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//这里之前做的准备工作：go get 下载github.com/go-sql-driver/mysql，然后把这个包拷贝到goroot下面。。很奇怪
//这下面的代码是不能直接运行的，会出问题，每次的stmt、res都要重新声明

//db.Prepare()函数用来返回准备要执行的sql操作，然后返回准备完毕的执行状态。

//db.Query()函数用来直接执行Sql返回Rows结果。

//stmt.Exec()函数用来执行stmt准备好的SQL语句

func main() {

	//sql.Open 支持以下格式
	//user@unix(/path/to/socket)/dbname?charset=utf8
	//user:password@tcp(localhost:5555)/dbname?charset=utf8
	//user:password@/dbname
	//user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	db, err := sql.Open("mysql", "root:root@/testgo?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("ziyuan", "研发部门", "2017-02-27")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt1, err := db.Prepare("update userinfo set username=?,departname=? where uid=?")
	checkErr(err)

	res1, err := stmt1.Exec("ZiYuan", "后端开发", 1)
	checkErr(err)

	affect, err := res1.RowsAffected()
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(3)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
