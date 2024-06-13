package main

import (
	"dataclearing/conf"
	"dataclearing/gkld-jingkao-23-22-21/model"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"strings"
)

// 只有京考2023,2022,2021的数据
// 从gkld 获取 2023 2022，2021 的数据

type Job struct {
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
	V3LdCode    string `json:"v3_ld_code" gorm:"v3_ld_code"`
}

func (*Job) TableName() string {
	return "v2_guokao_jingkao"
}
func main() {
}

func read2023() {
	filename := fmt.Sprintf("./2023.txt")
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := string(content)
	//fmt.Println(body)
	var jobs []Job

	ids := strings.Split(body, "\n")
	fmt.Println("总 count", len(ids))

	var nerr, nsuccess int
	for _, id := range ids {
		if id == "" {
			continue
		}
		ctx, err := os.ReadFile(fmt.Sprintf("./data/2023/%s.json", id))
		if err != nil {
			log.Print(err)
			nerr++
			continue
		}

		resp := model.Response{}
		err1 := json.Unmarshal(ctx, &resp)
		if err1 != nil {
			log.Print("json.Unmarshal", err)
			nerr++
			continue
		}
		//fmt.Println(resp.Data)

		var job = Job{
			Year:                     "2023",
			JobCodeBase:              "", //~
			JobCode:                  resp.Data.JobDataList.Get("职位代码"),
			JobName:                  resp.Data.JobDataList.Get("职位名称"),
			UnitName:                 resp.Data.JobDataList.Get("上级单位"),
			EmployDepartment:         resp.Data.JobDataList.Get("工作单位"),
			InstitutionalNatural:     resp.Data.JobDataList.Get("单位性质"),
			JobLevel:                 resp.Data.JobOtherList.Get("职位层级"), //区级机关
			JobCategory:              resp.Data.JobOtherList.Get("职位类别"),
			JobDescription:           resp.Data.JobDataList.Get("职位介绍"),
			ExamType:                 "-",
			GrassrootsExperience:     "-",
			Count:                    strconv.FormatInt(int64(resp.Data.Info.Hires), 10),
			EducationalRequire:       resp.Data.JobConditionList.Get("学历要求"),
			DegreeRequire:            resp.Data.JobConditionList.Get("学位要求"),
			ProfessionalRequire:      resp.Data.JobConditionList.Get("专业要求"),
			PoliticalStatus:          resp.Data.JobConditionList.Get("政治面貌"),
			Other:                    resp.Data.JobConditionList.Get("年龄要求"),
			ProfessionalAblilityTest: resp.Data.JobOtherList.Get("是否组织专业能力测试"),
			InterviewRatio:           resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例"),
			GrassrootsExperienceYear: resp.Data.JobConditionList.Get("服务基层项目"),
			EnquirtyTel:              resp.Data.JobDataList.Get("联系电话"),
			Website:                  resp.Data.JobOtherList.Get("单位网站"),
			Remark:                   resp.Data.JobConditionList.Get("备注"),
			Type1:                    "京考",
			Type2:                    resp.Data.JobOtherList.Get("招聘对象"),                        //大学生士兵
			V2SubmitNum:              resp.Data.JobEnrollInfo.EnrollFieldsList.Get("过审人数"),      //过审人数
			V2MaxScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最高进面分数"), //最高进面分
			V2MinScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最低进面分数"), //最低进面分
		}
		jobs = append(jobs, job)
		//resp.Data.JobOtherList.Get("职位层级")
		//resp.Data.JobOtherList.Get("职位类别")           //普通职位
		//resp.Data.JobOtherList.Get("招聘对象")       //普通职位，大学生士兵，等
		//resp.Data.JobOtherList.Get("是否组织专业能力测试") //否
		//resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例") //3:1
		//resp.Data.JobOtherList.Get("单位网站") //http://www.bjdch.gov.cn/

		name := resp.Data.JobDataList.Get("职位名称") //复议审理岗
		//resp.Data.JobDataList.Get("报考地区")         //北京北京市东城区
		dm := resp.Data.JobDataList.Get("职位代码") //220119205
		//resp.Data.JobDataList.Get("工作单位")         //复议审理一科
		//resp.Data.JobDataList.Get("上级单位")         //北京市东城区司法局
		//resp.Data.JobDataList.Get("联系电话") //010-84217021\r、010-89945039

		//resp.Data.JobConditionList.Get("专业要求")   //本科：法学类（0301）研究生：法学（0301），法律（0351）
		//resp.Data.JobConditionList.Get("学历要求")   //本科及以上
		//resp.Data.JobConditionList.Get("学位要求")   //取得相应学位
		//resp.Data.JobConditionList.Get("政治面貌")   //不限
		//resp.Data.JobConditionList.Get("服务基层项目") //不限
		//resp.Data.JobConditionList.Get("年龄要求") //18周岁以上、35周岁以下（1986年11月至2004年11月期间出生）
		//resp.Data.JobConditionList.Get("备注")     //通过国家司法考试或国家统一法律职业资格考试（A类）。

		fmt.Println("dm", dm, "name", name)
		nsuccess++
	}
	fmt.Println("失败个数：", nerr)
	fmt.Println("成功个数：", nsuccess)
	//Insert(jobs)
}

