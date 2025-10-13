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

var year = "2021"
var tableName = "tb_guokao"

// Job undefined
type GuoJob struct {
	Year        string `json:"year" gorm:"year"`              // 年份
	JobCodeBase string `json:"job_code_base" gorm:"job_code"` // 京— — — —，国部门代码
	JobCode     string `json:"job_code" gorm:"job_code"`      // 京职位代码，国职位代码
	// job_distribute   			职位分布
	// grassroots_experience_year 	服务基层工作年龄
	// work_space 					工作地点
	// settled_space 				落户地点
	JobDistribute            string `json:"job_distribute" gorm:"job_distribute"`                         // 职位分布
	GrassrootsExperienceYear string `json:"grassroots_experience_year" gorm:"grassroots_experience_year"` // 服务基层工作年龄
	WorkSpace                string `json:"work_space" gorm:"work_space"`                                 // 工作地点
	SettledSpace             string `json:"settled_space" gorm:"settled_space"`                           // 落户地点
}

// TableName 表名称
func (*GuoJob) TableName() string {
	return tableName
}
func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 使用标准输出
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别：Silent, Error, Warn, Info
			Colorful:      true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{
		Logger: newLogger,
	})

	f, err := excelize.OpenFile("./conf/guo/2021guo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var jobs []GuoJob

	// 获取 Sheet1 上所有单元格
	rows := f.GetRows("中央国家行政机关省级以下直属机构")
	for idx, row := range rows {
		if idx <= 1 {
			continue
		}
		var job GuoJob
		job.Year = year
		job.JobCodeBase = row[0]
		job.JobCode = row[8]

		job.JobDistribute = row[6]
		job.WorkSpace = row[20]
		job.SettledSpace = row[21]
		job.GrassrootsExperienceYear = row[16]

		jobs = append(jobs, job)
	}
	fmt.Println("将要插入jobs个数:", len(jobs))

	var count int64
	db.Table(tableName).Count(&count)
	fmt.Println("表原始长度：", count)

	for _, job := range jobs {
		if job.Year == "" || job.JobCodeBase == "" || job.JobCode == "" {
			continue
		}
		db.Table("tb_guokao").Where("year = ?", job.Year).
			Where("job_code_base = ?", job.JobCodeBase).
			Where("job_code = ?", job.JobCode).
			Updates(map[string]interface{}{
				"job_distribute":             job.JobDistribute,
				"grassroots_experience_year": job.GrassrootsExperienceYear,
				"work_space":                 job.WorkSpace,
				"settled_space":              job.SettledSpace,
			})
		//fmt.Println("ccc")
		//break
	}
}
