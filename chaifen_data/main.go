package main

import (
	"dataclearing/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type TbGuokaoJingkaoCopy1 struct {
	ID                       int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // 主键编码
	Year                     string         `gorm:"column:year" json:"year"`                                             // 年份
	Type1                    string         `gorm:"column:type1" json:"type1"`                                           // 国考，京考
	Type2                    string         `gorm:"column:type2" json:"type2"`                                           // 职位分类说明
	JobCodeBase              string         `gorm:"column:job_code_base" json:"job_code_base"`                           // 部门代码
	JobCode                  string         `gorm:"column:job_code;not null" json:"job_code"`                            // 职位代码
	JobName                  string         `gorm:"column:job_name" json:"job_name"`                                     // 职位名称
	UnitName                 string         `gorm:"column:unit_name" json:"unit_name"`                                   // 单位名称
	EmployDepartment         string         `gorm:"column:employ_department" json:"employ_department"`                   // 用人部门
	InstitutionalNatural     string         `gorm:"column:institutional_natural" json:"institutional_natural"`           // 机构性质
	JobLevel                 string         `gorm:"column:job_level" json:"job_level"`                                   // 职位层级
	JobCategory              string         `gorm:"column:job_category" json:"job_category"`                             // 职位类别
	JobDescription           string         `gorm:"column:job_description" json:"job_description"`                       // 职位简介
	ExamType                 string         `gorm:"column:exam_type" json:"exam_type"`                                   // 考试类别
	GrassrootsExperience     string         `gorm:"column:grassroots_experience" json:"grassroots_experience"`           // 基层项目工作经历
	Count                    string         `gorm:"column:count" json:"count"`                                           // 招考人数
	EducationalRequire       string         `gorm:"column:educational_require" json:"educational_require"`               // 学历要求
	DegreeRequire            string         `gorm:"column:degree_require" json:"degree_require"`                         // 学位要求
	ProfessionalRequire      string         `gorm:"column:professional_require" json:"professional_require"`             // 专业要求
	PoliticalStatus          string         `gorm:"column:political_status" json:"political_status"`                     // 政治面貌
	Other                    string         `gorm:"column:other" json:"other"`                                           // 其他条件
	ProfessionalAbilityTest  string         `gorm:"column:professional_ablility_test" json:"professional_ability_test"`  // 是否组织专业能力测试
	InterviewRatio           string         `gorm:"column:interview_ratio" json:"interview_ratio"`                       // 面试人员比例
	GrassrootsExperienceYear string         `gorm:"column:grassroots_experience_year" json:"grassroots_experience_year"` // 基层工作最低年限
	EnquiryTel               string         `gorm:"column:enquirty_tel" json:"enquiry_tel"`                              // 咨询电话
	Website                  string         `gorm:"column:website" json:"website"`                                       // 单位网站
	Remark                   string         `gorm:"column:remark" json:"remark"`                                         // 备注
	V2SubmitNum              string         `gorm:"column:v2_submit_num" json:"v2_submit_num"`                           // 过审人数
	V2MaxScore               string         `gorm:"column:v2_max_score" json:"v2_max_score"`                             // 最高进面分
	V2MinScore               string         `gorm:"column:v2_min_score" json:"v2_min_score"`                             // 最低进面分
	V3LdCode                 string         `gorm:"column:v3_ld_code" json:"v3_ld_code"`                                 // 公考雷达的ID
	CreatedAt                time.Time      `gorm:"column:created_at" json:"created_at"`                                 // 创建时间
	UpdatedAt                time.Time      `gorm:"column:updated_at" json:"updated_at"`                                 // 最后更新时间
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                 // 删除时间
	CreateBy                 int64          `gorm:"column:create_by" json:"create_by"`                                   // 创建者
	UpdateBy                 int64          `gorm:"column:update_by" json:"update_by"`                                   // 更新者
}

func (TbGuokaoJingkaoCopy1) TableName() string {
	return "tb_guokao_jingkao_copy1"
}