func read2022() {
	filename := fmt.Sprintf("./2022.txt")
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := string(content)
	//fmt.Println(body)
	var jobs []Job

	ids := strings.Split(body, "\n")
	fmt.Println("总 count", len(ids))

	var nerr, nsuccess int
	for _, id := range ids {
		if id == "" {
			continue
		}
		ctx, err := os.ReadFile(fmt.Sprintf("./data/2022/%s.json", id))
		if err != nil {
			log.Print(err)
			nerr++
			continue
		}

		resp := model.Response{}
		err1 := json.Unmarshal(ctx, &resp)
		if err1 != nil {
			log.Print("json.Unmarshal", err)
			nerr++
			continue
		}
		//fmt.Println(resp.Data)

		var job = Job{
			Year:                     "2022",
			JobCodeBase:              "", //~
			JobCode:                  resp.Data.JobDataList.Get("职位代码"),
			JobName:                  resp.Data.JobDataList.Get("职位名称"),
			UnitName:                 resp.Data.JobDataList.Get("上级单位"),
			EmployDepartment:         resp.Data.JobDataList.Get("工作单位"),
			InstitutionalNatural:     resp.Data.JobDataList.Get("单位性质"),
			JobLevel:                 resp.Data.JobOtherList.Get("职位层级"), //区级机关
			JobCategory:              resp.Data.JobOtherList.Get("职位类别"),
			JobDescription:           resp.Data.JobDataList.Get("职位介绍"),
			ExamType:                 "-",
			GrassrootsExperience:     "-",
			Count:                    strconv.FormatInt(int64(resp.Data.Info.Hires), 10),
			EducationalRequire:       resp.Data.JobConditionList.Get("学历要求"),
			DegreeRequire:            resp.Data.JobConditionList.Get("学位要求"),
			ProfessionalRequire:      resp.Data.JobConditionList.Get("专业要求"),
			PoliticalStatus:          resp.Data.JobConditionList.Get("政治面貌"),
			Other:                    resp.Data.JobConditionList.Get("年龄要求"),
			ProfessionalAblilityTest: resp.Data.JobOtherList.Get("是否组织专业能力测试"),
			InterviewRatio:           resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例"),
			GrassrootsExperienceYear: resp.Data.JobConditionList.Get("服务基层项目"),
			EnquirtyTel:              resp.Data.JobDataList.Get("联系电话"),
			Website:                  resp.Data.JobOtherList.Get("单位网站"),
			Remark:                   resp.Data.JobConditionList.Get("备注"),
			Type1:                    "京考",
			Type2:                    resp.Data.JobConditionList.Get("服务基层项目"),                  //大学生士兵
			V2SubmitNum:              resp.Data.JobEnrollInfo.EnrollFieldsList.Get("过审人数"),      //过审人数
			V2MaxScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最高进面分数"), //最高进面分
			V2MinScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最低进面分数"), //最低进面分
			V3LdCode:                 id,                                                        //公考雷达code
		}
		jobs = append(jobs, job)
		//resp.Data.JobOtherList.Get("职位层级")
		//resp.Data.JobOtherList.Get("职位类别")           //普通职位
		//resp.Data.JobOtherList.Get("招聘对象")       //普通职位，大学生士兵，等
		//resp.Data.JobOtherList.Get("是否组织专业能力测试") //否
		//resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例") //3:1
		//resp.Data.JobOtherList.Get("单位网站") //http://www.bjdch.gov.cn/

		name := resp.Data.JobDataList.Get("职位名称") //复议审理岗
		//resp.Data.JobDataList.Get("报考地区")         //北京北京市东城区
		dm := resp.Data.JobDataList.Get("职位代码") //220119205
		//resp.Data.JobDataList.Get("工作单位")         //复议审理一科
		//resp.Data.JobDataList.Get("上级单位")         //北京市东城区司法局
		//resp.Data.JobDataList.Get("联系电话") //010-84217021\r、010-89945039

		//resp.Data.JobConditionList.Get("专业要求")   //本科：法学类（0301）研究生：法学（0301），法律（0351）
		//resp.Data.JobConditionList.Get("学历要求")   //本科及以上
		//resp.Data.JobConditionList.Get("学位要求")   //取得相应学位
		//resp.Data.JobConditionList.Get("政治面貌")   //不限
		//resp.Data.JobConditionList.Get("服务基层项目") //不限
		//resp.Data.JobConditionList.Get("年龄要求") //18周岁以上、35周岁以下（1986年11月至2004年11月期间出生）
		//resp.Data.JobConditionList.Get("备注")     //通过国家司法考试或国家统一法律职业资格考试（A类）。

		fmt.Println("dm", dm, "name", name)
		nsuccess++
	}
	fmt.Println("失败个数：", nerr)
	fmt.Println("成功个数：", nsuccess)
	Insert(jobs)
}

