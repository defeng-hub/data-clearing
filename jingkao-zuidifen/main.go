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

	var year = "2022"
	type2 := "Sheet1" //国考更改（sheet名）
	f, err := excelize.OpenFile(fmt.Sprintf("./jingkao-zuidifen/zuidifen-%s.xlsx", year))
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(type2)
	for idx, row := range rows {
		if idx <= 1 {
			continue
		}
		if row[3] != "" {
			fmt.Println(row[1], row[2], row[3])
			if row[5] != "" {
				db.Table("tb_jingkao").
					Where("year = ?", year).
					Where(fmt.Sprintf("unit_name = ? && employ_department = ? && (job_level LIKE %v || job_name LIKE %v)", "'%"+row[3]+"%'", "'%"+row[3]+"%'"), row[1], row[2]).
					Update("v2_min_score", row[5])
			}
		} else {
			fmt.Println(row[1], row[2])
			db.Table("tb_jingkao").
				Where("year = ?", year).
				Where("unit_name = ? && employ_department = ?", row[1], row[2]).
				Update("v2_min_score", row[5])
		}
		//db.Table("tb_jingkao").Where("year = ? && job_code = ?", year, row[0]).Update("v2_min_score", row[1])
	}
}