type TbGuokao struct {
	ID                       int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // 主键
	JobCodeBase              string         `gorm:"column:job_code_base" json:"job_code_base"`                           // 部门代码
	UnitName                 string         `gorm:"column:unit_name" json:"unit_name"`                                   // 部门名称
	EmployDepartment         string         `gorm:"column:employ_department" json:"employ_department"`                   // 用人司局
	InstitutionalNatural     string         `gorm:"column:instituonal_natural" json:"institutional_natural"`             // 机构与性质
	JobName                  string         `gorm:"column:job_name" json:"job_name"`                                     // 招考职位
	JobCategory              string         `gorm:"column:job_category" json:"job_category"`                             // 职位属性
	JobDistribute            string         `gorm:"column:job_distribute" json:"job_distribute"`                         // 职位分布
	JobDescription           string         `gorm:"column:job_description" json:"job_description"`                       // 职位简介
	JobCode                  string         `gorm:"column:job_code" json:"job_code"`                                     // 职位代码
	JobLevel                 string         `gorm:"column:job_level" json:"job_level"`                                   // 机构层级
	ExamType                 string         `gorm:"column:exam_type" json:"exam_type"`                                   // 考试类别
	Count                    string         `gorm:"column:count" json:"count"`                                           // 招考人数
	ProfessionalRequire      string         `gorm:"column:professional_require" json:"professional_require"`             // 专业要求
	EducationalRequire       string         `gorm:"column:educational_require" json:"educational_require"`               // 学历要求
	DegreeRequire            string         `gorm:"column:degree_require" json:"degree_require"`                         // 学位要求
	PoliticalStatus          string         `gorm:"column:political_status" json:"political_status"`                     // 政治面貌
	GrassrootsExperienceYear string         `gorm:"column:grassroots_experience_year" json:"grassroots_experience_year"` // 基层最低工作年限
	ServeExperienceYear      string         `gorm:"column:serve_experience_year" json:"serve_experience_year"`           // 服务基层项目工作经历
	ProfessionalAbilityTest  string         `gorm:"column:professional_ability_test" json:"professional_ability_test"`   // 是否在面试阶段组织专业能力测试
	InterviewRatio           string         `gorm:"column:interview_ratio" json:"interview_ratio"`                       // 面试人员比例
	WorkSpace                string         `gorm:"column:work_space" json:"work_space"`                                 // 工作地点
	SettledSpace             string         `gorm:"column:settled_apace" json:"settled_space"`                           // 落户地点
	Remark                   string         `gorm:"column:remark" json:"remark"`                                         // 备注
	Website                  string         `gorm:"column:website" json:"website"`                                       // 部门网站
	EnquiryTel               string         `gorm:"column:enquirty_tel" json:"enquiry_tel"`                              // 咨询电话
	V2SubmitNum              string         `gorm:"column:v2_submit_num" json:"v2_submit_num"`                           // 最终过审人数
	V2MinScore               string         `gorm:"column:v2_min_score" json:"v2_min_score"`                             // 最低进面分
	V3SubmitInfo             string         `gorm:"column:v3_submit_info" json:"v3_submit_info"`                         // 历史报名人数
	V3Ext                    string         `gorm:"column:v3_ext" json:"v3_ext"`                                         // 拓展字段
	CreatedAt                time.Time      `gorm:"column:created_at" json:"created_at"`                                 // 创建时间
	UpdatedAt                time.Time      `gorm:"column:updated_at" json:"updated_at"`                                 // 最后更新时间
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                 // 删除时间
	CreateBy                 int64          `gorm:"column:create_by" json:"create_by"`                                   // 创建者
	UpdateBy                 int64          `gorm:"column:update_by" json:"update_by"`                                   // 更新者
	Year                     string         `gorm:"column:year" json:"year"`                                             // 年份
}

func (TbGuokao) TableName() string {
	return "tb_guokao"
}

