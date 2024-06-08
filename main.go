package main

import (
	"dataclearing/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) <= 2 {
		fmt.Println("软件使用方式：check <产品id> <产品名称>")
		return
	}

	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Table("p_product").Where("id = ?", os.Args[1])

}
