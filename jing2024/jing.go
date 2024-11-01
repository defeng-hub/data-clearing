package main

import (
	"dataclearing/conf"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

type TbJingkao struct {
	//ID                       int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Type2                    string `gorm:"type:varchar(255);comment:'职位类型'" json:"type2"`
	JobCode                  string `gorm:"type:varchar(50);comment:'职位代码'" json:"job_code"`
	UnitName                 string `gorm:"type:varchar(100);comment:'单位名称'" json:"unit_name"`
	EmployDepartment         string `gorm:"type:varchar(80);comment:'用人部门'" json:"employ_department"`
	InstituonalNatural       string `gorm:"type:varchar(100);comment:'机构性质'" json:"instituonal_natural"`
	JobName                  string `gorm:"type:varchar(80);comment:'职位名称'" json:"job_name"`
	JobLevel                 string `gorm:"type:varchar(50);comment:'职位层级'" json:"job_level"`
	JobCategory              string `gorm:"type:varchar(50);comment:'职位类别'" json:"job_category"`
	JobDescription           string `gorm:"type:varchar(255);comment:'职位简介'" json:"job_description"`
	Count                    string `gorm:"type:varchar(10);comment:'招考人数'" json:"count"`
	EducationalRequire       string `gorm:"type:varchar(80);comment:'学历要求'" json:"educational_require"`
	DegreeRequire            string `gorm:"type:varchar(80);comment:'学位要求'" json:"degree_require"`
	ProfessionalRequire      string `gorm:"type:varchar(255);comment:'专业要求'" json:"professional_require"`
	PoliticalStatus          string `gorm:"type:varchar(50);comment:'政治面貌'" json:"political_status"`
	Other                    string `gorm:"type:varchar(255);comment:'其他条件'" json:"other"`
	ProfessionalAbilityTest  string `gorm:"type:varchar(10);comment:'是否在面试阶段组织专业能力测试'" json:"professional_ability_test"`
	InterviewRatio           string `gorm:"type:varchar(10);comment:'面试人员比例'" json:"interview_ratio"`
	GrassrootsExperienceYear string `gorm:"type:varchar(50);comment:'基层最低工作年限'" json:"grassroots_experience_year"`
	EnquirtyTel              string `gorm:"type:varchar(255);comment:'考生咨询电话'" json:"enquirty_tel"`
	Website                  string `gorm:"type:varchar(1000);comment:'单位网站'" json:"website"`
	Remark                   string `gorm:"type:varchar(2000);comment:'备注'" json:"remark"`
	V2SubmitNum              string `gorm:"type:varchar(50);default:'';comment:'过审人数'" json:"v2_submit_num"`
	V2MinScore               string `gorm:"type:varchar(50);comment:'最低进面分'" json:"v2_min_score"`
	V3SubmitInfo             string `gorm:"type:varchar(4000);comment:'历史报名人数'" json:"v3_submit_info"`
	V3Ext                    string `gorm:"type:varchar(4000);comment:'拓展字段'" json:"v3_ext"`
	gorm.Model
	CreateBy int64  `gorm:"type:bigint(20);comment:'创建者'" json:"create_by"`
	UpdateBy int64  `gorm:"type:bigint(20);comment:'更新者'" json:"update_by"`
	Year     string `gorm:"type:varchar(50);comment:'年份'" json:"year"`
}

// TableName specifies the table name for GORM
func (TbJingkao) TableName() string {
	return "tb_jingkao"
}

func main() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		panic("failed to connect database")
	}
	var year = "2024"
	type2 := "普通职位" //国考更改（sheet名）
	f, err := excelize.OpenFile(fmt.Sprintf("./conf/jing/jing%s.xlsx", year))
	if err != nil {
		fmt.Println(err)
		return
	}
	var jobs []TbJingkao

	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(type2)
	for idx, row := range rows {
		if idx <= 1 {
			continue
		}
		var job TbJingkao
		job.Year = year

		job.Type2 = type2
		job.JobCode = row[1]

		job.UnitName = row[2]
		job.EmployDepartment = row[3]
		job.InstituonalNatural = row[4]
		job.JobName = row[5]
		job.JobLevel = row[6]
		job.JobCategory = row[7]
		job.JobDescription = row[8]
		job.Count = row[9]
		job.EducationalRequire = row[10]
		job.DegreeRequire = row[11]
		job.ProfessionalRequire = row[12]
		job.PoliticalStatus = row[13]
		job.Other = row[14]
		job.ProfessionalAbilityTest = row[15]
		job.InterviewRatio = row[16]
		job.GrassrootsExperienceYear = row[17]
		job.EnquirtyTel = row[18]
		job.Website = row[19]
		job.EnquirtyTel = strings.Replace(job.EnquirtyTel, "\r", "", -1)
		job.EnquirtyTel = strings.Replace(job.EnquirtyTel, "\n", ",", -1)
		job.Remark = row[20]

		jobs = append(jobs, job)
		//fmt.Printf("job:%v\n", job)

	}

	// 获取数据完成
	fmt.Println("获取数据完成")

	var count int64
	db.Table("tb_jingkao").Count(&count)
	fmt.Println("表原始长度：", count)

	//fmt.Println(jobs)
	fmt.Println("获取出来全部jobs", len(jobs))

	// 开始事务
	tx := db.Begin()

	// 滚动插入
	var temp []TbJingkao
	for _, job := range jobs {
		temp = append(temp, job)
		if len(temp) == 100 {
			err = tx.Create(&temp).Error
			if err != nil {
				fmt.Println("插入数据失败", err)
			}
			temp = temp[:0]
		}
	}
	err = tx.Create(&temp).Error
	if err != nil {
		fmt.Println("插入数据失败", err)
	}
	var count2 int64

	tx.Table("tb_jingkao").Count(&count2)

	fmt.Println("count2-count:", count2-count)

	if count2-count == int64(len(jobs)) {
		fmt.Println("提交事务:", tx.Commit().Error)
	} else {
		fmt.Println("回滚事务:", tx.Rollback().Error)
	}

}
