package main

import (
	"fmt"

	"dataclearing/conf"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var year = "2026"
var tableName = "tb_guokao"

// Job undefined
type GuoJob struct {
	ID                       int64  `json:"id" gorm:"id"`
	Year                     string `json:"year" gorm:"year"`                                             // 年份
	JobCodeBase              string `json:"job_code_base" gorm:"job_code"`                                // 京— — — —，国部门代码
	JobCode                  string `json:"job_code" gorm:"job_code"`                                     // 京职位代码，国职位代码
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

	WorkSpace    string `json:"work_space" gorm:"work_space"`
	SettledSpace string `json:"settled_space" gorm:"settled_space"`

	GrassrootsExperienceYear string `json:"grassroots_experience_year" gorm:"grassroots_experience_year"` // 基层工作最低年限
	EnquirtyTel              string `json:"enquirty_tel" gorm:"enquirty_tel"`                             // 咨询电话
	Website                  string `json:"website" gorm:"website"`                                       // 单位网站
	Remark                   string `json:"remark" gorm:"remark"`                                         // 备注
	Type1                    string `json:"type1" gorm:"type1"`                                           // 国考，京考
	Type2                    string `json:"type2" gorm:"type2"`                                           // 京考：普通职位，面向大学生士兵职位，面向乡村振兴协理员等服务基层项目人员职位，面向残疾人职位，特殊职位\n国考：中央党群机关，中央国家行政机关（本级），中央国家行政机关省级以下直属机构，中央国家行政机关参照公务员法管理事业单位

	//v_submit_num
	//v_max_score
	//v_min_score
	V2SubmitNum string `json:"v_submit_num" gorm:"v2_submit_num"`
	V2MaxScore  string `json:"v_max_score" gorm:"v2_max_score"`
	V2MinScore  string `json:"v_min_score" gorm:"v2_min_score"`
}

// TableName 表名称
func (*GuoJob) TableName() string {
	return tableName
}
func main() {

	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect database")
	}

	f, err := excelize.OpenFile("./conf/guo/2026guo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	var jobs []GuoJob

	type1 := "国考"
	type2 := "中央党群机关"
	// 获取 Sheet1 上所有单元格
	rows := f.GetRows(type2)
	for idx, row := range rows {
		if idx <= 1 {
			continue
		}
		var job GuoJob
		job.Type1 = type1
		job.Type2 = type2
		//job.JobCode = row[0]
		job.UnitName = row[1]
		job.EmployDepartment = row[2]
		job.InstitutionalNatural = row[3]
		job.JobName = row[4]
		job.JobCategory = row[5]
		job.JobDescription = row[7]
		job.JobCode = row[8]
		job.JobCodeBase = row[0]
		job.JobLevel = row[9]
		job.ExamType = row[10]
		job.Count = row[11]
		job.ProfessionalRequire = row[12]
		job.EducationalRequire = row[13]
		job.DegreeRequire = row[14]

		job.PoliticalStatus = row[15]
		job.GrassrootsExperienceYear = row[16]
		job.GrassrootsExperience = row[17]

		job.ProfessionalAblilityTest = row[18]
		job.InterviewRatio = row[19]

		//工作地点
		job.WorkSpace = row[20]
		//落户地点
		job.SettledSpace = row[21]

		job.Remark = row[22]
		job.Website = row[23]
		job.EnquirtyTel = row[24]

		if row[25] != "" {
			job.EnquirtyTel += "," + row[25]

		}
		if row[26] != "" {
			job.EnquirtyTel += "," + row[26]
		}

		job.Year = year

		jobs = append(jobs, job)
		//fmt.Printf("job:%v", job)

		//fmt.Println()
	}
	fmt.Println("将要插入jobs个数:", len(jobs))

	for _, job := range jobs {
		db.Model(&GuoJob{}).
			Where("job_code_base = ? and job_code = ? and year = ?", job.JobCodeBase, job.JobCode, "2026").
			Updates(map[string]interface{}{
				"work_space":    job.WorkSpace,
				"settled_space": job.SettledSpace,
			})
	}

}