func Guokao() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Default})
	if err != nil {
		panic("failed to connect database")
	}
	year := "2025"

	var count int64
	db.Model(TbGuokaoJingkaoCopy1{}).Count(&count)
	fmt.Println("表原始长度：", count)

	db.Model(TbGuokaoJingkaoCopy1{}).Where("type1 = '国考' and year = ?", year).Count(&count)
	fmt.Println("year 表 长度：", count)

	var all []TbGuokaoJingkaoCopy1
	db.Model(TbGuokaoJingkaoCopy1{}).Where("type1 = '国考' and year = ?", year).Find(&all)

	fmt.Println(len(all))
	var temps []TbGuokao
	// 开始事务
	tx := db.Begin()
	for _, copy1 := range all {
		var temp = TbGuokao{
			ID:                       0,
			JobCodeBase:              copy1.JobCodeBase,
			UnitName:                 copy1.UnitName,
			EmployDepartment:         copy1.EmployDepartment,
			InstitutionalNatural:     copy1.InstitutionalNatural,
			JobName:                  copy1.JobName,
			JobCategory:              copy1.JobCategory,
			JobDistribute:            "",
			JobDescription:           copy1.JobDescription,
			JobCode:                  copy1.JobCode,
			JobLevel:                 copy1.JobLevel,
			ExamType:                 copy1.ExamType,
			Count:                    copy1.Count,
			ProfessionalRequire:      copy1.ProfessionalRequire,
			EducationalRequire:       copy1.EducationalRequire,
			DegreeRequire:            copy1.DegreeRequire,
			PoliticalStatus:          copy1.PoliticalStatus,
			GrassrootsExperienceYear: copy1.GrassrootsExperience,
			ServeExperienceYear:      "",
			ProfessionalAbilityTest:  copy1.ProfessionalAbilityTest,
			InterviewRatio:           copy1.InterviewRatio,
			WorkSpace:                "",
			SettledSpace:             "",
			Remark:                   copy1.Remark,
			Website:                  copy1.Website,
			EnquiryTel:               copy1.EnquiryTel,
			V2SubmitNum:              copy1.V2SubmitNum,
			V2MinScore:               copy1.V2MinScore,
			V3SubmitInfo:             "",
			V3Ext:                    "",
			CreatedAt:                copy1.CreatedAt,
			UpdatedAt:                copy1.UpdatedAt,
			DeletedAt:                copy1.DeletedAt,
			CreateBy:                 copy1.CreateBy,
			UpdateBy:                 copy1.CreateBy,
			Year:                     copy1.Year,
		}
		temps = append(temps, temp)
		//err1 := db.Create(&temp).Error
		//if err1 != nil {
		//	fmt.Println("errrrrrrrrrrrrrrrrr1", err1)
		//	countA++
		//} else {
		//	fmt.Println("success", temp.ID)
		//	countB++
		//}
	}

	// 批量插入，每批插入 100 条
	if err := tx.CreateInBatches(temps, 50).Error; err != nil {
		tx.Rollback() // 如果发生错误，则回滚事务
		fmt.Println("插入失败，事务已回滚：", err)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		fmt.Println("提交事务失败：", err)
	} else {
		fmt.Println("批量插入成功，事务已提交")
	}
	return
}

type TbJingkao struct {
	ID                       int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`                        // 主键
	Type2                    string         `gorm:"column:type2" json:"type2"`                                           // 职位类型
	JobCode                  string         `gorm:"column:job_code" json:"job_code"`                                     // 职位代码
	UnitName                 string         `gorm:"column:unit_name" json:"unit_name"`                                   // 单位名称
	EmployDepartment         string         `gorm:"column:employ_department" json:"employ_department"`                   // 用人部门
	InstitutionalNatural     string         `gorm:"column:instituonal_natural" json:"institutional_natural"`             // 机构性质
	JobName                  string         `gorm:"column:job_name" json:"job_name"`                                     // 职位名称
	JobLevel                 string         `gorm:"column:job_level" json:"job_level"`                                   // 职位层级
	JobCategory              string         `gorm:"column:job_category" json:"job_category"`                             // 职位类别
	JobDescription           string         `gorm:"column:job_description" json:"job_description"`                       // 职位简介
	Count                    string         `gorm:"column:count" json:"count"`                                           // 招考人数
	EducationalRequire       string         `gorm:"column:educational_require" json:"educational_require"`               // 学历要求
	DegreeRequire            string         `gorm:"column:degree_require" json:"degree_require"`                         // 学位要求
	ProfessionalRequire      string         `gorm:"column:professional_require" json:"professional_require"`             // 专业要求
	PoliticalStatus          string         `gorm:"column:political_status" json:"political_status"`                     // 政治面貌
	Other                    string         `gorm:"column:other" json:"other"`                                           // 其他条件
	ProfessionalAbilityTest  string         `gorm:"column:professional_ability_test" json:"professional_ability_test"`   // 是否在面试阶段组织专业能力测试
	InterviewRatio           string         `gorm:"column:interview_ratio" json:"interview_ratio"`                       // 面试人员比例
	GrassrootsExperienceYear string         `gorm:"column:grassroots_experience_year" json:"grassroots_experience_year"` // 基层最低工作年限
	EnquiryTel               string         `gorm:"column:enquirty_tel" json:"enquiry_tel"`                              // 考生咨询电话
	Website                  string         `gorm:"column:website" json:"website"`                                       // 单位网站
	Remark                   string         `gorm:"column:remark" json:"remark"`                                         // 备注
	V2SubmitNum              string         `gorm:"column:v2_submit_num" json:"v2_submit_num"`                           // 过审人数
	V2MinScore               string         `gorm:"column:v2_min_score" json:"v2_min_score"`                             // 最低进面分
	V3SubmitInfo             string         `gorm:"column:v3_submit_info" json:"v3_submit_info"`                         // 历史报名人数
	V3Ext                    string         `gorm:"column:v3_ext" json:"v3_ext"`                                         // 拓展字段
	CreatedAt                time.Time      `gorm:"column:created_at" json:"created_at"`                                 // 创建时间
	UpdatedAt                time.Time      `gorm:"column:updated_at" json:"updated_at"`                                 // 最后更新时间
	DeletedAt                gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`                                 // 删除时间
	CreateBy                 int64          `gorm:"column:create_by" json:"create_by"`                                   // 创建者
	UpdateBy                 int64          `gorm:"column:update_by" json:"update_by"`                                   // 更新者
	Year                     string         `gorm:"column:year" json:"year"`                                             // 年份
}

