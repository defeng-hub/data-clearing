package main

import (
	"dataclearing/yeardata/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	read2023()
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

	ids := strings.Split(body, "\n")
	fmt.Println("总 count", len(ids))

	var nerr int
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

		resp.Data.JobOtherList.Get("职位层级")           //区级机关
		resp.Data.JobOtherList.Get("职位类别")           //普通职位
		resp.Data.JobOtherList.Get("招聘对象")           //普通职位
		resp.Data.JobOtherList.Get("是否组织专业能力测试")     //否
		resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例") //3:1
		resp.Data.JobOtherList.Get("单位网站")           //http://www.bjdch.gov.cn/

		name := resp.Data.JobDataList.Get("职位名称") //复议审理岗
		resp.Data.JobDataList.Get("报考地区")         //北京北京市东城区
		dm := resp.Data.JobDataList.Get("职位代码")   //220119205
		resp.Data.JobDataList.Get("工作单位")         //复议审理一科
		resp.Data.JobDataList.Get("上级单位")         //北京市东城区司法局
		resp.Data.JobDataList.Get("联系电话")         //010-84217021\r、010-89945039

		resp.Data.JobConditionList.Get("专业要求")   //本科：法学类（0301）研究生：法学（0301），法律（0351）
		resp.Data.JobConditionList.Get("学历要求")   //本科及以上
		resp.Data.JobConditionList.Get("学位要求")   //取得相应学位
		resp.Data.JobConditionList.Get("政治面貌")   //不限
		resp.Data.JobConditionList.Get("服务基层项目") //不限
		resp.Data.JobConditionList.Get("年龄要求")   //18周岁以上、35周岁以下（1986年11月至2004年11月期间出生）
		resp.Data.JobConditionList.Get("备注")     //通过国家司法考试或国家统一法律职业资格考试（A类）。

		//fmt.Println(resp.Data)
		fmt.Println("name", name, "dm", dm)
	}
	fmt.Println("失败个数：", nerr)
	fmt.Println("成功个数：")
}