func read2021() {
	filename := fmt.Sprintf("./2021.txt")
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := string(content)
	//fmt.Println(body)
	var jobs []Job

	ids := strings.Split(body, "\n")
	fmt.Println("总 count", len(ids))

	var nerr, nsuccess int
	for _, id := range ids {
		if id == "" {
			continue
		}
		ctx, err := os.ReadFile(fmt.Sprintf("./data/2021/%s.json", id))
		if err != nil {
			log.Print(err)
			nerr++
			continue
		}

		resp := model.Response{}
		err1 := json.Unmarshal(ctx, &resp)
		if err1 != nil {
			log.Print("json.Unmarshal", err)
			nerr++
			continue
		}
		//fmt.Println(resp.Data)

		var job = Job{
			Year:                     "2021",
			JobCodeBase:              "", //~
			JobCode:                  resp.Data.JobDataList.Get("职位代码"),
			JobName:                  resp.Data.JobDataList.Get("职位名称"),
			UnitName:                 resp.Data.JobDataList.Get("上级单位"),
			EmployDepartment:         resp.Data.JobDataList.Get("工作单位"),
			InstitutionalNatural:     resp.Data.JobDataList.Get("单位性质"),
			JobLevel:                 resp.Data.JobOtherList.Get("职位层级"), //区级机关
			JobCategory:              resp.Data.JobOtherList.Get("职位类别"),
			JobDescription:           resp.Data.JobDataList.Get("职位介绍"),
			ExamType:                 "-",
			GrassrootsExperience:     "-",
			Count:                    strconv.FormatInt(int64(resp.Data.Info.Hires), 10),
			EducationalRequire:       resp.Data.JobConditionList.Get("学历要求"),
			DegreeRequire:            resp.Data.JobConditionList.Get("学位要求"),
			ProfessionalRequire:      resp.Data.JobConditionList.Get("专业要求"),
			PoliticalStatus:          resp.Data.JobConditionList.Get("政治面貌"),
			Other:                    resp.Data.JobConditionList.Get("年龄要求"),
			ProfessionalAblilityTest: resp.Data.JobOtherList.Get("是否组织专业能力测试"),
			InterviewRatio:           resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例"),
			GrassrootsExperienceYear: resp.Data.JobConditionList.Get("服务基层项目"),
			EnquirtyTel:              resp.Data.JobDataList.Get("联系电话"),
			Website:                  resp.Data.JobOtherList.Get("单位网站"),
			Remark:                   resp.Data.JobConditionList.Get("备注"),
			Type1:                    "京考",
			Type2:                    resp.Data.JobConditionList.Get("服务基层项目"),                  //大学生士兵
			V2SubmitNum:              resp.Data.JobEnrollInfo.EnrollFieldsList.Get("过审人数"),      //过审人数
			V2MaxScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最高进面分数"), //最高进面分
			V2MinScore:               resp.Data.JobInterviewInfo.EnrollFieldsList.Get("最低进面分数"), //最低进面分
			V3LdCode:                 id,                                                        //公考雷达code
		}
		jobs = append(jobs, job)
		//resp.Data.JobOtherList.Get("职位层级")
		//resp.Data.JobOtherList.Get("职位类别")           //普通职位
		//resp.Data.JobOtherList.Get("招聘对象")       //普通职位，大学生士兵，等
		//resp.Data.JobOtherList.Get("是否组织专业能力测试") //否
		//resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例") //3:1
		//resp.Data.JobOtherList.Get("单位网站") //http://www.bjdch.gov.cn/

		name := resp.Data.JobDataList.Get("职位名称") //复议审理岗
		//resp.Data.JobDataList.Get("报考地区")         //北京北京市东城区
		dm := resp.Data.JobDataList.Get("职位代码") //220119205
		//resp.Data.JobDataList.Get("工作单位")         //复议审理一科
		//resp.Data.JobDataList.Get("上级单位")         //北京市东城区司法局
		//resp.Data.JobDataList.Get("联系电话") //010-84217021\r、010-89945039

		//resp.Data.JobConditionList.Get("专业要求")   //本科：法学类（0301）研究生：法学（0301），法律（0351）
		//resp.Data.JobConditionList.Get("学历要求")   //本科及以上
		//resp.Data.JobConditionList.Get("学位要求")   //取得相应学位
		//resp.Data.JobConditionList.Get("政治面貌")   //不限
		//resp.Data.JobConditionList.Get("服务基层项目") //不限
		//resp.Data.JobConditionList.Get("年龄要求") //18周岁以上、35周岁以下（1986年11月至2004年11月期间出生）
		//resp.Data.JobConditionList.Get("备注")     //通过国家司法考试或国家统一法律职业资格考试（A类）。
		fmt.Println("dm", dm, "name", name)
		nsuccess++
	}
	fmt.Println("失败个数：", nerr)
	fmt.Println("成功个数：", nsuccess)
	Insert(jobs)
}

func Insert(jobs []Job) {

	db, err := gorm.Open(mysql.Open(conf.DSNL), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		panic("failed to connect database")
	}
	// 开始事务
	tx := db.Begin()
	var count int64
	db.Table("v2_guokao_jingkao").Count(&count)
	fmt.Println("表原始长度：", count)

	// 滚动插入
	var temp []Job
	for _, job := range jobs {
		temp = append(temp, job)
		if len(temp) == 50 {
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

	tx.Table("v2_guokao_jingkao").Count(&count2)

	fmt.Println("count2-count:", count2-count)

	if count2-count == int64(len(jobs)) {
		fmt.Println("提交事务:", tx.Commit().Error)
	} else {
		fmt.Println("回滚事务:", tx.Rollback().Error)
	}
}
