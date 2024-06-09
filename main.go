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

type JingJob struct {
	ID                       int64  `json:"id" gorm:"id"`
	Year                     string `json:"year" gorm:"year"`                                             // 年份
	JobCode                  string `json:"job_code" gorm:"job_code"`                                     // 京职位代码，国部门代码-职位代码
	JobName                  string `json:"job_name" gorm:"job_name"`                                     // 京职位名称，国招考职位
	UnitName                 string `json:"unit_name" gorm:"unit_name"`                                   // 京单位名称，国部门名称
	EmployDepartment         string `json:"employ_department" gorm:"employ_department"`                   // 京用人部门，国用人司局
	InstitutionalNatural     string `json:"institutional_natural" gorm:"institutional_natural"`           // 京机构性质，国机构性质
	JobLevel                 string `json:"job_level" gorm:"job_level"`                                   // 京职位层级，国机构层级
	JobCategory              string `json:"job_category" gorm:"job_category"`                             // 京职位类别，国职位属性
	JobDescription           string `json:"job_description" gorm:"job_description"`                       // 京职位简介，国职位简介
	ExamType                 string `json:"exam_type" gorm:"exam_type"`                                   // 京— — — —，国考试类别
	GrassrootsExperience     string `json:"grassroots_experience" gorm:"grassroots_experience"`           // 京— — — —，国服务基层项目工作经历
	Count                    string `json:"count" gorm:"count"`                                           // 招考人数
	EducationalRequire       string `json:"educational_require" gorm:"educational_require"`               // 学历要求
	DegreeRequire            string `json:"degree_require" gorm:"degree_require"`                         // 学位要求
	ProfessionalRequire      string `json:"professional_require" gorm:"professional_require"`             // 专业要求
	PoliticalStatus          string `json:"political_status" gorm:"political_status"`                     // 政治面貌
	Other                    string `json:"other" gorm:"other"`                                           // 其他条件
	ProfessionalAblilityTest string `json:"professional_ablility_test" gorm:"professional_ablility_test"` // 是否组织专业能力测试
	InterviewRatio           string `json:"interview_ratio" gorm:"interview_ratio"`                       // 面试人员比例 5:1
	GrassrootsExperienceYear string `json:"grassroots_experience_year" gorm:"grassroots_experience_year"` // 基层工作最低年限
	EnquirtyTel              string `json:"enquirty_tel" gorm:"enquirty_tel"`                             // 咨询电话
	Website                  string `json:"website" gorm:"website"`                                       // 单位网站
	Remark                   string `json:"remark" gorm:"remark"`                                         // 备注
	Type1                    string `json:"type1" gorm:"type1"`                                           // 国考，京考
	Type2                    string `json:"type2" gorm:"type2"`                                           // 京考：普通职位，面向大学生士兵职位，面向乡村振兴协理员等服务基层项目人员职位，面向残疾人职位，特殊职位\n国考：中央党群机关，中央国家行政机关（本级），中央国家行政机关省级以下直属机构，中央国家行政机关参照公务员法管理事业单位
}

// TableName 表名称
func (*JingJob) TableName() string {
	return "job"
}

func main() {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		panic("failed to connect database")
	}

	f, err := excelize.OpenFile("./conf/jing.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var jobs []JingJob

	type1 := "京考"
	type2 := "特殊职位" //国考更改（sheet名）
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(type2)
	for idx, row := range rows {
		if idx <= 1 {
			continue
		}
		var job JingJob
		job.Type1 = type1
		job.Type2 = type2
		job.JobCode = row[1]

		job.UnitName = row[2]
		job.EmployDepartment = row[3]
		job.InstitutionalNatural = row[4]
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
		job.ProfessionalAblilityTest = row[15]
		job.InterviewRatio = row[16]
		job.GrassrootsExperienceYear = row[17]
		job.EnquirtyTel = row[18]
		job.Website = row[19]
		job.Remark = row[20]
		job.EnquirtyTel = strings.Replace(job.EnquirtyTel, "\r", "", -1)
		job.EnquirtyTel = strings.Replace(job.EnquirtyTel, "\n", ",", -1)

		job.ExamType = "-"
		job.GrassrootsExperience = "-"

		job.Year = "2024"
		jobs = append(jobs, job)
		//fmt.Printf("job:%v\n", job)

	}

	// 获取数据完成
	fmt.Println("获取数据完成")

	var count int64
	db.Table("job").Count(&count)
	fmt.Println("表原始长度：", count)

	//fmt.Println(jobs)
	fmt.Println("获取出来全部jobs", len(jobs))

	// 开始事务
	tx := db.Begin()

	// 滚动插入
	var temp []JingJob
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

	tx.Table("job").Count(&count2)

	fmt.Println("count2-count:", count2-count)

	if count2-count == int64(len(jobs)) {
		fmt.Println("提交事务:", tx.Commit().Error)
	} else {
		fmt.Println("回滚事务:", tx.Rollback().Error)
	}

}
