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

	var n int
	for _, id := range ids {
		if id == "" {
			continue
		}
		ctx, err := os.ReadFile(fmt.Sprintf("./data/2023/%s.json", id))
		if err != nil {
			log.Print(err)
			n++
			continue
		}

		resp := model.Response{}
		err1 := json.Unmarshal(ctx, &resp)
		if err1 != nil {
			log.Print("json.Unmarshal", err)
			n++
			continue
		}

		resp.Data.JobOtherList.Get("职位层级")           //区级机关
		resp.Data.JobOtherList.Get("职位类别")           //普通职位
		resp.Data.JobOtherList.Get("招聘对象")           //普通职位
		resp.Data.JobOtherList.Get("是否组织专业能力测试")     //否
		resp.Data.JobOtherList.Get("面试人数与计划录用人数的比例") //3:1
		resp.Data.JobOtherList.Get("单位网站")           //http://www.bjdch.gov.cn/

		//fmt.Println(resp.Data)
		break
	}
	fmt.Println("err count", n)
}
