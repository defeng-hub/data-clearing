package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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
	fmt.Println("aa")
	for page := 1; page <= 75; page++ {
		if page == 30 {
			continue
		}
		filename := fmt.Sprintf("./yeardata/2023/%d.html", page)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}
		body := string(content)
		//fmt.Println(body)

		re := regexp.MustCompile(`<td><a href="([^"]+)"`)
		matches := re.FindAllStringSubmatch(body, -1)

		for _, match := range matches {
			//fmt.Println(match[1])
			quchong[match[1]] = ""
		}
		fmt.Println(len(quchong))
		//time.Sleep(1 * time.Second)
	}

	fmt.Println("end")
	fmt.Println(len(quchong))
	for key, _ := range quchong {
		ll := strings.Split(key, "/")
		write.WriteString(fmt.Sprintf("%v\n", ll[len(ll)-1]))
	}
}

func year2022(write *bufio.Writer) {
	var quchong = make(map[string]string, 10)
	fmt.Println("aa")
	for page := 1; page <= 80; page++ {
		filename := fmt.Sprintf("./yeardata/2022/%d.html", page)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}
		body := string(content)
		//fmt.Println(body)

		re := regexp.MustCompile(`<td><a href="([^"]+)"`)
		matches := re.FindAllStringSubmatch(body, -1)

		for _, match := range matches {
			//fmt.Println(match[1])
			quchong[match[1]] = ""
		}
		fmt.Println(len(quchong))
		//time.Sleep(1 * time.Second)
	}

	fmt.Println("end")
	fmt.Println(len(quchong))
	for key, _ := range quchong {
		ll := strings.Split(key, "/")
		write.WriteString(fmt.Sprintf("%v\n", ll[len(ll)-1]))
	}
}

func year2021(write *bufio.Writer) {
	var quchong = make(map[string]string, 10)
	for page := 1; page <= 60; page++ {
		filename := fmt.Sprintf("./yeardata/2021/%d.html", page)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
			continue
		}
		body := string(content)
		//fmt.Println(body)

		re := regexp.MustCompile(`<td><a href="([^"]+)"`)
		matches := re.FindAllStringSubmatch(body, -1)

		for _, match := range matches {
			//fmt.Println(match[1])
			quchong[match[1]] = ""
		}
		fmt.Println(len(quchong))
		//time.Sleep(1 * time.Second)
	}

	fmt.Println("end")
	fmt.Println(len(quchong))
	for key, _ := range quchong {
		ll := strings.Split(key, "/")
		write.WriteString(fmt.Sprintf("%v\n", ll[len(ll)-1]))
	}
}
