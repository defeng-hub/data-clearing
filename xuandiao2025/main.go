package main

import (
	"bytes"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func main() {
	readExcel()
	main1("221728801")
}

func readExcel() {

	f, err := excelize.OpenFile("./xuandiao2025/222.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("定向选调")
	for i, row := range rows {
		if i < 3 {
			continue
		}
		fmt.Println(">>", row[2])

		ll := strings.Split(row[2], "、")

		ss := "**"

		if len(ll) == 3 {
			ss = "**" + main1(ll[0]) + "、" + main1(ll[1]) + "、" + main1(ll[2])
			fmt.Println(ss)
		}
		if len(ll) == 2 {
			ss = "**" + main1(ll[0]) + "、" + main1(ll[1])
			fmt.Println(ss)
		}
		if len(ll) == 1 {
			ss = "**" + main1(ll[0])
			fmt.Println(ss)
		}
	}

}
func main1(iii string) string {
	// 表单数据
	formData := fmt.Sprintf("yhid=121713&zwdm=%s&mslb=1", iii)
	reqBody := bytes.NewBufferString(formData)

	// 创建请求
	req, err := http.NewRequest("POST", "https://fuwu.rsj.beijing.gov.cn/dxxdgwy/publicQuery/queryMsrymd", reqBody)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "0"
	}

	// 设置请求头
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "route=547d87bfefd67bded59518668b6a65b9; JSESSIONID=655c0760-69c6-4fe4-95cf-b01a066742a8; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%2218e0d55fe3825e-0cf174c0327105-7e433c49-3686400-18e0d55fe391998%22%2C%22props%22%3A%7B%22%24latest_referrer%22%3A%22%22%2C%22%24latest_referrer_host%22%3A%22%22%7D%7D; _trs_uv=lwrm2k8y_365_1d15; _va_ses=*; arialoadData=false; 69867b41007d48a1bc1a8c080fe543d6=WyIxMTI5NjA1MTMyIl0; 1391502a78b24453b4934f56e419c2c3=WyIyMTY2MTM4NDY1Il0; _trs_ua_s_1=m2znzy82_365_ebsh; _va_id=cd65882ef11bf752.1709620663.25.1730522414.1730521767.")
	req.Header.Set("Host", "fuwu.rsj.beijing.gov.cn")
	req.Header.Set("Origin", "https://fuwu.rsj.beijing.gov.cn")
	req.Header.Set("Referer", "https://fuwu.rsj.beijing.gov.cn/dxxdgwy/publicQuery/queryMsrymd")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "0"
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "0"
	}

	// 打印响应
	//fmt.Println("Response Status:", resp.Status)
	//fmt.Println("Response Body:", string(body))
	return main2(string(body))
}

func main2(htmlContent string) string {
	// 使用 goquery 加载 HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		panic(err)
	}

	// 查找最低分数
	var minScore string
	doc.Find("#tbl tr").Each(func(i int, s *goquery.Selection) {
		if i > 0 { // 跳过表头
			score := s.Find("td").Last().Text() // 获取最后一个 <td> 的文本
			minScore = strings.TrimSpace(score) // 去除空格
		}
	})

	if minScore != "" {
		//fmt.Printf("最低分数是: %s\n", minScore)
		return minScore
	} else {
		//fmt.Println("没有找到分数。", minScore)
		return "0"
	}
	return "0"
}
