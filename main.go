package main

import (
	"bufio"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"os"
	"regexp"
	"strings"
	"time"
)

// 2023年：https://www.gongkaoleida.com/user/exam/detail/414875?page=1
//
// 2022年：https://www.gongkaoleida.com/user/exam/detail/284219?page=1
//
// 2021年：https://www.gongkaoleida.com/user/exam/detail/173331
func main() {
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	filePath := "./yeardata/2023.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	year2023(write)
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func year2023(write *bufio.Writer) {
	var quchong = make(map[string]string, 10)

	request := gorequest.New()
	for page := 1; page <= 75; page++ {
		request.Header = map[string]string{
			":authority":                "www.gongkaoleida.com",
			":method":                   "GET",
			":path":                     fmt.Sprintf("/user/exam/detail/414875?page=%v", page),
			":scheme":                   "https",
			"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"Accept-Encoding":           "gzip, deflate, br, zstd",
			"Accept-Language":           "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6",
			"Cache-Control":             "no-cache",
			"Cookie":                    "\nHm_lvt_f721d958b1ffbdd95625a927f9bbe719=1712907152; Hm_lvt_a85566772a4d8c7093230e45128ffa8f=1717744454,1717918367; Hm_lvt_f721d958b1ffbdd95625a927f9bbe719=; XSRF-TOKEN=eyJpdiI6IkI5ZGJFa1FcL3o3WWk3WFwvcWx4MTg2dz09IiwidmFsdWUiOiJuR29MVzc0U1dUYWlQRlppcjRlSkM1TlNHVjlqb1dsUVRkRGs0RitHamNsQTRROGZEMlFuemFXNFd3VjFYVDlUIiwibWFjIjoiZmQxMzZhNWRjN2MxNGUzYjI3YzVjNGM3Y2YwODllYzU5ZTE0ZDUxMWRiMDNmZmMwODUyOTc3ZmQzZGFjNjFlYyJ9; gkld_session=eyJpdiI6IlVlVFBZM1RCRW1HcldWNE9Edkwxb3c9PSIsInZhbHVlIjoiKzdVWE16WjI1NWdQWFJpVmxzNllZVFJ1bjI1alA2WVpwQURZcjlONkNwbWVnTWQrRWlrVzJpR0NqYmhodmxWUUhleFJKWEVOZm9weHZtZ2EzOW9scVB5eHMwcUJUcmg1dnN3ekxuVXpRb0tZVVh3aE9lblFRRG90RHhwNWN2eFciLCJtYWMiOiJjOTVmMTYwZTVlNWJhYWQ0NjQ4MjM2NGJkYzM2NGJkNzFkNzM3OGQ2ZGNkNmQ2NTk0M2U2NmU4YjIwOTQ3MjliIn0%3D; Hm_lpvt_f721d958b1ffbdd95625a927f9bbe719=1717928623; Hm_lpvt_a85566772a4d8c7093230e45128ffa8f=1717928624",
			"Pragma":                    "no-cache",
			"Priority":                  "u=0, i",
			"Sec-Ch-Ua":                 `Microsoft Edge";v="125", "Chromium";v="125", "Not.A/Brand";v="24"`,
			"Sec-Ch-Ua-Mobile":          "?0",
			"Sec-Fetch-Dest":            "document",
			"Sec-Fetch-Mode":            "navigate",
			"Sec-Fetch-Site":            "none",
			"Sec-Fetch-User":            "?1",
			"Upgrade-Insecure-Requests": "1",
			"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0",
		}
		_, body, errs := request.Get(fmt.Sprintf("https://www.gongkaoleida.com/user/exam/detail/414875?page=%v", page)).End()
		fmt.Println(body, errs)
		if errs != nil {
			fmt.Println("出错！", errs)
			return
		}

		re := regexp.MustCompile(`<td><a href="([^"]+)"`)
		matches := re.FindAllStringSubmatch(body, -1)

		for _, match := range matches {
			fmt.Println(match[1])
			quchong[match[1]] = ""
		}
		fmt.Println(len(quchong))
		time.Sleep(1 * time.Second)
	}

	for key, _ := range quchong {
		ll := strings.Split(key, "/")
		write.WriteString(fmt.Sprintf("%v\n", ll[len(ll)-1]))
	}
}
