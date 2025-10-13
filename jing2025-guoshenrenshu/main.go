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

func main() {
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
	fmt.Println(db)

	var year = "2025"
	type2 := "Data"
	f, err := excelize.OpenFile("./jing2025-guoshenrenshu/原始数据.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(type2)
	for idx, row := range rows {
		if idx <= 0 {
			continue
		}
		fmt.Println("row", year, row[1], row[6])
		db.Table("tb_jingkao").Where("year = ? && job_code = ?", year, row[1]).Update("v2_submit_num", fmt.Sprintf("%v人，%v", row[6], "截止到11月22日9点"))
	}
}