func (TbJingkao) TableName() string {
	return "tb_jingkao"
}
func Jingkao() {
	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Default})
	if err != nil {
		panic("failed to connect database")
	}
	year := "2025"

	var count int64
	db.Model(TbGuokaoJingkaoCopy1{}).Count(&count)
	fmt.Println("表原始长度：", count)

	db.Model(TbGuokaoJingkaoCopy1{}).Where("type1 = '京考' and year = ?", year).Count(&count)
	fmt.Println("year 表 长度：", count)

	var all []TbGuokaoJingkaoCopy1
	db.Model(TbGuokaoJingkaoCopy1{}).Where("type1 = '京考' and year = ?", year).Find(&all)

	fmt.Println(len(all))
	var temps []TbJingkao
	// 开始事务
	tx := db.Begin()
	for _, copy1 := range all {
		var temp = TbJingkao{
			ID:                       0,
			Type2:                    copy1.Type2,
			JobCode:                  copy1.JobCode,
			UnitName:                 copy1.UnitName,
			EmployDepartment:         copy1.EmployDepartment,
			InstitutionalNatural:     copy1.InstitutionalNatural,
			JobName:                  copy1.JobName,
			JobLevel:                 copy1.JobLevel,
			JobCategory:              copy1.JobCategory,
			JobDescription:           copy1.JobDescription,
			Count:                    copy1.Count,
			EducationalRequire:       copy1.EducationalRequire,
			DegreeRequire:            copy1.DegreeRequire,
			ProfessionalRequire:      copy1.ProfessionalRequire,
			PoliticalStatus:          copy1.PoliticalStatus,
			Other:                    copy1.Other,
			ProfessionalAbilityTest:  copy1.ProfessionalAbilityTest,
			InterviewRatio:           copy1.InterviewRatio,
			GrassrootsExperienceYear: copy1.GrassrootsExperienceYear,
			EnquiryTel:               copy1.EnquiryTel,
			Website:                  copy1.Website,
			Remark:                   copy1.Remark,
			V2SubmitNum:              copy1.V2SubmitNum,
			V2MinScore:               copy1.V2MinScore,
			V3Ext:                    "",
			V3SubmitInfo:             "",
			CreatedAt:                copy1.CreatedAt,
			UpdatedAt:                copy1.UpdatedAt,
			DeletedAt:                copy1.DeletedAt,
			CreateBy:                 copy1.CreateBy,
			UpdateBy:                 copy1.UpdateBy,
			Year:                     copy1.Year,
		}
		temps = append(temps, temp)
	}

	// 批量插入，每批插入 100 条
	if err := tx.CreateInBatches(temps, 50).Error; err != nil {
		tx.Rollback() // 如果发生错误，则回滚事务
		fmt.Println("插入失败，事务已回滚：", err)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		fmt.Println("提交事务失败：", err)
	} else {
		fmt.Println("批量插入成功，事务已提交")
	}
	return
}
func main() {
	Jingkao()
}
