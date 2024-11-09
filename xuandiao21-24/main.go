package main

import (
	"dataclearing/conf"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func dataclean2021() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 设定日志级别为 Info（也可以是 Debug）
				IgnoreRecordNotFoundError: true,
				Colorful:                  true, // 启用彩色打印
			},
		),
	})
	if err != nil {
		panic("failed to connect database")
	}

	var year = "2021"
	f, err := excelize.OpenFile("./xuandiao21-24/dxxd" + year + "-最低分.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("aaa")
	for i, row := range rows {
		if i <= 2 {
			continue
		}
		fmt.Println(row[0], row[1])
		db.Table("tb_dingxiangxuandiao").
			Where("year = ? && zhiwei_daima = ?", year, row[0]).
			Update("v2_min_score", row[1])
	}

}
func dataclean2022() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 设定日志级别为 Info（也可以是 Debug）
				IgnoreRecordNotFoundError: true,
				Colorful:                  true, // 启用彩色打印
			},
		),
	})

	fmt.Println(db)
	if err != nil {
		panic("failed to connect database")
	}

	var year = "2022"
	f, err := excelize.OpenFile("./xuandiao21-24/dxxd" + year + "-最低分.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("bbb")
	for i, row := range rows {
		if i <= 2 {
			continue
		}
		//fmt.Println(row[1], row[3])
		db.Table("tb_dingxiangxuandiao").
			Where("year = ? && zhiwei_daima = ?", year, row[1]).
			Update("v2_min_score", row[3])
	}

}

func main() {
	dataclean2022()
}
